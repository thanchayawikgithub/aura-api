package handler

import (
	"aura/auraapi"
	"aura/internal/model"
	"context"
)

func (s *PostService) GetPostsByUserID(ctx context.Context, userID uint) (*auraapi.GetPostsByUserIDRes, error) {
	posts, err := s.PostStorage.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &auraapi.GetPostsByUserIDRes{
		Posts: model.ToPostList(posts),
	}, nil
}
