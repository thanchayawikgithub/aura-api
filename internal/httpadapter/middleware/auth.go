package middleware

import (
	"aura/internal/config"
	"aura/internal/pkg/auth"
	"aura/internal/pkg/response"
	"aura/internal/util"
	"context"
	"log"

	"github.com/labstack/echo/v4"
)

func Auth(config *config.JWT) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie(auth.AccessTokenCookieName)
			if err != nil {
				return response.Unauthorized(c, "Missing authentication token")
			}

			claims, err := auth.ValidateToken(cookie.Value, config)
			if err != nil {
				log.Println(err)
				return response.Unauthorized(c, err.Error())
			}

			setContextValue(c, claims)

			return next(c)
		}
	}
}

func setContextValue(c echo.Context, claims *auth.Claims) {
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, util.UserID, claims.UserID)
	ctx = context.WithValue(ctx, util.UserEmail, claims.Email)
	c.SetRequest(c.Request().WithContext(ctx))
}
