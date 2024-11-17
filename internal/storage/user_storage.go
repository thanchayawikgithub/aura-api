package storage

import (
	"aura/internal/model"
	"context"

	"github.com/stretchr/testify/mock"
)

type (
	IUserStorage interface {
		IStorage[*model.User]
		FindByEmail(ctx context.Context, email string) (result *model.User, err error)
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

func (s *UserStorage) FindByEmail(ctx context.Context, email string) (result *model.User, err error) {
	err = s.db.WithContext(ctx).Where("email = ?", email).First(&result).Error
	return result, err
}
