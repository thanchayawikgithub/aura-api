package exception

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HttpError represents a custom HTTP error
type HttpError struct {
	Status  int
	Code    string
	Message string
}

type ValidateError struct {
	Message string
}

// Error implements error.
func (ve *ValidateError) Error() string {
	return ve.Message
}

// Create a new HTTP error with default status code (400)
func New(ctx context.Context, code, message string) *HttpError {
	return &HttpError{
		Status:  http.StatusBadRequest,
		Code:    code,
		Message: message,
	}
}

// Create a new HTTP error with custom status code
func NewWithStatus(ctx context.Context, status int, code, message string) *HttpError {
	return &HttpError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

// Implement the error interface
func (he *HttpError) Error() string {
	return he.Message
}

// CustomHTTPErrorHandler handles all errors in the application
func CustomHTTPErrorHandler(err error, c echo.Context) {
	status := http.StatusInternalServerError
	message := "Internal Server Error"

	if httpError, ok := err.(*HttpError); ok {
		status = httpError.Status
		message = httpError.Message
	} else if echoError, ok := err.(*echo.HTTPError); ok {
		status = echoError.Code
		message = fmt.Sprintf("%v", echoError.Message)
	}

	if !c.Response().Committed {
		if err := c.JSON(status, map[string]string{
			"code":    fmt.Sprintf("%d", status),
			"message": message,
		}); err != nil {
			c.Logger().Error(err)
		}
	}
}
