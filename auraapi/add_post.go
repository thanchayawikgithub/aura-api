package auraapi

import "aura/auradomain"

type (
	AddPostReq struct {
		Content string `json:"content" validate:"required"`
		UserID  uint   `json:"user_id" validate:"required"`
	}

	AddPostRes struct {
		*auradomain.Post
	}
)
