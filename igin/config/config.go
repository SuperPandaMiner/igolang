package config

import (
	"github.com/igolang/iconfig"
	"github.com/igolang/iconfig/iviper"
)

func Init(file string) {
	iviper.Register(file)
	iconfig.Init()
}
