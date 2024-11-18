package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"igin/config"
	"igin/engine"
	"igin/routers"
	"testing"
)

func init() {
	config.Init("../../config.yml")

	routers.HandlerRegisterFunc = func(root *gin.RouterGroup) {
		root.Use(tokenFilter)
		root.GET("/success", success)
		root.GET("/400", badRequest)
		root.GET("/401", authentication)
		root.GET("/403", forbidden)
		root.GET("/500", internalServerError)
		root.GET("/panic", _panic)
		root.GET("/query", query)
		root.GET("/path/:id", path)
		root.POST("/post", post)
	}

}

func success(ctx *gin.Context) {
	writeSuccess(ctx, "success")
}

func badRequest(ctx *gin.Context) {
	ctx.AbortWithError(400, errors.New("400"))
}

func authentication(ctx *gin.Context) {
	ctx.AbortWithError(401, errors.New("401"))
}

func forbidden(ctx *gin.Context) {
	ctx.AbortWithError(403, errors.New("403"))
}

func internalServerError(ctx *gin.Context) {
	abort(ctx, errors.New("internal server error"))
}

func _panic(ctx *gin.Context) {
	panic("panic")
}

func query(ctx *gin.Context) {
	params := make(map[string]interface{})
	params["string"] = ctx.Query("string")
	i, _ := getInt64(ctx, "int")
	params["int"] = i
	b, _ := getBool(ctx, "bool")
	params["bool"] = b
	writeSuccess(ctx, params)
}

func path(ctx *gin.Context) {
	id := ctx.Param("id")
	writeSuccess(ctx, id)
}

func post(ctx *gin.Context) {
	type param struct {
		String string `json:"string" binding:"required" validate:"不能为空"`
		Int    int    `json:"int" binding:"min=1,max=10"`
		Bool   bool   `json:"bool"`
		Name   string `json:"name" binding:"max=5"`
	}
	body := &param{}
	if err := ctx.Bind(&body); err == nil {
		writeSuccess(ctx, body)
	}
}

func tokenFilter(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		abortWithCode(ctx, 401, errors.New("authorization required"))
	}
}

func Test(t *testing.T) {
	engine.Run()

	select {}
}
