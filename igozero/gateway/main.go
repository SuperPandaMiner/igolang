package main

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/gateway"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var c gateway.GatewayConf
	conf.MustLoad("config.yml", &c)

	engine := gateway.MustNewServer(c)

	go func() {
		engine.Start()
	}()

	quit := make(chan os.Signal, 1)
	// kill -2，kill -15
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	logx.Info("received signal: %s\n", sig)

	engine.Stop()

	logx.Info("server shutdown success")

}
