package database

import (
	"aura/internal/storage"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type AbstractStorage[T storage.ModelType] struct {
	mock.Mock
}

func (s *AbstractStorage[T]) WithTx(tx *gorm.DB) storage.IStorage[T] {
	args := s.Called(tx)
	return args.Get(0).(storage.IStorage[T])
}

func (s *AbstractStorage[T]) WithPreload(preloads ...string) storage.IStorage[T] {
	args := s.Called(preloads)
	return args.Get(0).(storage.IStorage[T])
}

func (s *AbstractStorage[T]) Save(ctx context.Context, model T) (result T, err error) {
	args := s.Called(ctx, model)
	return args.Get(0).(T), args.Error(1)
}

func (s *AbstractStorage[T]) FindByID(ctx context.Context, id uint) (result T, err error) {
	args := s.Called(ctx, id)
	return args.Get(0).(T), args.Error(1)
}

func (s *AbstractStorage[T]) FindAll(ctx context.Context) (result []T, err error) {
	args := s.Called(ctx)
	return args.Get(0).([]T), args.Error(1)
}

func (s *AbstractStorage[T]) Update(ctx context.Context, id uint, model T) (result T, err error) {
	args := s.Called(ctx, id, model)
	return args.Get(0).(T), args.Error(1)
}

func (s *AbstractStorage[T]) Delete(ctx context.Context, model T) (err error) {
	args := s.Called(ctx, model)
	return args.Error(0)
}
