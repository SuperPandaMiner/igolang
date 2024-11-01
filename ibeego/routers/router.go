package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"ibeego/controllers"
)

func Init() {

	// request body 可多次读取
	web.BConfig.CopyRequestBody = true

	// 关闭 controller error handler 默认的页面渲染，开启的话会在响应时自动寻找 tpl
	web.BConfig.WebConfig.AutoRender = false
	// 注册 controller 错误自定义处理
	web.ErrorController(&controllers.ErrorController{})

	// cors
	web.InsertFilter("/", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// api
	root := web.NewNamespace("/api")
	ControllerRegisterFunc(root)
	web.AddNamespace(root)
}

// ControllerRegisterFunc 在根路由下添加 api
var ControllerRegisterFunc func(root *web.Namespace) = func(root *web.Namespace) {}
