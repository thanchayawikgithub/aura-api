package main

import (
	"aura/internal/config"
	"aura/internal/handler"
	"aura/internal/httpadapter"
	"aura/internal/storage"
	"fmt"
	"net/http"
	"time"

	mdw "aura/internal/httpadapter/middleware"

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

	mdwAuth := mdw.Auth(&cfg.JWT, handler.NewRefreshTokenService(service))
	mdws := []echo.MiddlewareFunc{
		mdw.WithTx(storage),
	}

	// Version 1
	v1 := e.Group("/v1")

	auth := v1.Group("/auth", mdws...)
	setUpAuth(auth, adapter)

	user := v1.Group("/user", mdws...)
	setUpUser(user, adapter, mdwAuth)

	post := v1.Group("/post", mdws...)
	setUpPost(post, adapter, mdwAuth)

	comment := v1.Group("/comment", mdws...)
	setUpComment(comment, adapter)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Server.Port)))
}

func setUpAuth(e *echo.Group, adapter *httpadapter.Adapter) {
	e.POST("/login", adapter.Login)
	e.POST("/logout", adapter.Logout)
}

func setUpUser(e *echo.Group, adapter *httpadapter.Adapter, mdwAuth echo.MiddlewareFunc) {
	e.POST("", adapter.AddUser)
	e.GET(pathParamUserID, adapter.GetUserByID, mdwAuth)
}

func setUpPost(e *echo.Group, adapter *httpadapter.Adapter, mdwAuth echo.MiddlewareFunc) {
	e.POST("", adapter.AddPost)
	e.GET(pathParamPostID, adapter.GetPostByID, mdwAuth)
	e.GET("/user"+pathParamUserID, adapter.GetPostsByUserID, mdwAuth)
	e.PATCH(pathParamPostID, adapter.EditPost, mdwAuth)
	e.DELETE(pathParamPostID, adapter.DeletePost, mdwAuth)
}

func setUpComment(e *echo.Group, adapter *httpadapter.Adapter) {
	e.POST("", adapter.AddComment)
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
