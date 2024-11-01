package main

import (
	"beego/docs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

type ViewVO struct {
	// 字符串
	String string `json:"string"`
	// 布尔
	Bool bool `json:"bool"`
	// 数字
	Number int `json:"number"`
}

// @Tags api
// @Router /swaggo/get/{id} [get]
// @Summary get
// @Description get
// @Security ACCESS_TOKEN
// @Accept  json
// @Produce  json
// @Param id path string false "id"
// @Param string query string false "字符串"
// @Param bool query bool false "布尔"
// @Param number query int false "数字"
// @success 200 {object} ViewVO
func GetFunc(ctx *context.Context) {
	ctx.WriteString("GetFunc")
}

// @Tags api
// @Router /swaggo/post [post]
// @Summary post
// @Description post
// @Security ACCESS_TOKEN
// @Accept  json
// @Produce  json
// @Param params body ViewVO true "params"
// @success 200 {array} ViewVO
func PostFunc(ctx *context.Context) {
	ctx.WriteString("PostFunc")
}

// @securityDefinitions.apikey ACCESS_TOKEN
// @in header
// @name access_token
// @Security ACCESS_TOKEN
func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "docs"

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
	}))

	// 获取 swagger.json 的接口
	beego.Get("/swagger/json", docs.SwaggerJson)

	beego.Get("/swaggo/get/:id", GetFunc)
	beego.Post("/swaggo/post/", PostFunc)

	docs.Init("1.0.0", "127.0.0.1:8080", "/", "Swaggo API")

	beego.Run("127.0.0.1:8080")
}
