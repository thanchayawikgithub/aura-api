package model

import "gorm.io/gorm"

const (
	PostTableName = "post"
)

type Post struct {
	gorm.Model
	Content string `gorm:"column:content"`
	UserID  uint
	User    User
}

func (Post) TableName() string {
	return PostTableName
}
