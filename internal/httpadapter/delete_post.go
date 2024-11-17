package httpadapter

import (
	"aura/internal/handler"
	"aura/internal/pkg/response"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) DeletePost(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := a.postService.DeletePost(c.Request().Context(), uint(postID)); err != nil {
		if errors.Is(err, handler.ErrNoPermission) {
			return response.Forbidden(c, err.Error())
		}

		return response.InternalServerError(c, err.Error())
	}

	return response.NoContent(c)
}
