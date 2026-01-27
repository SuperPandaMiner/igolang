package main

import (
	"fmt"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
)

type TestConfiguration struct {
	Value string `json:"value"`
	Name  string `json:"name,optional"`
}

func main() {
	// 创建 etcd subscriber
	ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
		Hosts: []string{"10.8.4.117:2379"},
		Key:   "/config.yml",
	})

	// 创建 configurator
	cc := configurator.MustNewConfigCenter[TestConfiguration](configurator.Config{
		Type: "yaml", // 配置值类型：json,yaml,toml
	}, ss)

	// 获取配置
	v, err := cc.GetConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	// 添加 listener 监听配置变化
	cc.AddListener(func() {
		v, err := cc.GetConfig()
		if err != nil {
			panic(err)
		}
		fmt.Println(v)
	})

	select {}
}
