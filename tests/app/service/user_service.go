package service

import (
	"aura/auraapi"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) AddUser(ctx context.Context, req *auraapi.AddUserReq) (*auraapi.AddUserRes, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*auraapi.AddUserRes), args.Error(1)
}

func (m *MockUserService) GetUserByID(ctx context.Context, id uint) (*auraapi.GetUserByIdRes, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*auraapi.GetUserByIdRes), args.Error(1)
}

func (m *MockUserService) Login(ctx context.Context, req *auraapi.LoginReq) (*auraapi.LoginRes, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*auraapi.LoginRes), args.Error(1)
}
