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

func (receiver File) Create(userID int) (int, error) {
	file := models.File{
		UserID: userID,
	}

	result := receiver.db.Create(&file)
	if result.Error != nil {
		return 0, result.Error
	}

	return file.ID, nil
}
