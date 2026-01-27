package config

import (
	"fmt"
	"github.com/igolang/iconfig"
	"testing"
)

func init() {
	Init("../../iconfig/config.yml")
}

func Test(t *testing.T) {
	fmt.Println(iconfig.Server)
	fmt.Println(iconfig.Database)
	fmt.Println(iconfig.Logger)
}
