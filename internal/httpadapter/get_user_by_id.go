package httpadapter

import (
	"aura/internal/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) GetUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	result, err := a.userService.GetUserByID(c.Request().Context(), uint(userID))
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.OK(c, result)
}
