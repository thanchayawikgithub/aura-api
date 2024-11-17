package storage

import (
	"aura/internal/model"

	"github.com/stretchr/testify/mock"
)

type (
	IUserStorage interface {
		IStorage[*model.User]
	}

	UserStorage struct {
		AbstractStorage[*model.User]
	}

	MockUserStorage struct {
		mock.Mock
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

func NewMockUserStorage() *MockUserStorage {
	return &MockUserStorage{}
}
