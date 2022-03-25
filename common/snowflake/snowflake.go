package util

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

// 传入起始时间,机械ID
func SnowFlake(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}
func ()  {
	
}