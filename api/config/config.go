package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Database struct {
		ConnectionString string `mapstructure:"connection-string"`
	}
	Server struct {
		Mode    string
		Address string
	}
	App struct {
		Issuer string
		Valid  uint
	}
	PKI struct {
		PublicKey  string `mapstructure:"public-key"`
		PrivateKey string `mapstructure:"private-key"`
	}
	S3 struct {
		Endpoint           string
		Bucket             string
		AwsAccessKeyId     string `mapstructure:"aws-access-key-id"`
		AwsSecretAccessKey string `mapstructure:"aws-secret-access-key"`
		AwsDefaultRegion   string `mapstructure:"aws-default-region"`
	}
	TLS struct {
		Cert string `mapstructure:"certificate"`
		Key  string
	}
}

func NewConfig(path, filename, filetype string) (*Config, error) {
	cfg := &Config{}
	err := cfg.loadConfig(path, filename, filetype)

	if err != nil {
		return nil, err
	}

	err = cfg.set()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (receiver *Config) loadConfig(path, filename, filetype string) error {
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	viper.SetConfigType(filetype)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&receiver)
}

func (receiver *Config) set() error {
	err := os.Setenv("AWS_ACCESS_KEY_ID", receiver.S3.AwsAccessKeyId)
	if err != nil {
		return err
	}

	err = os.Setenv("AWS_SECRET_ACCESS_KEY", receiver.S3.AwsSecretAccessKey)
	if err != nil {
		return err
	}

	err = os.Setenv("AWS_DEFAULT_REGION", receiver.S3.AwsDefaultRegion)
	if err != nil {
		return err
	}

	return nil
}
