package utils

import (
	"github.com/bwmarrin/snowflake"
	"log"
	"strconv"
)

// GetID using snowflake algorithm to generate ID
func GetID() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Panic(err)
	}
	return node.Generate().Int64()
}

// Stoi transfer string to int
func Stoi(s string, n int) interface{} {
	t, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	switch n {
	case 8:
		return int8(t)
	case 16:
		return int16(t)
	case 32:
		return int32(t)
	case 64:
		return int64(t)
	default:
		return t
	}
}
