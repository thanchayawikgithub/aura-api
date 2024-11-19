package database

import (
	"aura/internal/model"
	"context"
)

type MockUserStorage struct {
	AbstractStorage[*model.User]
}

func (m *MockUserStorage) FindByEmail(ctx context.Context, email string) (result *model.User, err error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*model.User), args.Error(1)
}
