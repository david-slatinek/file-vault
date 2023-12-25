package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/db"
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

	context.JSON(http.StatusCreated, response.OTP{Key: key, URL: url})
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
