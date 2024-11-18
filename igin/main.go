package main

import (
	"github.com/gin-gonic/gin"
	"igin/config"
	"igin/engine"
	"igin/logger"
	"igin/orm"
	"igin/routers"
	"ilogger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	config.Init("config.yml")

	logger.Init()

	orm.Init()

	routers.HandlerRegisterFunc = func(root *gin.RouterGroup) {
		root.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "hello world")
		})
	}
	engine.Run()

	quit := make(chan os.Signal, 1)
	// kill -2，kill -15
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	ilogger.Warn("received signal: %s", sig)

	engine.ShutDown()

	orm.Close()

	ilogger.Close()

	ilogger.Info("server shutdown success")
}
