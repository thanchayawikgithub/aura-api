package httpadapter

import (
	"aura/auraapi"
	"aura/internal/pkg/response"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) AddComment(c echo.Context) error {
	req, err := BindAndValidate[auraapi.AddCommentReq](c)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	result, err := a.commentService.AddComment(c.Request().Context(), req)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.Created(c, result)
}
