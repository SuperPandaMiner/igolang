package main

import (
	_ "github.com/beego/beego/v2/core/config/yaml"
	"github.com/beego/beego/v2/server/web"
	"ibeego/conf"
	"ibeego/controllers"
	"ibeego/iorm"
	"ibeego/logger"
	"ibeego/routers"
)

func main() {
	conf.Init()
	logger.Init()
	iorm.Init()
	routerExample()
	routers.Init()
	web.Run()
}

func routerExample() {
	routers.ControllerRegisterFunc = func(root *web.Namespace) {
		root.CtrlGet("success", controllers.ExampleController.Success)
		root.CtrlGet("400", controllers.ExampleController.BadRequest)
		root.CtrlGet("401", controllers.ExampleController.Authentication)
		root.CtrlGet("403", controllers.ExampleController.Forbidden)
		root.CtrlGet("500", controllers.ExampleController.InternalServerError)
		root.CtrlGet("abort", controllers.ExampleController.AbortWithCode)
		root.CtrlGet("panic", controllers.ExampleController.Panic)
		root.CtrlGet("query", controllers.ExampleController.Query)
		root.CtrlGet("path/:id", controllers.ExampleController.Path)
		root.CtrlPost("body", controllers.ExampleController.Body)
		root.CtrlPost("upload", controllers.ExampleController.Upload)
		root.CtrlGet("download", controllers.ExampleController.Download)
	}
	web.InsertFilter("/*", web.BeforeRouter, controllers.TokenFilter)
}
