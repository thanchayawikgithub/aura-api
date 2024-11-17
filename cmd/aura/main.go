package main

import (
	"aura/internal/config"
	"aura/internal/handler"
	"aura/internal/httpadapter"
	"aura/internal/storage"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	pathParamUserID = "/:user_id"
	pathParamPostID = "/:post_id"
)

func main() {
	cfg := config.LoadConfig()

	storage := storage.New(&cfg.Database)
	service := handler.New(storage, cfg)
	adapter := httpadapter.New(service, cfg)

	e := echo.New()
	e.Validator = httpadapter.NewCustomValidator()
	setupMiddleware(e)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	// Version 1
	v1 := e.Group("/v1")

	user := v1.Group("/user")
	setUpUser(user, adapter)

	post := v1.Group("/post")
	setUpPost(post, adapter)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Server.Port)))
}

func setupMiddleware(e *echo.Echo) {
	e.Use(
		middleware.Recover(),
		middleware.RequestID(),
		middleware.Logger(),
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: 30 * time.Second,
		}),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentType,
				echo.HeaderAccept,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				http.MethodPatch,
			},
		}),
	)
}

func setUpUser(e *echo.Group, adapter *httpadapter.Adapter) {
	e.POST("", adapter.AddUser)
	e.GET(pathParamUserID, adapter.GetUserByID)
}

func setUpPost(e *echo.Group, adapter *httpadapter.Adapter) {
	e.POST("", adapter.AddPost)
	e.GET(pathParamPostID, adapter.GetPostByID)
	e.GET("/user"+pathParamUserID, adapter.GetPostsByUserID)
}
