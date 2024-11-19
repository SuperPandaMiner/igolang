package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"iecho/config"
	"iecho/engine"
	"iecho/logger"
	"iecho/routers"
	"strconv"
	"testing"
	"time"
	"utils"
)

type customType string

func (t *customType) UnmarshalParam(src string) error {
	*t = customType("custom" + src)
	return nil
}

func init() {
	config.Init("../../config.yml")

	logger.Init()

	routers.HandlerRegisterFunc = func(root *echo.Group) {
		root.Use(tokenFilter)
		root.GET("/success", success)
		root.GET("/400", badRequest)
		root.GET("/401", authentication)
		root.GET("/403", forbidden)
		root.GET("/500", internalServerError)
		root.GET("/log", EchoLog)
		root.GET("/panic", _panic)
		root.GET("/query", query)
		root.GET("/path/:id", path)
		root.POST("/post", post)
		root.POST("/upload", upload)
		root.GET("/download", download)
	}
}

func success(ctx echo.Context) error {
	return writeSuccess(ctx, "success")
}

func badRequest(ctx echo.Context) error {
	return abortWithCode(400, errors.New("400"))
}

func authentication(ctx echo.Context) error {
	return abortWithCode(401, errors.New("401"))
}

func forbidden(ctx echo.Context) error {
	return abortWithCode(403, errors.New("403"))
}

func internalServerError(ctx echo.Context) error {
	return abort(errors.New("internal server error"))
}

func EchoLog(ctx echo.Context) error {
	ctx.Echo().Logger.Debug("debug log")
	ctx.Echo().Logger.Info("info log")
	ctx.Echo().Logger.Warn("warn log")
	ctx.Echo().Logger.Error("error log")
	return nil
}

func _panic(ctx echo.Context) error {
	panic("panic")
}

// http://127.0.0.1:8080/iecho/query?p=p&string=hello&int=123&bool=true&name=zhangsan&ints=1&ints=2&delimiterString=1,2,3&customType=type&token=token
func query(ctx echo.Context) error {
	type params struct {
		P               string     `json:"p"`
		String          string     `json:"string"`
		Int             int        `json:"int"`
		Bool            bool       `json:"bool"`
		Name            string     `json:"name"`
		Ints            []int      `json:"ints"`
		DelimiterString []string   `json:"delimiterString"`
		CustomType      customType `json:"customType"`
	}
	md := params{}
	md.P = ctx.QueryParam("p")

	err := echo.QueryParamsBinder(ctx).
		String("string", &md.String).Int("int", &md.Int).Bool("bool", &md.Bool).
		MustString("name", &md.Name).
		Ints("ints", &md.Ints).
		BindWithDelimiter("delimiterString", &md.DelimiterString, ",").
		BindUnmarshaler("customType", &md.CustomType).
		BindError()
	if err != nil {
		return err
	}
	return writeSuccess(ctx, &md)
}

func path(ctx echo.Context) error {
	id := ctx.Param("id")
	return writeSuccess(ctx, id)
}

func post(ctx echo.Context) error {
	type param struct {
		String string `json:"string" validate:"required"`
		Int    int    `json:"int" validate:"min=1,max=10"`
		Bool   bool   `json:"bool"`
		Name   string `json:"name" validate:"max=5"`
	}

	md := &param{}
	err := ctx.Bind(md) // (&echo.DefaultBinder{}).BindBody(ctx, &md)
	if err != nil {
		return err
	}

	// valid
	err = ctx.Validate(md)
	if err != nil {
		return err
	}

	return writeSuccess(ctx, md)
}

var resourceDir = "../static"
var resourceBuffer = make(map[string]string)

func upload(ctx echo.Context) error {
	uploadErr := errors.New("upload failed")

	file, err := ctx.FormFile("file")
	if err != nil {
		return abort(uploadErr)
	}
	srcFile, err := file.Open()
	if err != nil {
		return abort(uploadErr)
	}
	defer srcFile.Close()

	id := strconv.FormatInt(time.Now().UnixMilli(), 10)
	err = utils.SaveFile(srcFile, resourceDir, id)
	if err != nil {
		return abort(uploadErr)
	}
	resourceBuffer[id] = resourceDir + "/" + id
	return writeSuccess(ctx, id)
}

func download(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	buffer := resourceBuffer[id]
	if buffer == "" {
		return abortWithCode(404, errors.New("file not found"))
	}
	//err := ctx.File(buffer)
	//err := ctx.Inline(buffer, id)
	err := ctx.Attachment(buffer, id)
	return err
}

// Header Token
func tokenFilter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var headers map[string]string
		err := (&echo.DefaultBinder{}).BindHeaders(ctx, &headers)
		if err != nil {
			return err
		}
		if headers["Token"] == "" {
			return abortWithCode(401, errors.New("authorization required"))
		}
		return next(ctx)
	}
}

func Test(t *testing.T) {

	engine.Run()

	select {}
}
