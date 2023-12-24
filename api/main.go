package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"main/auth"
	"main/config"
	"main/controllers"
	"main/db"
	"main/models/response"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	userDb, err := db.New(cfg)
	if err != nil {
		log.Fatalf("error connecting to users database: %v", err)
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

	usersGroup := router.Group("api/v1").Use(auth.ValidateToken)
	{
		usersGroup.POST("/register", userController.Register)
		usersGroup.POST("/login", userController.Login)
	}

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		log.Println("server is up at: " + srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
