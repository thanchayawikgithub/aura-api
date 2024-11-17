package httpadapter

import (
	"aura/auraapi"
	"aura/internal/pkg/auth"
	"aura/internal/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) Login(c echo.Context) error {
	req, err := BindAndValidate[auraapi.LoginReq](c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	result, err := a.userService.Login(c.Request().Context(), req)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}

	// Generate JWT token
	token, err := auth.GenerateToken(result.UserID, result.Email, &a.config.JWT)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	// Set cookie
	c.SetCookie(&http.Cookie{
		Name:     auth.AccessTokenCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	return response.OK(c, result)
}
