package itf

import (
	"aura/auraapi"
	"context"
)

type IUserService interface {
	// GetAll(ctx context.Context) ([]*model.User, error)
	// GetByID(ctx context.Context, id uint) (*model.User, error)
	AddUser(ctx context.Context, req *auraapi.AddUserReq) (*auraapi.AddUserRes, error)
	// Update(ctx context.Context, user *model.User) error
	// Delete(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id uint) (*auraapi.GetUserByIdRes, error)
}
