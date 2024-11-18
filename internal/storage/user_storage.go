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
		SaveRefreshToken(ctx context.Context, refreshToken *model.RefreshToken) error
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

func (s *UserStorage) SaveRefreshToken(ctx context.Context, refreshToken *model.RefreshToken) error {
	return s.db.WithContext(ctx).Model(&model.RefreshToken{}).Save(&refreshToken).Error
}
