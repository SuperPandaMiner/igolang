package main

import (
	"context"
	"github.com/igolang/igozero/app/apprpc"
	"github.com/igolang/igozero/app/apprpcclient"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	clientConf := zrpc.RpcClientConf{}
	// 填充默认值，比如 trace 透传等
	err := conf.FillDefault(&clientConf)
	if err != nil {
		return
	}

	// 单节点
	clientConf.Target = "dns:///127.0.0.1:8080"
	// 多节点
	clientConf.Endpoints = []string{"127.0.0.1:8080"}
	// 注册中心
	clientConf.Etcd = discov.EtcdConf{
		Hosts:              []string{"127.0.0.1:2379"},
		Key:                "app.rpc",
		User:               "",
		Pass:               "",
		CertFile:           "",
		CertKeyFile:        "",
		CACertFile:         "",
		InsecureSkipVerify: false,
	}

	conn := zrpc.MustNewClient(clientConf)
	client := apprpcclient.NewAppRpc(conn)
	resp, err := client.Ping(context.Background(), &apprpc.PingRequest{Ping: 1})
	if err != nil {
		logx.Error(err)
	} else {
		logx.Info(resp)
	}
}
