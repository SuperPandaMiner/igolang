package routers

import (
	"github.com/gin-gonic/gin"
)

// HandlerRegisterFunc 在根路由组下添加 api
var HandlerRegisterFunc = func(root *gin.RouterGroup) {}

func Router() *gin.Engine {
	engine := gin.Default()

	// error 处理
	engine.Use(errorHandler())
	// 404
	engine.NoRoute(notFoundHandler())

	// 先设置响应头 #errorHandler()
	engine.Use(func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
	})
	// cors
	engine.Use(corsHandler())
	// 允许 option 请求
	engine.OPTIONS("/*any")

	rootGroup := engine.Group("/igin")

	// api
	HandlerRegisterFunc(rootGroup)

	return engine
}
