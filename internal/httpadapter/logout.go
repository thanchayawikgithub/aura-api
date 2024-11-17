package httpadapter

import (
	"aura/internal/pkg/auth"
	"aura/internal/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     auth.AccessTokenCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})

	return response.OK(c, nil)
}
