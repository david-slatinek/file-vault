package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/db"
	"main/models/response"
	"net/http"
)

type User struct {
	UserDB *db.User
}

func (receiver User) Register(context *gin.Context) {
	key, url, err := receiver.UserDB.Register(context.MustGet("email").(string))

	if errors.Is(err, db.UserAlreadyExists) {
		context.JSON(http.StatusBadRequest, response.Error{Message: "email already exists"})
		return
	}

	context.JSON(http.StatusCreated, response.OTP{Key: key, URL: url})
}
