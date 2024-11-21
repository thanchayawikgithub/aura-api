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
		JWT      JWT      `mapstructure:"jwt" validate:"required"`
		MinIO    MinIO    `mapstructure:"minio" validate:"required"`
		Redis    Redis    `mapstructure:"redis" validate:"required"`
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

	JWT struct {
		SecretKey             string        `mapstructure:"secretKey" validate:"required"`
		AccessTokenExpiresAt  time.Duration `mapstructure:"accessTokenExpiresAt" validate:"required"`
		RefreshTokenExpiresAt time.Duration `mapstructure:"refreshTokenExpiresAt" validate:"required"`
	}

	MinIO struct {
		Endpoint        string `mapstructure:"endpoint" validate:"required"`
		AccessKeyID     string `mapstructure:"accessKeyID" validate:"required"`
		SecretAccessKey string `mapstructure:"secretAccessKey" validate:"required"`
		UseSSL          bool   `mapstructure:"useSSL" validate:"required"`
	}

	Redis struct {
		Host     string        `mapstructure:"host" validate:"required"`
		Port     int           `mapstructure:"port" validate:"required"`
		DB       int           `mapstructure:"db"`
		Duration time.Duration `mapstructure:"duration" validate:"required"`
	}
)

const (
	defaultConfigPath = "./internal/etc"
	defaultConfigName = "config"
)

func LoadConfig(name string) *Config {
	if name == "" {
		name = defaultConfigName
	}

	viper.SetConfigName(name)
	viper.AddConfigPath(defaultConfigPath)
	viper.SetConfigType("yml")
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
