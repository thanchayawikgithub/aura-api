package handler

import (
	"aura/internal/model"
	"context"

	"gorm.io/gorm"
)

func (s *PostService) DeletePost(ctx context.Context, postID uint) error {
	return s.PostStorage.Delete(ctx, &model.Post{Model: gorm.Model{ID: postID}})
}
