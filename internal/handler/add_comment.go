package handler

import (
	"aura/auraapi"
	"aura/internal/model"
	"context"
)

func (s *CommentService) AddComment(ctx context.Context, req *auraapi.AddCommentReq) (*auraapi.AddCommentRes, error) {

	result, err := s.CommentStorage.Save(ctx, &model.Comment{
		UserID:  req.UserID,
		PostID:  req.PostID,
		Content: req.Content,
	})

	if err != nil {
		return nil, err
	}

	return &auraapi.AddCommentRes{
		Comment: result.ToDomain(),
	}, nil
}
