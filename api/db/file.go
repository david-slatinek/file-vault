package db

import (
	"gorm.io/gorm"
	"main/config"
	"main/models"
)

type File struct {
	db *gorm.DB
}

func NewFile(cfg *config.Config) (*File, error) {
	db, err := NewDB(*cfg)
	if err != nil {
		return nil, err
	}

	return &File{
		db: db,
	}, nil
}

func (receiver File) Create(userID int, id string) error {
	file := models.File{
		ID:     id,
		UserID: userID,
	}

	return receiver.db.Create(&file).Error
}
