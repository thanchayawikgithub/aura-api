package handler

import (
	"aura/auraapi"
	"context"
)

func (s *CommentService) GetCommentByID(ctx context.Context, id uint) (*auraapi.GetCommentByIdRes, error) {
	result, err := s.CommentStorage.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &auraapi.GetCommentByIdRes{Comment: result.ToDomain()}, nil
}
