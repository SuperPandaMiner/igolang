package config

import (
	"iconfig"
	"iconfig/iviper"
)

func Init(file string) {
	iviper.Register(file)
	iconfig.Init()
}
