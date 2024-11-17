package handler

import (
	"aura/auraapi"
	"context"
)

func (s *PostService) GetPostByID(ctx context.Context, id uint) (*auraapi.GetPostByIdRes, error) {
	post, err := s.PostStorage.WithPreload("User").FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &auraapi.GetPostByIdRes{
		Post: post.ToDomain(),
		User: post.User.ToDomain(),
	}, nil
}
