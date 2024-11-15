package storage

import (
	"aura/internal/model"
)

type (
	IPostStorage interface {
		IStorage[*model.Post]
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
