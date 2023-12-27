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
	UserAlreadyExists = errors.New("user already exists")
	InvalidCode       = errors.New("invalid code")
	UserNotFound      = errors.New("user not found")
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

	p, err := pki.NewPKI(*cfg)
	if err != nil {
		return nil, err
	}

	return &User{
		db:        db,
		otpClient: otp.NewOtp(cfg),
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

func (receiver User) GetByEmail(email string) (models.User, error) {
	user := models.User{}

	result := receiver.db.Where("email = ?", email).Preload("Files", func(db *gorm.DB) *gorm.DB {
		return db.Order("accessed_at desc")
	}).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
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
		return false, InvalidCode
	}

	return true, nil
}

func (receiver User) UpdateAccessedAt(user models.User) error {
	return receiver.db.Model(&user).Update("accessed_at", "now()").Error
}
