package handler

import (
	"aura/auraapi"
	"aura/internal/model"
	"aura/internal/util"
	"context"
)

func (s *PostService) AddPost(ctx context.Context, req *auraapi.AddPostReq) (*auraapi.AddPostRes, error) {
	userID, err := util.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	post, err := s.PostStorage.Save(ctx, &model.Post{
		Content: req.Content,
		UserID:  userID,
	})
	if err != nil {
		return nil, err
	}

	return &auraapi.AddPostRes{
		Post: post.ToDomain(),
	}, nil
}
