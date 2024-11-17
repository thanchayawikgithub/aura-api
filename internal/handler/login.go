package handler

import (
	"aura/auraapi"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Login(ctx context.Context, req *auraapi.LoginReq) (*auraapi.LoginRes, error) {
	user, err := s.UserStorage.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &auraapi.LoginRes{
		UserID: user.ID,
		Email:  user.Email,
	}, nil
}
