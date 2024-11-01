package utils

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

func GenSnowFlakeId() snowflake.ID {
	return node.Generate()
}
