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
	pathParamUserID    = "/:user_id"
	pathParamPostID    = "/:post_id"
	pathParamCommentID = "/:comment_id"
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
	setUpComment(comment, adapter, mdwAuth)

	attachment := v1.Group("/attachment", append(mdws, mdwAuth)...)
	setUpAttachment(attachment, adapter)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Server.Port)))
}

func setUpAuth(auth *echo.Group, adapter *httpadapter.Adapter) {
	auth.POST("/login", adapter.Login)
	auth.POST("/logout", adapter.Logout)
}

func setUpUser(user *echo.Group, adapter *httpadapter.Adapter, mdwAuth echo.MiddlewareFunc) {
	user.POST("", adapter.AddUser)
	user.GET(pathParamUserID, adapter.GetUserByID, mdwAuth)
	user.GET("/export", adapter.ExportUsers)
}

func setUpPost(post *echo.Group, adapter *httpadapter.Adapter, mdwAuth echo.MiddlewareFunc) {
	post.POST("", adapter.AddPost, mdwAuth)
	post.GET(pathParamPostID, adapter.GetPostByID, mdwAuth)
	post.GET("/user"+pathParamUserID, adapter.GetPostsByUserID, mdwAuth)
	post.PATCH(pathParamPostID, adapter.EditPost, mdwAuth)
	post.DELETE(pathParamPostID, adapter.DeletePost, mdwAuth)
}

func setUpComment(comment *echo.Group, adapter *httpadapter.Adapter, mdwAuth echo.MiddlewareFunc) {
	comment.POST("", adapter.AddComment)
	comment.GET(pathParamCommentID, adapter.GetCommentByID, mdwAuth)
	comment.DELETE(pathParamCommentID, adapter.DeleteComment, mdwAuth)
}

func setUpAttachment(attachment *echo.Group, adapter *httpadapter.Adapter) {
	attachment.POST("", adapter.UploadFile)
	attachment.GET("", adapter.DownloadFile)
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
