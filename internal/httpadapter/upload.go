package httpadapter

import (
	"aura/internal/pkg/response"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.BadRequest(c, "File is required")
	}

	err = a.attachmentService.UploadFile(c.Request().Context(), file)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.OK(c, nil)
}
