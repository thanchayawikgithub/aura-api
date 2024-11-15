package httpadapter

import (
	"aura/internal/config"
	"aura/internal/handler"
	"aura/internal/handler/itf"

	"github.com/labstack/echo/v4"
)

type Adapter struct {
	config      *config.Config
	service     *handler.Service
	userService itf.IUserService
	postService itf.IPostService
}

func New(s *handler.Service, cfg *config.Config) *Adapter {
	return &Adapter{
		config:      cfg,
		service:     s,
		userService: handler.NewUserService(s),
		postService: handler.NewPostService(s),
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
