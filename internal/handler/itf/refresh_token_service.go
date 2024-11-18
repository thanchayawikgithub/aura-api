package itf

import (
	"aura/internal/model"
	"context"
)

type IRefreshTokenService interface {
	GetByToken(ctx context.Context, token string) (*model.RefreshToken, error)
	Save(ctx context.Context, token string, userID uint) error
	Rotate(ctx context.Context, oldToken *model.RefreshToken, newToken string) error
}
