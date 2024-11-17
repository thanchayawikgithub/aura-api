package auraapi

import "aura/auradomain"

type (
	EditPostReq struct {
		Content string `json:"content" validate:"required"`
	}

	EditPostRes struct {
		*auradomain.Post
	}
)
