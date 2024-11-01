package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// HandlerRegisterFunc 在根路由组下添加 api
var HandlerRegisterFunc = func(root *gin.RouterGroup) {}

func Router() *gin.Engine {
	engine := gin.Default()

	//参数校验， 响应结果
	engine.Use(responseHandler())
	// 404
	engine.NoRoute(notFound())

	// 先设置响应头 #responseHandler()
	engine.Use(func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
	})
	// cors
	engine.Use(cors.New(corsConfig()))
	// 允许 option 请求
	engine.OPTIONS("/*any")

	rootGroup := engine.Group("")

	// api
	HandlerRegisterFunc(rootGroup)

	return engine
}
