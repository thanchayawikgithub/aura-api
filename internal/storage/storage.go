package storage

import (
	"aura/internal/config"
	"aura/internal/model"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	ModelType interface {
		*model.Post | *model.User
	}

	IStorage[T ModelType] interface {
		FindByID(ctx context.Context, id uint) (data T, err error)
		FindAll(ctx context.Context) (data []T, err error)
		Insert(ctx context.Context, data T) error
		Update(ctx context.Context, id uint, data T) error
		Delete(ctx context.Context, data T) error
	}

	AbstractStorage[T ModelType] struct {
		db        *gorm.DB
		tableName string
	}

	Storage struct {
		db *gorm.DB
	}
)

func New(cfg *config.Database) *Storage {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		panic(err)
	}

	db, err := conn.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.MaxIdleTime * time.Second)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.MaxLifeTime * time.Second)

	return &Storage{db: conn}
}

func (s *AbstractStorage[T]) FindByID(ctx context.Context, id uint) (data T, err error) {
	err = s.db.Table(s.tableName).WithContext(ctx).Where("id = ?", id).First(&data).Error
	return data, err
}

func (s *AbstractStorage[T]) FindAll(ctx context.Context) (data []T, err error) {
	err = s.db.Table(s.tableName).WithContext(ctx).Find(&data).Error
	return data, err
}

func (s *AbstractStorage[T]) Insert(ctx context.Context, data T) error {
	err := s.db.Table(s.tableName).WithContext(ctx).Create(&data).Error
	return err
}

func (s *AbstractStorage[T]) Update(ctx context.Context, id uint, data T) error {
	err := s.db.Table(s.tableName).WithContext(ctx).Where("id = ?", id).Updates(&data).Error
	return err
}

func (s *AbstractStorage[T]) Delete(ctx context.Context, data T) error {
	err := s.db.Table(s.tableName).WithContext(ctx).Delete(&data).Error
	return err
}
