package models

import "time"

type User struct {
	ID         int       `gorm:"primaryKey,autoIncrement"`
	Email      string    `gorm:"unique"`
	Secret     string    `gorm:"unique"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	AccessedAt time.Time `gorm:"default:now()"`
	Files      []File    `gorm:"foreignKey:user_id"`
}
