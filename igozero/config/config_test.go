package config

import (
	"fmt"
	"testing"
)

func init() {
	Init("../config.yml")
}

func Test(t *testing.T) {
	fmt.Println(Config)
}
