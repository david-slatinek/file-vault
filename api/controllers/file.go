package controllers

import (
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/sha3"
	"log"
	"main/db"
	"main/models/request"
	"main/models/response"
	"main/storage"
	"main/validator"
	"net/http"
)

type File struct {
	FileDB  *db.File
	UserDB  *db.User
	Storage *storage.Storage
}

func (receiver File) Upload(context *gin.Context) {
	email := context.MustGet("email").(string)

	user, err := receiver.UserDB.GetByEmail(email)
	if errors.Is(err, db.UserNotFound) {
		context.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
		return
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	var req request.FileUpload
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
		return
	}

	validPassword := validator.ValidatePassword(req.Password)
	if !validPassword {
		context.JSON(http.StatusBadRequest, response.Error{Message: validator.InvalidPassword})
		return
	}

	valid, err := receiver.UserDB.ValidCode(user, req.Code)
	if err != nil || !valid {
		context.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
		return
	}

	file, err := req.File.Open()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	buf := make([]byte, req.File.Size)

	_, err = file.Read(buf)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	hash := sha3.New512()
	_, err = hash.Write(buf)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	id := hex.EncodeToString(hash.Sum(nil))

	for _, value := range user.Files {
		if value.ID == id {
			context.JSON(http.StatusBadRequest, response.Error{Message: "file already exists"})
			return
		}
	}

	createdFile, err := receiver.FileDB.Create(user.ID, id, req.File.Filename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	salt, err := receiver.Storage.Upload(id, req.Password, buf)
	if err != nil {
		err = receiver.FileDB.Delete(createdFile)
		if err != nil {
			log.Printf("failed to delete file: %s\n", err)
		}

		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}
	createdFile.Salt = salt

	err = receiver.FileDB.UpdateSalt(createdFile, salt)
	if err != nil {
		err = receiver.FileDB.Delete(createdFile)
		if err != nil {
			log.Printf("failed to delete file: %s\n", err)
		}

		err = receiver.Storage.Delete(id)
		if err != nil {
			log.Printf("failed to delete file: %s\n", err)
		}

		context.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"fileID": id})
}
