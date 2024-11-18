package handler

import (
	"aura/internal/model"
	"context"
	"time"
)

func (s *UserService) SaveRefreshToken(ctx context.Context, token string, userID uint) error {
	return s.UserStorage.SaveRefreshToken(ctx, &model.RefreshToken{
		Token:     token,
		UserID:    userID,
		ExpiresIn: time.Now().Add(time.Duration(s.cfg.JWT.RefreshTokenExpiresIn) * time.Second),
	})
}
