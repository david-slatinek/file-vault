package db

import (
	"database/sql"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/config"
	"main/models"
	"main/otp"
	"time"
)

var (
	UserAlreadyExists         = errors.New("user already exists")
	UserNotFoundOrInvalidCode = errors.New("user not found or invalid code")
)

type User struct {
	db        *gorm.DB
	otpClient *otp.OTP
}

func New(cfg *config.Config) (*User, error) {
	sqlDB, err := sql.Open("pgx", cfg.Database.ConnectionString)
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

	return &User{
		db:        db,
		otpClient: otp.New(cfg),
	}, nil
}

func (receiver User) Register(email string) (string, string, error) {
	user := models.User{}
	result := receiver.db.Where("email = ?", email).First(&user)

	if result.RowsAffected != 0 {
		return "", "", UserAlreadyExists
	}

	user.Email = email

	secret, url, err := receiver.otpClient.GenerateOTP(user.Email)
	if err != nil {
		return "", "", err
	}
	user.Secret = secret

	result = receiver.db.Create(&user)
	if result.Error != nil {
		return "", "", result.Error
	}

	return secret, url, nil
}

func (receiver User) Login(email, code string) error {
	user := models.User{}
	result := receiver.db.Where("email = ?", email).First(&user)

	if result.RowsAffected == 0 {
		return UserNotFoundOrInvalidCode
	}

	if !receiver.otpClient.Valid(code, user.Secret) {
		return UserNotFoundOrInvalidCode
	}

	return nil
}
