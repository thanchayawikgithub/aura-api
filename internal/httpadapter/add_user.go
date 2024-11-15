package httpadapter

import (
	"aura/auraapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) AddUser(c echo.Context) error {
	req, err := BindAndValidate[auraapi.AddUserReq](c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := a.userService.AddUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
