package handler

import (
	"aura/internal/config"
	"aura/internal/storage"
)

type (
	Service struct {
		cfg         *config.Config
		UserStorage storage.IUserStorage
		PostStorage storage.IPostStorage
	}

	UserService struct {
		*Service
	}

	PostService struct {
		*Service
	}
)

func New(s *storage.Storage, cfg *config.Config) *Service {
	return &Service{
		cfg:         cfg,
		UserStorage: storage.NewUserStorage(s),
		PostStorage: storage.NewPostStorage(s),
	}
}

func NewUserService(service *Service) *UserService {
	return &UserService{
		Service: service,
	}
}

func NewPostService(service *Service) *PostService {
	return &PostService{
		Service: service,
	}
}
