package routers

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"igin/models"
	"net/http"
	"strings"
)

func notFoundHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.ErrorResponseWithCode(http.StatusNotFound, "page not found"))
	}
}

func corsHandler() gin.HandlerFunc {
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowHeaders:    []string{"Authorization", "content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host"},
	}

	return cors.New(corsConf)
}

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				switch e.Type {
				case gin.ErrorTypeBind:
					var errs validator.ValidationErrors
					ok := errors.As(e.Err, &errs)

					if !ok {
						writeError(c, e.Error())
						return
					}

					var stringErrors []string
					for _, err := range errs {
						stringErrors = append(stringErrors, translate(err))
					}
					writeError(c, strings.Join(stringErrors, "; "))
				default:
					writeError(c, e.Err.Error())
				}
			}
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

func writeError(ctx *gin.Context, errString string) {
	status := ctx.Writer.Status()
	ctx.JSON(status, models.ErrorResponseWithCode(status, errString))
}
