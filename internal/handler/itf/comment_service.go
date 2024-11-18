package itf

import (
	"aura/auraapi"
	"context"
)

type ICommentService interface {
	AddComment(ctx context.Context, req *auraapi.AddCommentReq) (*auraapi.AddCommentRes, error)
	GetCommentByID(ctx context.Context, id uint) (*auraapi.GetCommentByIdRes, error)
	DeleteComment(ctx context.Context, id uint) error
}
