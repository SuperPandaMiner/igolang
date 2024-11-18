package config

import (
	"iconfig"
	"iconfig/jinzhu"
)

func Init(file string) {
	jinzhu.Register(file)
	iconfig.Init()
}
