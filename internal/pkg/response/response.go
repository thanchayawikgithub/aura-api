package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Error(c echo.Context, code int, message string) error {
	return c.JSON(code, &Response{
		Message: message,
		Data:    nil,
	})
}

func Success(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, &Response{
		Message: message,
		Data:    data,
	})
}

func OK(c echo.Context, data interface{}) error {
	return Success(c, http.StatusOK, "Success", data)
}

func Created(c echo.Context, data interface{}) error {
	return Success(c, http.StatusCreated, "Created", data)
}

func BadRequest(c echo.Context, message string) error {
	return Error(c, http.StatusBadRequest, message)
}

func Unauthorized(c echo.Context, message string) error {
	return Error(c, http.StatusUnauthorized, message)
}

func Forbidden(c echo.Context, message string) error {
	return Error(c, http.StatusForbidden, message)
}

func NotFound(c echo.Context, message string) error {
	return Error(c, http.StatusNotFound, message)
}

func InternalServerError(c echo.Context, message string) error {
	if message == "" {
		message = "Internal Server Error"
	}
	return Error(c, http.StatusInternalServerError, message)
}