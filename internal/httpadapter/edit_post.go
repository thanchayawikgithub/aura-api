package httpadapter

import (
	"aura/auraapi"
	"aura/internal/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) EditPost(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	req, err := BindAndValidate[auraapi.EditPostReq](c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	result, err := a.postService.EditPost(c.Request().Context(), &auraapi.EditPostReq{
		Content: req.Content,
	}, uint(postID))
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.OK(c, result)
}
