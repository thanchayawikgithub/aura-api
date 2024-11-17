package handler

import (
	"aura/auraapi"
	"aura/internal/model"
	"aura/internal/util"
	"context"
)

func (s *PostService) EditPost(ctx context.Context, req *auraapi.EditPostReq, postID uint) (*auraapi.EditPostRes, error) {
	userID, err := util.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	post, err := s.PostStorage.FindByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	if post.UserID != userID {
		return nil, ErrNoPermission
	}

	post, err = s.PostStorage.Update(ctx, postID, &model.Post{
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &auraapi.EditPostRes{
		Post: post.ToDomain(),
	}, nil
}
