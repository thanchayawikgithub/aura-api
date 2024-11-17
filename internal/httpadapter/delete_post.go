package httpadapter

import (
	"aura/internal/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) DeletePost(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := a.postService.DeletePost(c.Request().Context(), uint(postID)); err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.OK(c, nil)
}
