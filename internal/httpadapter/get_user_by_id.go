package httpadapter

import (
	"aura/internal/pkg/response"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (a *Adapter) GetUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	result, err := a.userService.GetUserByID(c.Request().Context(), uint(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, err.Error())
		}
		return response.InternalServerError(c, err.Error())
	}

	return response.OK(c, result)
}
