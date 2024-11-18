package main

import (
	"github.com/labstack/echo/v4"
	"iecho/config"
	"iecho/engine"
	"iecho/logger"
	"iecho/orm"
	"iecho/routers"
	"ilogger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Init("../config.yml")

	logger.Init()

	orm.Init()

	routers.HandlerRegisterFunc = func(root *echo.Group) {
		root.GET("/ping", func(ctx echo.Context) error {
			return ctx.JSON(http.StatusOK, "hello world")
		})
	}
	engine.Run()

	quit := make(chan os.Signal, 1)
	// kill -2，kill -15
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	ilogger.Warn("received signal: %s\n", sig)

	engine.ShutDown()

	orm.Close()

	ilogger.Close()

	ilogger.Info("server shutdown success")
}
