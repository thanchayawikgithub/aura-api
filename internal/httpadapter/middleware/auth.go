package middleware

import (
	"aura/internal/config"
	"aura/internal/handler"

	"aura/internal/pkg/auth"
	"aura/internal/pkg/response"
	"aura/internal/util"
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Auth(config *config.JWT, refreshTokenService *handler.RefreshTokenService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessTokenCookie, err := c.Cookie(auth.AccessTokenCookieName)
			if err != nil {
				return response.Unauthorized(c, "Missing authentication token")
			}

			claims, err := auth.ValidateToken(accessTokenCookie.Value, config)
			if err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					return refreshToken(c, next, config, refreshTokenService)
				}

				return response.Unauthorized(c, "Invalid token")
			}

			setClaimsToContext(c, claims)
			return next(c)
		}
	}
}

func refreshToken(c echo.Context, next echo.HandlerFunc, config *config.JWT, refreshTokenService *handler.RefreshTokenService) error {
	ctx := c.Request().Context()
	refreshTokenCookie, err := c.Cookie(auth.RefreshTokenCookieName)
	if err != nil {
		return response.Unauthorized(c, "Missing refresh token")
	}

	refreshTokenClaims, err := auth.ValidateToken(refreshTokenCookie.Value, config)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return response.Unauthorized(c, "Refresh token expired, please login again")
		}
		return response.Unauthorized(c, "Invalid refresh token")
	}

	if refreshTokenClaims.TokenType != "refresh" {
		return response.Unauthorized(c, "Invalid token type")
	}

	refreshToken, err := refreshTokenService.GetByToken(ctx, refreshTokenCookie.Value)
	if err != nil {
		return response.Unauthorized(c, "Refresh token not found")
	}

	if refreshToken.IsRevoked {
		return response.Unauthorized(c, "Refresh token revoked")
	}

	newAccessToken, err := auth.GenerateAccessToken(refreshTokenClaims.UserID, refreshTokenClaims.Email, config)
	if err != nil {
		return response.InternalServerError(c, "Failed to generate access token")
	}

	newAccessTokenClaims, err := auth.ValidateToken(newAccessToken, config)
	if err != nil {
		return response.InternalServerError(c, "Failed to validate access token")
	}

	newRefreshToken, err := auth.GenerateRefreshToken(refreshTokenClaims.UserID, refreshTokenClaims.Email, config, refreshTokenClaims)
	if err != nil {
		return response.InternalServerError(c, "Failed to generate refresh token")
	}

	err = refreshTokenService.Rotate(ctx, refreshToken, newRefreshToken)
	if err != nil {
		return response.InternalServerError(c, "Failed to rotate refresh token")
	}

	util.SetSecureCookie(c, auth.AccessTokenCookieName, newAccessToken)
	util.SetSecureCookie(c, auth.RefreshTokenCookieName, newRefreshToken)
	setClaimsToContext(c, newAccessTokenClaims)

	return next(c)
}

func setClaimsToContext(c echo.Context, claims *auth.Claims) {
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, util.UserID, claims.UserID)
	ctx = context.WithValue(ctx, util.UserEmail, claims.Email)
	c.SetRequest(c.Request().WithContext(ctx))
}
