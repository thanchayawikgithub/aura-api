package httpadapter

import (
	"aura/internal/pkg/response"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (a *Adapter) GetCommentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}
	res, err := a.commentService.GetCommentByID(c.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, err.Error())
		}
		return response.InternalServerError(c, err.Error())
	}
	return response.OK(c, res)
}
