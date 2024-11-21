package handler

import (
	"aura/auraapi"
	"aura/internal/model"
	"context"
)

// func (s *PostService) GetPostByID(ctx context.Context, id uint) (*auraapi.GetPostByIdRes, error) {
// 	cacheKey := fmt.Sprintf("post:%d", id)

// 	var cachedPost *auraapi.GetPostByIdRes
// 	if err := s.RedisClient.Get(ctx, cacheKey, &cachedPost); err == nil {
// 		log.Printf("return cached post: %s", cacheKey)
// 		return cachedPost, nil
// 	}

// 	post, err := s.PostStorage.WithPreload("User", "Comments").FindByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	_ = s.RedisClient.Set(ctx, cacheKey, post)

// 	return &auraapi.GetPostByIdRes{
// 		Post:     post.ToDomain(),
// 		User:     post.User.ToDomain(),
// 		Comments: model.ToCommentList(post.Comments),
// 	}, nil
// }

func (s *PostService) GetPostByID(ctx context.Context, id uint) (*auraapi.GetPostByIdRes, error) {
	post, err := s.PostStorage.WithPreload("User", "Comments").FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &auraapi.GetPostByIdRes{
		Post:     post.ToDomain(),
		User:     post.User.ToDomain(),
		Comments: model.ToCommentList(post.Comments),
	}, nil
}
