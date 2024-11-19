package database

import (
	"aura/internal/model"
	"context"
)

type MockPostStorage struct {
	AbstractStorage[*model.Post]
}

func (m *MockPostStorage) FindByUserID(ctx context.Context, userID uint) (result []*model.Post, err error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*model.Post), args.Error(1)
}
