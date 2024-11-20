package auraapi

import "aura/auradomain"

type (
	AddPostReq struct {
		Content string `json:"content" validate:"required"`
	}

	AddPostRes struct {
		*auradomain.Post
	}
)
