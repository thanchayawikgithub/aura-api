package storage

import (
	"aura/internal/model"
)

type (
	IUserStorage interface {
		IStorage[*model.User]
	}

	UserStorage struct {
		AbstractStorage[*model.User]
	}
)
