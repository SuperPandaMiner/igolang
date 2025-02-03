package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var Config Configuration

type Configuration struct {
	rest.RestConf
}

func Init(file string) {
	conf.MustLoad(file, &Config)
}
