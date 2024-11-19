package handler

import (
	"aura/internal/model"
	"aura/internal/util"
	"context"
)

func (s *CommentService) DeleteComment(ctx context.Context, id uint) error {
	userID, err := util.GetUserID(ctx)
	if err != nil {
		return err
	}

	comment, err := s.CommentStorage.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return ErrNoPermission
	}

	return s.CommentStorage.Delete(ctx, &model.Comment{ID: id})
}
