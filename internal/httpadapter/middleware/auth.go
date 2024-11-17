package middleware

import (
	"aura/internal/config"
	"aura/internal/pkg/auth"
	"aura/internal/pkg/response"
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

			// Set user info in context
			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)

			return next(c)
		}
	}
}
