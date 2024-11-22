package test

import (
	"fmt"
	"iconfig"
	"iconfig/iviper"
	"testing"
)

func init() {
	// 使用 jinzhu config
	//jinzhu.Register("../config.yml")
	// 使用 viper config
	iviper.Register("../config.yml")
	iconfig.Init()
}

func Test(t *testing.T) {
	fmt.Println(iconfig.Server)
	fmt.Println(iconfig.Database)
	fmt.Println(iconfig.Logger)
	fmt.Println(iconfig.Zap)
}
