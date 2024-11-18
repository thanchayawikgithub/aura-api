package model

import (
	"time"

	"gorm.io/gorm"
)

const RefreshTokenTableName = "refresh_token"

type RefreshToken struct {
	gorm.Model
	Token     string    `gorm:"column:token;not null"`
	UserID    uint      `gorm:"column:user_id;not null"`
	ExpiresIn time.Time `gorm:"column:expires_in;not null"`
	IsRevoked bool      `gorm:"column:is_revoked;default:false"`
}

func (RefreshToken) TableName() string {
	return RefreshTokenTableName
}
