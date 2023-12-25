package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/db"
	"main/models/response"
	"net/http"
)

type File struct {
	FileDB *db.File
	UserDB *db.User
}

func (receiver File) Upload(context *gin.Context) {
	email := context.MustGet("email").(string)

	user, err := receiver.UserDB.GetByEmail(email)
	if err != nil && errors.Is(err, db.UserNotFound) {
		context.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
		return
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	fileID, err := receiver.FileDB.Create(user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"fileID": fileID})
}
