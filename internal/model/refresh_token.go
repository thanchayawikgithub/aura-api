package model

import (
	"time"

	"gorm.io/gorm"
)

const RefreshTokenTableName = "refresh_token"

type RefreshToken struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Token     string         `gorm:"column:token;not null"`
	UserID    uint           `gorm:"column:user_id;not null"`
	IsRevoked bool           `gorm:"column:is_revoked;default:false;type:boolean"`
}

func (RefreshToken) TableName() string {
	return RefreshTokenTableName
}
