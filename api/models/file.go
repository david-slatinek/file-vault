package models

import "time"

type File struct {
	ID         string    `gorm:"primaryKey"`
	Filename   string    `gorm:"not null"`
	Salt       string    `json:"salt,omitempty"`
	UserID     int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	AccessedAt time.Time `gorm:"default:now()"`
}

type FileDto struct {
	ID         string `json:"id"`
	Filename   string `json:"filename"`
	CreatedAt  string `json:"createdAt"`
	AccessedAt string `json:"accessedAt"`
}
