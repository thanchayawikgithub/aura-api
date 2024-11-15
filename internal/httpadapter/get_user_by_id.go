package httpadapter

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) GetUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := a.userService.GetUserByID(c.Request().Context(), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
