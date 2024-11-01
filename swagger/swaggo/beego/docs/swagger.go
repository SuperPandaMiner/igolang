package docs

import (
	_ "embed"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/swaggo/swag"
	"os/exec"
)

// Init 修改生成的 swagger 通用信息
func Init(version, host, basePath, title string) {
	SwaggerInfo.Version = version
	SwaggerInfo.Host = host
	SwaggerInfo.BasePath = basePath
	SwaggerInfo.Title = title
	SwaggerInfo.Schemes = []string{"http", "https"}
}

func generateSwaggerDocs(host string) {
	// 执行 swag init 命令
	if err := exec.Command("swag", "init", "-g", "main.go", "-o", "docs").Run(); err != nil {
		panic("Failed to generate Swagger docs")
	}
}

func SwaggerJson(ctx *context.Context) {
	json, _ := swag.ReadDoc(SwaggerInfo.InfoInstanceName)
	ctx.WriteString(json)
}
