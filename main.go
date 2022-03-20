package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/limitcool/blog/common/setting"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/model"
	"github.com/limitcool/blog/route"
	"log"
	"time"
)

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

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
	//初始化casbin
	{
		var err error
		global.Enforcer, err = InitCasbin()
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
	err = setting.ReadSection("Pay", &global.PaySetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func InitCasbin() (*casbin.Enforcer, error) {
	//a, err := gormadapter.NewAdapterByDBWithCustomTable(global.DB, &CasbinRule{}, "casbin_rule")
	a, err := gormadapter.NewAdapterByDBUseTableName(global.DB, "", "casbin_rule")
	if err != nil {
		return nil, err
	}
	e, _ := casbin.NewEnforcer("configs/casbin.conf", a)
	e.LoadPolicy()
	// 传递用户组别,请求地址,请求方式,对应v0,v1,v2
	//ok, errs := e.Enforce("1", "/", "GET")
	//if errs != nil {
	//	fmt.Println("104:", errs)
	//	return nil, errs
	//}
	//if ok {
	//	fmt.Println("hello")
	//} else {
	//	fmt.Println("false!!!")
	//}

	return e, err
}
