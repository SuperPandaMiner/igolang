package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"igozero/config"
	"igozero/engine"
	"igozero/logger"
	"igozero/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Init("./config.yml")

	logger.Init()

	routers.HandlerRegisterFunc = func(engine *rest.Server) {
		engine.AddRoute(routers.Route(http.MethodGet, "/ping", func(r *http.Request) (any, error) {
			return "hello world", nil
		}))
	}
	engine.Run()

	quit := make(chan os.Signal, 1)
	// kill -2，kill -15
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	logx.Info("received signal: %s\n", sig)

	engine.ShutDown()

	logx.Info("server shutdown success")
}
