package handler

import (
	"aura/auraapi"
	"aura/internal/model"
	"context"
)

func (s *PostService) AddPost(ctx context.Context, req *auraapi.AddPostReq) (*auraapi.AddPostRes, error) {
	post, err := s.PostStorage.Insert(ctx, &model.Post{
		Content: req.Content,
		UserID:  req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &auraapi.AddPostRes{
		Post: post.ToDomain(),
	}, nil
}
