package httpadapter

import (
	"aura/auraapi"
	"aura/internal/pkg/auth"
	"aura/internal/pkg/response"
	"aura/internal/util"

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

	refreshToken, err := auth.GenerateRefreshToken(result.UserID, result.Email, &a.config.JWT, nil)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	// Save refresh token
	err = a.refreshTokenService.Save(ctx, refreshToken, result.UserID)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	util.SetSecureCookie(c, auth.AccessTokenCookieName, accessToken)
	util.SetSecureCookie(c, auth.RefreshTokenCookieName, refreshToken)

	return response.OK(c, result)
}
