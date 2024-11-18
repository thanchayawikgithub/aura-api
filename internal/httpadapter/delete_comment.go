package httpadapter

import (
	"aura/internal/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) DeleteComment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	err = a.commentService.DeleteComment(c.Request().Context(), uint(id))
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}
	return response.NoContent(c)
}
