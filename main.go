package main

import (
	_ "github.com/limitcool/blog/bootstrap"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/route"
)

// @title           Blog
// @version         1.0
// @description     个人自用的Go博客系统的后端服务,采用gin框架+mysql数据库构建,目前正在实现中。
// @termsOfService  https://github.com/limitcool/cool-blog

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

func main() {

	route := route.NewRouter()
	route.Run("0.0.0.0:" + global.ServerSetting.HttpPort)

}
