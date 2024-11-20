package httpadapter

import (
	"aura/internal/config"
	"aura/internal/handler"
	"aura/internal/handler/itf"

	"github.com/labstack/echo/v4"
)

type Adapter struct {
	cfg                 *config.Config
	service             *handler.Service
	userService         itf.IUserService
	postService         itf.IPostService
	refreshTokenService itf.IRefreshTokenService
	commentService      itf.ICommentService
	attachmentService   itf.IAttachmentService
}

func New(s *handler.Service, cfg *config.Config) *Adapter {
	return &Adapter{
		cfg:                 cfg,
		service:             s,
		userService:         handler.NewUserService(s),
		refreshTokenService: handler.NewRefreshTokenService(s),
		postService:         handler.NewPostService(s),
		commentService:      handler.NewCommentService(s),
		attachmentService:   handler.NewAttachmentService(s),
	}
}

func BindAndValidate[T interface{}](c echo.Context) (req *T, err error) {
	req = new(T)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	err = c.Validate(req)
	return req, err
}
