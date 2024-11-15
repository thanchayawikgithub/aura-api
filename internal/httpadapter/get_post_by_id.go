package httpadapter

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) GetPostByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := a.postService.GetPostByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
