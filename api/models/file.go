package models

import "time"

type File struct {
	ID         string    `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	AccessedAt time.Time `gorm:"default:now()"`
}
