package model

import "gorm.io/gorm"

const (
	UserTableName = "user"
)

type User struct {
	gorm.Model
	Email    string `gorm:"column:email;unique"`
	Password string `gorm:"column:password"`
	Posts    []Post
}

func (User) TableName() string {
	return UserTableName
}
