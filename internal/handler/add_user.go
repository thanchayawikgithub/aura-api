package handler

import (
	"aura/auraapi"
	"aura/internal/model"
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) AddUser(ctx context.Context, req *auraapi.AddUserReq) (*auraapi.AddUserRes, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := s.UserStorage.Save(ctx, &model.User{
		Email:       req.Email,
		Username:    req.Username,
		DisplayName: req.DisplayName,
		Password:    string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	return &auraapi.AddUserRes{
		User: user.ToDomain(),
	}, nil
}
