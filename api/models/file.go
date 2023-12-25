package models

import "time"

type File struct {
	ID         string `gorm:"primaryKey"`
	Filename   string `gorm:"not null"`
	Salt       string
	UserID     int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	AccessedAt time.Time `gorm:"default:now()"`
}
