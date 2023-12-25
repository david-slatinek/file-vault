package models

import "time"

type File struct {
	ID         string    `gorm:"primaryKey"`
	Filename   string    `gorm:"not null"`
	Salt       string    `json:"salt,omitempty"`
	UserID     int       `gorm:"not null" json:"userID,omitempty"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	AccessedAt time.Time `gorm:"default:now()"`
}
