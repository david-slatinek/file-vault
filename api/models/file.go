package models

import "time"

type File struct {
	ID         int       `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	AccessedAt time.Time `gorm:"default:now()"`
}
