package auraapi

import "aura/auradomain"

type (
	AddCommentReq struct {
		Content string `json:"content" validate:"required"`
		UserID  uint   `json:"user_id" validate:"required"`
		PostID  uint   `json:"post_id" validate:"required"`
	}

	AddCommentRes struct {
		*auradomain.Comment
	}
)
