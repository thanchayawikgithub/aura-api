package model

import (
	"aura/auradomain"

	"gorm.io/gorm"
)

const (
	PostTableName = "post"
)

type Post struct {
	gorm.Model
	Content string `gorm:"column:content"`
	UserID  uint
	User    User
}

func (p *Post) ToDomain() *auradomain.Post {
	return &auradomain.Post{
		ID:        p.ID,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
	}
}

func ToPostList(posts []*Post) []*auradomain.Post {
	result := make([]*auradomain.Post, len(posts))
	for i, post := range posts {
		result[i] = post.ToDomain()
	}
	return result
}

func (Post) TableName() string {
	return PostTableName
}
