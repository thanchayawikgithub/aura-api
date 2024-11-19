package model

import (
	"aura/auradomain"
	"time"

	"gorm.io/gorm"
)

const (
	CommentTableName = "comment"
)

type Comment struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Content   string         `gorm:"column:content"`
	UserID    uint
	User      User
	PostID    uint
	Post      Post
}

func (c *Comment) ToDomain() *auradomain.Comment {
	return &auradomain.Comment{
		ID:        c.ID,
		UserID:    c.UserID,
		PostID:    c.PostID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
}

func ToCommentList(comments []Comment) []*auradomain.Comment {
	result := make([]*auradomain.Comment, len(comments))
	for i, comment := range comments {
		result[i] = comment.ToDomain()
	}
	return result
}

func (Comment) TableName() string {
	return CommentTableName
}
