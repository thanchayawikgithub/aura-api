package database

import "aura/internal/model"

type MockCommentStorage struct {
	AbstractStorage[*model.Comment]
}
