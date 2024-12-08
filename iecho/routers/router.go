package routers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"iecho/models"
	"net/http"
)

// HandlerRegisterFunc 在根路由组下添加 api
var HandlerRegisterFunc = func(root *echo.Group) {}

func Router(engine *echo.Echo) {

	// 入参校验
	engine.Validator = &iValidator{validator: validator.New()}

	// error 处理
	engine.HTTPErrorHandler = errorHandler()
	// 404
	engine.RouteNotFound("/*", func(ctx echo.Context) error {
		response := models.ErrorResponseWithCode(http.StatusNotFound, "page not found")
		return ctx.JSON(response.Code, response)
	})

	// cors
	engine.Use(cors())

	rootGroup := engine.Group("/iecho")

	// api
	HandlerRegisterFunc(rootGroup)
}
