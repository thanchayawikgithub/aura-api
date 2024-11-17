package httpadapter

import (
	"aura/internal/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) GetPostByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	result, err := a.postService.GetPostByID(c.Request().Context(), uint(id))
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.OK(c, result)
}
