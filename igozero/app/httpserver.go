package main

import (
	"github.com/igolang/igozero/core/config"
	"github.com/igolang/igozero/core/httpengine"
	"github.com/igolang/igozero/core/logger"
	"github.com/igolang/igozero/core/routers"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.InitRestConfig("./config.yml")

	logger.Init()

	routers.HandlerRegisterFunc = func(engine *rest.Server) {
		engine.AddRoute(routers.Route(http.MethodGet, "/app/ping", func(r *http.Request) (any, error) {
			return "hi! i am app", nil
		}))
		engine.AddRoute(routers.Route(http.MethodGet, "/app/pingping", func(r *http.Request) (any, error) {
			return "hi! i am app", nil
		}))
	}
	httpengine.Run()

	quit := make(chan os.Signal, 1)
	// kill -2，kill -15
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	logx.Info("received signal: %s\n", sig)

	httpengine.ShutDown()

	logx.Info("server shutdown success")
}
