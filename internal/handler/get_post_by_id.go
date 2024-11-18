package handler

import (
	"aura/auraapi"
	"aura/internal/model"
	"context"
)

func (s *PostService) GetPostByID(ctx context.Context, id uint) (*auraapi.GetPostByIdRes, error) {
	post, err := s.PostStorage.WithPreload("User", "Comments").FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &auraapi.GetPostByIdRes{
		Post:     post.ToDomain(),
		User:     post.User.ToDomain(),
		Comments: model.ToCommentList(post.Comments),
	}, nil
}
