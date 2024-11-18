package httpadapter

import (
	"aura/internal/pkg/response"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (a *Adapter) GetPostByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	result, err := a.postService.GetPostByID(c.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, err.Error())
		}
		return response.InternalServerError(c, err.Error())
	}

	return response.OK(c, result)
}
