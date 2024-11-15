package auraapi

import "aura/auradomain"

type (
	AddUserReq struct {
		Email       string `json:"email" validate:"required,email"`
		Password    string `json:"password" validate:"required"`
		Username    string `json:"username" validate:"required"`
		DisplayName string `json:"display_name" validate:"required"`
	}

	AddUserRes struct {
		*auradomain.User
	}
)
