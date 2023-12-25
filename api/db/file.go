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

func (receiver File) Create(userID int, id, filename string) (models.File, error) {
	file := models.File{
		ID:       id,
		Filename: filename,
		UserID:   userID,
	}

	res := receiver.db.Create(&file)
	if res.Error != nil {
		return models.File{}, res.Error
	}

	return file, nil
}

func (receiver File) UpdateSalt(file models.File, salt string) error {
	return receiver.db.Model(&file).Update("salt", salt).Error
}

func (receiver File) Delete(file models.File) error {
	return receiver.db.Delete(&file).Error
}

func (receiver File) UpdateAccessedAt(file models.File) error {
	return receiver.db.Model(&file).Update("accessed_at", "now()").Error
}
