package routers

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"iecho/models"
	"net/http"
	"strings"
)

type iValidator struct {
	validator *validator.Validate
}

func (cv *iValidator) Validate(md interface{}) error {
	return cv.validator.Struct(md)
}

func cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowHeaders: []string{"Authorization", "content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host"},
	})
}

func errorHandler() echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		if ctx.Response().Committed {
			return
		}

		var code int
		var message string
		var ve validator.ValidationErrors
		var be *echo.BindingError
		var he *echo.HTTPError
		switch {
		case errors.As(err, &ve):
			var stringErrors []string
			for _, e := range ve {
				stringErrors = append(stringErrors, translate(e))
			}
			code = http.StatusBadRequest
			message = strings.Join(stringErrors, "; ")
		case errors.As(err, &he):
			code = he.Code
			switch m := he.Message.(type) {
			case string:
				message = m
			}
		case errors.As(err, &be):
			code = be.Code
			switch m := be.Message.(type) {
			case string:
				message = m
			}
		default:
			code = http.StatusInternalServerError
			message = err.Error()
		}

		if message == "" {
			message = http.StatusText(code)
		}

		response := models.ErrorResponseWithCode(code, message)
		if err = ctx.JSON(response.Code, response); err != nil {
			ctx.Logger().Error(err)
		}
	}
}

func translate(e validator.FieldError) string {
	field := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("Field '%s' is required", field)
	case "max":
		return fmt.Sprintf("Field '%s' must be less or equal to %s", field, e.Param())
	case "min":
		return fmt.Sprintf("Field '%s' must be more or equal to %s", field, e.Param())
	}
	return e.Error()
}
