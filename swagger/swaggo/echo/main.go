package main

import (
	"echo/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os/exec"
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
func GetFunc(ctx echo.Context) error {
	return ctx.JSON(200, "GetFunc")
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
func PostFunc(ctx echo.Context) error {
	return ctx.JSON(200, "PostFunc")
}

// @securityDefinitions.apikey ACCESS_TOKEN
// @in header
// @name access_token
// @Security ACCESS_TOKEN
func main() {

	// 如果 swaggo 注释在其他包下，需要 import 一下

	initDocsInfo()
	engine := echo.New()

	engine.Use(cors())

	// 注册 swagger api
	engine.GET("/swagger/*any", echoSwagger.WrapHandler)
	engine.GET("/swaggo/get/:id", GetFunc)
	engine.POST("/swaggo/post", PostFunc)

	engine.Start("127.0.0.1:8080")

}

func generateSwaggerDocs(host string) {
	// 执行 swag init 命令
	if err := exec.Command("swag", "init", "-g", "main.go", "-o", "docs").Run(); err != nil {
		panic("Failed to generate Swagger docs")
	}
}

// 修改生成的 swagger 通用信息
func initDocsInfo() {
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "Swaggo API"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowHeaders: []string{"Authorization", "content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host"},
	})
}
