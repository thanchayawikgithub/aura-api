package httpadapter

import (
	"aura/internal/pkg/export"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) ExportUsers(c echo.Context) error {
	err := a.userService.ExportUsers(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.File(export.ExportUsersPath)
}
