package handler

import (
	"aura/auraapi"
	"aura/internal/model"
	"context"
)

func (s *PostService) EditPost(ctx context.Context, req *auraapi.EditPostReq, postID uint) (*auraapi.EditPostRes, error) {
	post, err := s.PostStorage.Update(ctx, postID, &model.Post{
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &auraapi.EditPostRes{
		Post: post.ToDomain(),
	}, nil
}
