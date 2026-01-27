package svc

import "github.com/igolang/igozero/core/config"

type ServiceContext struct {
	Config config.RpcConfiguration
}

func NewServiceContext(c config.RpcConfiguration) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
