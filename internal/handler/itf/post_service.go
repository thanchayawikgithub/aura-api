package itf

import (
	"aura/auraapi"
	"context"
)

type IPostService interface {
	AddPost(ctx context.Context, req *auraapi.AddPostReq) (*auraapi.AddPostRes, error)
	GetPostByID(ctx context.Context, id uint) (*auraapi.GetPostByIdRes, error)
}
