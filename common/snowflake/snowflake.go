package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/limitcool/blog/global"
	"log"
)

// 生成雪花id
func GenerateSnowFlakeId() int64 {
	var err error
	// Create a new Node with a Node number of 1
	global.Node, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("生成雪花ID失败:", err)
	}

	// Generate a snowflake ID.
	id := global.Node.Generate()

	return id.Int64()
}
