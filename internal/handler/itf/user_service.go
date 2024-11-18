package itf

import (
	"aura/auraapi"
	"context"
)

type IUserService interface {
	AddUser(ctx context.Context, req *auraapi.AddUserReq) (*auraapi.AddUserRes, error)
	GetUserByID(ctx context.Context, id uint) (*auraapi.GetUserByIdRes, error)
	Login(ctx context.Context, req *auraapi.LoginReq) (*auraapi.LoginRes, error)
	SaveRefreshToken(ctx context.Context, token string, userID uint) error

	// IsRefreshTokenRevoked(ctx context.Context, token string) (bool, error)
	// RotateRefreshToken(ctx context.Context, oldToken, newToken string, userID uint) error
}
