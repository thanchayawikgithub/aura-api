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
		Save(ctx context.Context, data T) (T, error)
		Update(ctx context.Context, id uint, data T) error
		Delete(ctx context.Context, data T) error
		WithPreload(preloads ...string) IStorage[T]
	}

	AbstractStorage[T ModelType] struct {
		db        *gorm.DB
		tableName string
		preloads  []string
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

	conn.AutoMigrate(&model.User{}, &model.Post{})
	return &Storage{db: conn}
}

func (s *AbstractStorage[T]) WithPreload(preloads ...string) IStorage[T] {
	s.preloads = preloads
	return s
}

// Modify FindByID to use preloads
func (s *AbstractStorage[T]) FindByID(ctx context.Context, id uint) (data T, err error) {
	query := s.db.Table(s.tableName).WithContext(ctx)
	for _, preload := range s.preloads {
		query = query.Preload(preload)
	}
	err = query.Where("id = ?", id).First(&data).Error
	s.preloads = nil // Reset preloads after query
	return data, err
}

// Modify FindAll to use preloads
func (s *AbstractStorage[T]) FindAll(ctx context.Context) (data []T, err error) {
	query := s.db.Table(s.tableName).WithContext(ctx)
	for _, preload := range s.preloads {
		query = query.Preload(preload)
	}
	err = query.Find(&data).Error
	s.preloads = nil // Reset preloads after query
	return data, err
}

// func (s *AbstractStorage[T]) FindByID(ctx context.Context, id uint) (data T, err error) {
// 	err = s.db.Table(s.tableName).WithContext(ctx).Where("id = ?", id).First(&data).Error
// 	return data, err
// }

// func (s *AbstractStorage[T]) FindAll(ctx context.Context) (data []T, err error) {
// 	err = s.db.Table(s.tableName).WithContext(ctx).Find(&data).Error
// 	return data, err
// }

func (s *AbstractStorage[T]) Save(ctx context.Context, data T) (T, error) {
	err := s.db.Table(s.tableName).WithContext(ctx).Save(&data).Error
	return data, err
}

func (s *AbstractStorage[T]) Update(ctx context.Context, id uint, data T) error {
	err := s.db.Table(s.tableName).WithContext(ctx).Where("id = ?", id).Updates(&data).Error
	return err
}

func (s *AbstractStorage[T]) Delete(ctx context.Context, data T) error {
	err := s.db.Table(s.tableName).WithContext(ctx).Delete(&data).Error
	return err
}
