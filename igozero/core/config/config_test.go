package config

import (
	"fmt"
	"testing"
)

func init() {
	InitRestConfig("../config.yml")
}

func Test(t *testing.T) {
	fmt.Println(RestConfig)
}
