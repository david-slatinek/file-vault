package controllers

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"main/db"
	"main/models"
	"main/models/request"
	"main/models/response"
	"net/http"
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

func (receiver User) Login(context *gin.Context) {
	var req request.Login

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
		return
	}

	err := receiver.UserDB.Login(context.MustGet("email").(string), req.Code)

	if errors.Is(err, db.UserNotFound) {
		context.JSON(http.StatusBadRequest, response.Error{Message: db.UserNotFound.Error()})
		return
	} else if errors.Is(err, db.InvalidCode) {
		context.JSON(http.StatusBadRequest, response.Error{Message: db.InvalidCode.Error()})
		return
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func (receiver User) Files(context *gin.Context) {
	email := context.MustGet("email").(string)

	user, err := receiver.UserDB.GetByEmail(email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	if len(user.Files) == 0 {
		context.JSON(http.StatusNoContent, nil)
		return
	}

	var files response.Files

	for _, file := range user.Files {
		files.Files = append(files.Files, models.File{
			ID:         file.ID,
			Filename:   file.Filename,
			CreatedAt:  file.CreatedAt,
			AccessedAt: file.AccessedAt,
		})
	}

	context.JSON(http.StatusOK, files)
}
