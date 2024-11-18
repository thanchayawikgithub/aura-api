package httpadapter

import (
	"aura/auraapi"
	"aura/internal/pkg/auth"
	"aura/internal/pkg/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req, err := BindAndValidate[auraapi.LoginReq](c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	result, err := a.userService.Login(ctx, req)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}

	// Generate JWT token
	accessToken, err := auth.GenerateAccessToken(result.UserID, result.Email, &a.config.JWT)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	refreshToken, err := auth.GenerateRefreshToken(result.UserID, result.Email, &a.config.JWT)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	// Save refresh token
	err = a.userService.SaveRefreshToken(ctx, refreshToken, result.UserID)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	setSecureCookie(c, auth.AccessTokenCookieName, accessToken,
		time.Duration(a.config.JWT.AccessTokenExpiresIn)*time.Second)
	setSecureCookie(c, auth.RefreshTokenCookieName, refreshToken,
		time.Duration(a.config.JWT.RefreshTokenExpiresIn)*time.Second)

	return response.OK(c, result)
}

func setSecureCookie(c echo.Context, name, value string, maxAge time.Duration) {
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(maxAge.Seconds()),
	})
}
