package db

import (
	"database/sql"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/models"
	"time"
)

var (
	UserAlreadyExists = errors.New("user already exists")
)

type User struct {
	db *gorm.DB
}

func NewUser(connectionString string) (*User, error) {
	sqlDB, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &User{db: db}, nil
}

func (receiver User) Register(email string) error {
	user := models.User{}
	result := receiver.db.Where("email = ?", email).First(&user)

	if result.RowsAffected != 0 {
		return UserAlreadyExists
	}

	user.Email = email
	user.Secret = "secret"

	return receiver.db.Create(&user).Error
}
