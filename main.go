package main

import (
	"github.com/limitcool/blog/common/setting"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/model"
	route "github.com/limitcool/blog/route"
	"log"
	"time"
)

func init() {
	// 读取配置文件
	{
		err := setupSetting()
		if err != nil {
			log.Fatalf("init.setupSetting err: %v", err)
		}
	}
	// 连接数据库
	{
		var err error
		global.DB, err = model.NewDBEngine(global.DatabaseSetting)
		if err != nil {
			log.Println(err)
		}
	}

}

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

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Jwt", &global.JwtSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
