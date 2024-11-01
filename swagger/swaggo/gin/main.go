package main

import (
	"gin/gin/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
func GetFunc(ctx *gin.Context) {
	ctx.JSON(200, "GetFunc")
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
func PostFunc(ctx *gin.Context) {
	ctx.JSON(200, "PostFunc")
}

// @securityDefinitions.apikey ACCESS_TOKEN
// @in header
// @name access_token
// @Security ACCESS_TOKEN
func main() {

	// 如果 swaggo 注释在其他包下，需要 import 一下

	initDocsInfo()
	engine := gin.Default()

	engine.Use(cors.New(corsConfig()))

	// 注册 swagger api
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/swaggo/get/:id", GetFunc)
	engine.POST("/swaggo/post", PostFunc)

	engine.Run("127.0.0.1:8080")

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

func corsConfig() cors.Config {
	corsConf := cors.Config{}
	corsConf.AllowAllOrigins = true
	return corsConf
}
