package storage

import (
	"aura/internal/model"
)

type (
	ICommentStorage interface {
		IStorage[*model.Comment]
	}

	CommentStorage struct {
		AbstractStorage[*model.Comment]
	}
)

func NewCommentStorage(s *Storage) *CommentStorage {
	return &CommentStorage{
		AbstractStorage: AbstractStorage[*model.Comment]{
			db:        s.db,
			tableName: model.CommentTableName,
		},
	}
}
