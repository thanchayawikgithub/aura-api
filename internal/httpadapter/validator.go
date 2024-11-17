package httpadapter

import (
	"aura/internal/pkg/exception"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	err := v.RegisterValidation("name", nameValidator)
	if err != nil {
		return &CustomValidator{
			Validator: validator.New(),
		}
	}

	return &CustomValidator{
		Validator: v,
	}
}

func (v *CustomValidator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		if validatationErrors, ok := err.(validator.ValidationErrors); ok {
			msg := []string{}
			for _, v := range validatationErrors {
				msg = append(msg, fmt.Sprintf("Key: %s Field: %s Error: %s", v.Namespace(), v.Field(), v.Tag()))
			}

			return &exception.ValidateError{
				Message: strings.Join(msg, ", "),
			}
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func nameValidator(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	return regexp.MustCompile(`^[a-zA-Z0-9_/\\. -]+$`).MatchString(name)
}
