package config

import (
	"github.com/spf13/viper"
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
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.loadConfig()

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (receiver *Config) loadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&receiver)
}
