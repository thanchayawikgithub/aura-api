package model

import (
	"aura/auradomain"

	"gorm.io/gorm"
)

const (
	UserTableName = "user"
)

type User struct {
	gorm.Model
	Email       string `gorm:"column:email;unique" validate:"required,email"`
	Username    string `gorm:"column:username;unique" validate:"required"`
	DisplayName string `gorm:"column:display_name" validate:"required"`
	Password    string `gorm:"column:password" validate:"required"`
	Posts       []Post
	Comments    []Comment
}

func (User) TableName() string {
	return UserTableName
}

func (u *User) ToDomain() *auradomain.User {
	return &auradomain.User{
		ID:          u.ID,
		Email:       u.Email,
		Username:    u.Username,
		DisplayName: u.DisplayName,
	}
}
