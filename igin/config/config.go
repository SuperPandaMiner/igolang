package config

import (
	"igin/config/iconfig"
	"igin/config/jinzhu"
)

func Init(file string) {
	jinzhu.Register(file)
	iconfig.Init()
}
