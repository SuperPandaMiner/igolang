package config

import (
	"fmt"
	"iconfig"
	"testing"
)

func init() {
	Init("../../config.yml")
}

func Test(t *testing.T) {
	fmt.Println(iconfig.Server)
	fmt.Println(iconfig.Database)
	fmt.Println(iconfig.Logger)
}
