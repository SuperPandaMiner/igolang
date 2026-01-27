package main

import (
	"github.com/igolang/igozero/app/apprpc"
	"github.com/igolang/igozero/core/config"
	"github.com/igolang/igozero/core/logger"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"os/signal"
	"syscall"

	"github.com/igolang/igozero/app/internal/server"
	"github.com/igolang/igozero/app/internal/svc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.InitRpcConfig("./config.yml")
	logger.Init()

	ctx := svc.NewServiceContext(config.RpcConfig)

	engine := zrpc.MustNewServer(config.RpcConfig.RpcServerConf, func(grpcServer *grpc.Server) {
		// 注册 grpc 服务
		apprpc.RegisterAppRpcServer(grpcServer, server.NewAppRpcServer(ctx))

		// 开启 api 反射
		if config.RpcConfig.RpcServerConf.Mode == service.DevMode || config.RpcConfig.RpcServerConf.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

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
