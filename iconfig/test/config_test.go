package test

import (
	"fmt"
	"iconfig"
	"iconfig/jinzhu"
	"testing"
)

func init() {
	// 使用 jinzhu config
	jinzhu.Register("config.yml")
	iconfig.Init()
}

func Test(t *testing.T) {
	fmt.Println(iconfig.Server)
	fmt.Println(iconfig.Database)
	fmt.Println(iconfig.Logger)
	fmt.Println(iconfig.Zap)
}
