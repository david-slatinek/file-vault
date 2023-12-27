package controllers

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"main/db"
	"main/models"
	"main/models/response"
	"net/http"
	"time"
)

type User struct {
	UserDB *db.User
}

func (receiver User) Register(context *gin.Context) {
	key, url, err := receiver.UserDB.Register(context.MustGet("email").(string))

	if errors.Is(err, db.UserAlreadyExists) {
		context.JSON(http.StatusBadRequest, response.Error{Message: db.UserAlreadyExists.Error()})
		return
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	pngBase64 := base64.StdEncoding.EncodeToString(png)

	context.JSON(http.StatusCreated, response.OTP{Key: key, URL: pngBase64})
}

func (receiver User) Files(context *gin.Context) {
	email := context.MustGet("email").(string)

	user, err := receiver.UserDB.GetByEmail(email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	user.AccessedAt = time.Now()
	_ = receiver.UserDB.UpdateAccessedAt(user)

	if len(user.Files) == 0 {
		context.JSON(http.StatusNoContent, nil)
		return
	}

	var files response.Files

	for _, file := range user.Files {
		files.Files = append(files.Files, models.FileDto{
			ID:         file.ID,
			Filename:   file.Filename,
			CreatedAt:  file.CreatedAt.Format("2006-01-02 15:04:05"),
			AccessedAt: file.AccessedAt.Format("2006-01-02 15:04:05"),
		})
	}

	context.JSON(http.StatusOK, files)
}
