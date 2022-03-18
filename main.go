package main

import (
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/model"
	"github.com/limitcool/blog/internal/pkg/setting"
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
