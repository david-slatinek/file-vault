package db

import (
	"database/sql"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/config"
	"main/models"
	"main/otp"
	"main/pki"
	"time"
)

var (
	UserAlreadyExists         = errors.New("user already exists")
	UserNotFoundOrInvalidCode = errors.New("user not found or invalid code")
	UserNotFound              = errors.New("user not found")
)

func NewDB(cfg config.Config) (*gorm.DB, error) {
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

	return db, nil
}

type User struct {
	db        *gorm.DB
	otpClient *otp.OTP
	pkiClient *pki.PKI
}

func NewUser(cfg *config.Config) (*User, error) {
	db, err := NewDB(*cfg)
	if err != nil {
		return nil, err
	}

	p, err := pki.New(*cfg)
	if err != nil {
		return nil, err
	}

	return &User{
		db:        db,
		otpClient: otp.New(cfg),
		pkiClient: p,
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

	encrypted, err := receiver.pkiClient.Encrypt(secret)
	if err != nil {
		return "", "", err
	}
	user.Secret = encrypted

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

	valid, err := receiver.ValidCode(user, code)
	if err != nil || !valid {
		return UserNotFoundOrInvalidCode
	}

	user.AccessedAt = time.Now()
	_ = receiver.db.Updates(&user)

	return nil
}

func (receiver User) GetByEmail(email string) (models.User, error) {
	user := models.User{}

	result := receiver.db.Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return models.User{}, UserNotFound
	}

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (receiver User) ValidCode(user models.User, code string) (bool, error) {
	secret, err := receiver.pkiClient.Decrypt(user.Secret)
	if err != nil {
		return false, err
	}

	if !receiver.otpClient.Valid(code, secret) {
		return false, UserNotFoundOrInvalidCode
	}

	return true, nil
}