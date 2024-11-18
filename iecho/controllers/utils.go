package controllers

import (
	"github.com/labstack/echo/v4"
	"iecho/models"
	"net/http"
)

func abort(err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

func abortWithCode(code int, err error) error {
	return echo.NewHTTPError(code, err.Error())
}

func writeSuccess(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, models.OkResponse(data))
}
