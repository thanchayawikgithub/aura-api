package itf

import (
	"aura/auraapi"
	"context"
)

type ICommentService interface {
	AddComment(ctx context.Context, req *auraapi.AddCommentReq) (*auraapi.AddCommentRes, error)
}
