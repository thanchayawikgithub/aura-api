package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetSecureCookie(c echo.Context, name, value string) {
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})
}
