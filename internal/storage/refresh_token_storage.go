package storage

import (
	"aura/internal/model"
	"context"
)

type (
	IRefreshTokenStorage interface {
		IStorage[*model.RefreshToken]
		GetByToken(ctx context.Context, token string) (result *model.RefreshToken, err error)
	}

	RefreshTokenStorage struct {
		AbstractStorage[*model.RefreshToken]
	}
)

func NewRefreshTokenStorage(s *Storage) *RefreshTokenStorage {
	return &RefreshTokenStorage{
		AbstractStorage: AbstractStorage[*model.RefreshToken]{
			db:        s.db,
			tableName: model.RefreshTokenTableName,
		},
	}
}

func (s *RefreshTokenStorage) GetByToken(ctx context.Context, token string) (result *model.RefreshToken, err error) {
	err = s.db.WithContext(ctx).Table(model.RefreshTokenTableName).
		Where("token = ?", token).
		First(&result).Error

	return result, err
}
