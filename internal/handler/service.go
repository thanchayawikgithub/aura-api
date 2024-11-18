package handler

import (
	"aura/internal/config"
	"aura/internal/storage"
)

type (
	Service struct {
		cfg                 *config.Config
		UserStorage         storage.IUserStorage
		PostStorage         storage.IPostStorage
		RefreshTokenStorage storage.IRefreshTokenStorage
		CommentStorage      storage.ICommentStorage
	}

	UserService struct {
		*Service
	}

	PostService struct {
		*Service
	}

	RefreshTokenService struct {
		*Service
	}

	CommentService struct {
		*Service
	}
)

func New(s *storage.Storage, cfg *config.Config) *Service {
	return &Service{
		cfg:                 cfg,
		UserStorage:         storage.NewUserStorage(s),
		PostStorage:         storage.NewPostStorage(s),
		RefreshTokenStorage: storage.NewRefreshTokenStorage(s),
		CommentStorage:      storage.NewCommentStorage(s),
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

func NewRefreshTokenService(service *Service) *RefreshTokenService {
	return &RefreshTokenService{
		Service: service,
	}
}

func NewCommentService(service *Service) *CommentService {
	return &CommentService{
		Service: service,
	}
}
