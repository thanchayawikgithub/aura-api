package auraapi

import "aura/auradomain"

type (
	GetPostByIdRes struct {
		*auradomain.Post
		User *auradomain.User `json:"user"`
	}
)
