package auraapi

type (
	LoginReq struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginRes struct {
		UserID uint   `json:"user_id"`
		Email  string `json:"email"`
	}
)
