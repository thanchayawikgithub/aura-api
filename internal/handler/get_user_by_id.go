package handler

import (
	"aura/auraapi"
	"context"
)

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*auraapi.GetUserByIdRes, error) {
	user, err := s.UserStorage.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &auraapi.GetUserByIdRes{
		User: user.ToDomain(),
	}, nil
}
