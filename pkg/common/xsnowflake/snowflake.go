package xsnowflake

import (
	"github.com/bwmarrin/snowflake"
	"lark/pkg/utils"
	"strconv"
)

var Snowflake *snowflakeNode

type snowflakeNode struct {
	node *snowflake.Node
}

func init() {
	var (
		node *snowflake.Node
		err  error
	)
	// Create a new Node with a Node number of 1
	node, err = snowflake.NewNode(1)
	if err != nil {
		return
	}
	Snowflake = &snowflakeNode{node}
}

// Generate a snowflake ID.
func NewSnowflakeID() int64 {
	return Snowflake.node.Generate().Int64()
}

func DefaultLarkId() string {
	return utils.SixteenMD5(strconv.FormatInt(NewSnowflakeID(), 10))
}
