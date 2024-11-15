package handler

import (
	"aura/internal/config"
	"aura/internal/storage"
)

type (
	Service struct {
		UserStorage storage.IUserStorage
	}

	UserService struct {
		*Service
	}
)

func New(s *storage.Storage, cfg *config.Config) *Service {
	return &Service{
		UserStorage: storage.NewUserStorage(s),
	}
}

func NewUserService(service *Service) *UserService {
	return &UserService{
		Service: service,
	}
}
