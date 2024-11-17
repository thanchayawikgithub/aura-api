package storage

import (
	"aura/internal/model"
	"context"
)

type (
	IPostStorage interface {
		IStorage[*model.Post]
		FindByUserID(ctx context.Context, userID uint) (result []*model.Post, err error)
	}

	PostStorage struct {
		AbstractStorage[*model.Post]
	}
)

func NewPostStorage(s *Storage) *PostStorage {
	return &PostStorage{
		AbstractStorage: AbstractStorage[*model.Post]{
			db:        s.db,
			tableName: model.PostTableName,
		},
	}
}

func (s *PostStorage) FindByUserID(ctx context.Context, userID uint) (result []*model.Post, err error) {
	err = s.db.WithContext(ctx).Where("user_id = ?", userID).Find(&result).Error
	return result, err
}
