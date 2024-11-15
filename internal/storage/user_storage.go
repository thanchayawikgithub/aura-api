package storage

import (
	"aura/internal/model"
)

type (
	IUserStorage interface {
		IStorage[*model.User]
	}

	UserStorage struct {
		AbstractStorage[*model.User]
	}
)

func NewUserStorage(s *Storage) *UserStorage {
	return &UserStorage{
		AbstractStorage: AbstractStorage[*model.User]{
			db:        s.db,
			tableName: model.UserTableName,
		},
	}
}
