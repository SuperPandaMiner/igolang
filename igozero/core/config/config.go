package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

var RestConfig RestConfiguration
var RpcConfig RpcConfiguration

type RestConfiguration struct {
	rest.RestConf
}

type RpcConfiguration struct {
	zrpc.RpcServerConf
}

func InitRestConfig(file string) {
	conf.MustLoad(file, &RestConfig)
}

func InitRpcConfig(file string) {
	conf.MustLoad(file, &RpcConfig)
}
