package main

import (
	"context"
	"crypto/tls"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"main/auth"
	"main/config"
	"main/controllers"
	"main/db"
	"main/models/response"
	"main/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configPathFlag = flag.String("config", "./config", "path to config file")
	configFileFlag = flag.String("file", "config", "config file name")
	configTypeFlag = flag.String("type", "yaml", "config file extension")
)

func main() {
	flag.Parse()

	cfg, err := config.NewConfig(*configPathFlag, *configFileFlag, *configTypeFlag)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	userDb, err := db.NewUser(cfg)
	if err != nil {
		log.Fatalf("error connecting to users database: %v", err)
	}

	fileDb, err := db.NewFile(cfg)
	if err != nil {
		log.Fatalf("error connecting to files database: %v", err)
	}

	stg, err := storage.New(cfg)
	if err != nil {
		log.Fatalf("error connecting to storage: %v", err)
	}

	gin.SetMode(cfg.Server.Mode)

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Error{Message: "endpoint not found"})
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Error{Message: "method not allowed"})
	})

	userController := controllers.User{
		UserDB: userDb,
	}

	fileController := controllers.File{
		FileDB:  fileDb,
		UserDB:  userDb,
		Storage: stg,
	}

	router.Use(auth.ValidateToken)

	usersGroup := router.Group("api/v1")
	{
		usersGroup.POST("/register", userController.Register)
		usersGroup.GET("/files", userController.Files)
	}

	filesGroup := router.Group("api/v1")
	{
		filesGroup.POST("/upload", fileController.Upload)
		filesGroup.GET("/download/:id", fileController.Download)
		filesGroup.DELETE("/delete/:id", fileController.Delete)
	}

	cer, err := tls.LoadX509KeyPair(cfg.TLS.Cert, cfg.TLS.Key)
	if err != nil {
		log.Fatalf("error loading tls certificate: %v", err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
		TLSConfig:    tlsConfig,
	}

	go func() {
		log.Println("server is up at: " + srv.Addr)
		if err := srv.ListenAndServeTLS("", ""); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ListenAndServe() error: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown() error: %s\n", err)
	}

	log.Println("shutting down")
}
