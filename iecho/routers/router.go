package routers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// HandlerRegisterFunc 在根路由组下添加 api
var HandlerRegisterFunc = func(root *echo.Group) {}

func Router(engine *echo.Echo) {

	// 入参校验
	engine.Validator = &iValidator{validator: validator.New()}

	// error 处理
	engine.HTTPErrorHandler = errorHandler()

	// cors
	engine.Use(cors())

	rootGroup := engine.Group("/iecho")

	// api
	HandlerRegisterFunc(rootGroup)
}
