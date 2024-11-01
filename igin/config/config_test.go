package config

import (
	"fmt"
	"igin/config/iconfig"
	"igin/config/jinzhu"
	"testing"
)

func init() {
	jinzhu.Register("../config.yml")
	iconfig.Init()
}

func Test(t *testing.T) {
	fmt.Println(iconfig.Server)
	fmt.Println(iconfig.Database)
	fmt.Println(iconfig.Logger)
}
