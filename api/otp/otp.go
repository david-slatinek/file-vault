package otp

import (
	"crypto/rand"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"main/config"
)

type OTP struct {
	options totp.GenerateOpts
}

func NewOtp(cfg *config.Config) *OTP {
	return &OTP{
		options: totp.GenerateOpts{
			Issuer:    cfg.App.Issuer,
			Period:    cfg.App.Valid,
			Algorithm: otp.AlgorithmSHA1,
			Rand:      rand.Reader,
		},
	}
}

func (receiver OTP) GenerateOTP(email string) (string, string, error) {
	receiver.options.AccountName = email

	key, err := totp.Generate(receiver.options)
	if err != nil {
		return "", "", err
	}

	return key.Secret(), key.URL(), nil
}

func (receiver OTP) Valid(code, secret string) bool {
	return totp.Validate(code, secret)
}
