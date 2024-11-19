package model

import (
	"aura/auradomain"
	"time"

	"gorm.io/gorm"
)

const (
	UserTableName = "user"
)

type User struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Email       string         `gorm:"column:email;unique"`
	Username    string         `gorm:"column:username;unique"`
	DisplayName string         `gorm:"column:display_name"`
	Password    string         `gorm:"column:password"`
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
