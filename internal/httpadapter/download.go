package httpadapter

import "github.com/labstack/echo/v4"

func (a *Adapter) DownloadFile(c echo.Context) error {
	return a.attachmentService.DownloadFile(c.Request().Context(), c.QueryParam("path"))
}
