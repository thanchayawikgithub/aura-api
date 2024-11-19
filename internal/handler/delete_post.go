package handler

import (
	"aura/internal/model"
	"aura/internal/util"
	"context"
)

func (s *PostService) DeletePost(ctx context.Context, postID uint) error {
	userID, err := util.GetUserID(ctx)
	if err != nil {
		return err
	}

	post, err := s.PostStorage.FindByID(ctx, postID)
	if err != nil {
		return err
	}

	if post.UserID != userID {
		return ErrNoPermission
	}

	return s.PostStorage.Delete(ctx, &model.Post{ID: postID})
}
