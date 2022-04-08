package global

import (
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
)

var (
	Node *snowflake.Node
	Vp   *viper.Viper
)
