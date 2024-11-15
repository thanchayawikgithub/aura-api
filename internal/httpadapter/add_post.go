package httpadapter

import (
	"aura/auraapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) AddPost(c echo.Context) error {
	req, err := BindAndValidate[auraapi.AddPostReq](c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := a.postService.AddPost(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
