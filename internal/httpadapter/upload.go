package httpadapter

import (
	"aura/internal/pkg/response"
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.BadRequest(c, "File is required")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err := os.MkdirAll("uploads", 0755); err != nil {
		return err
	}

	// Destination
	dst, err := os.Create("uploads/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return response.OK(c, nil)
}
