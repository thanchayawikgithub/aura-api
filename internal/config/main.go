package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   Server   `mapstructure:"server" validate:"required"`
		Database Database `mapstructure:"database" validate:"required"`
	}

	Server struct {
		Port        int    `mapstructure:"port" validate:"required"`
		Environment string `mapstructure:"environment" validate:"required"`
	}

	Database struct {
		Host         string        `mapstructure:"host" validate:"required"`
		Port         int           `mapstructure:"port" validate:"required"`
		Name         string        `mapstructure:"name" validate:"required"`
		Username     string        `mapstructure:"username" validate:"required"`
		Password     string        `mapstructure:"password" validate:"required"`
		MaxIdleConns int           `mapstructure:"maxIdleConns" validate:"required"`
		MaxIdleTime  time.Duration `mapstructure:"maxIdleTime" validate:"required"`
		MaxOpenConns int           `mapstructure:"maxOpenConns" validate:"required"`
		MaxLifeTime  time.Duration `mapstructure:"maxLifeTime" validate:"required"`
	}
)

func LoadConfig() *Config {
	viper.SetConfigName("config.local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/etc")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config: %w", err))
	}

	err = validator.New().Struct(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error validating config: %w", err))
	}

	return &config
}
