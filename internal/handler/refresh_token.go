package handler

import (
	"aura/internal/model"
	"aura/internal/util"
	"context"
	"log"
)

func (s *RefreshTokenService) Save(ctx context.Context, token string, userID uint) error {
	_, err := s.RefreshTokenStorage.Save(ctx, &model.RefreshToken{
		Token:  token,
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *RefreshTokenService) GetByToken(ctx context.Context, token string) (*model.RefreshToken, error) {
	refreshToken, err := s.RefreshTokenStorage.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}

func (s *RefreshTokenService) Rotate(ctx context.Context, oldRefreshToken *model.RefreshToken, newRefreshToken string) error {
	tx, err := util.GetTx(ctx)

	if err != nil {
		log.Printf("Transaction not found: %v", err)
		return err
	}

	oldRefreshToken.IsRevoked = true
	_, err = s.RefreshTokenStorage.WithTx(tx).Update(ctx, oldRefreshToken.ID, oldRefreshToken)
	if err != nil {
		return err
	}
	tx.Commit()

	_, err = s.RefreshTokenStorage.WithTx(tx).Save(ctx, &model.RefreshToken{
		Token:  newRefreshToken,
		UserID: oldRefreshToken.UserID,
	})
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}
