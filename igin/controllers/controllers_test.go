package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"igin/config"
	"igin/engine"
	"igin/routers"
	"strconv"
	"testing"
	"time"
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
		root.POST("/upload", upload)
		root.GET("/download", download)
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

var resourceDir = "../static"
var resourceBuffer = make(map[string]string)

func upload(ctx *gin.Context) {
	uploadErr := errors.New("upload failed")

	file, err := ctx.FormFile("file")
	if err != nil {
		abort(ctx, uploadErr)
		return
	}

	id := strconv.FormatInt(time.Now().UnixMilli(), 10)
	err = ctx.SaveUploadedFile(file, resourceDir+"/"+id)
	if err != nil {
		abort(ctx, uploadErr)
		return
	}
	resourceBuffer[id] = resourceDir + "/" + id
	writeSuccess(ctx, id)
	return
}

func download(ctx *gin.Context) {
	id := ctx.Query("id")
	buffer := resourceBuffer[id]
	if buffer == "" {
		abortWithCode(ctx, 404, errors.New("file not found"))
		return
	}
	ctx.Header("Content-Type", "")
	//ctx.File(buffer)
	ctx.FileAttachment(buffer, id)
	return
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
