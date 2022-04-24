package cmd

import (
	"context"
	"fmt"
	_ "github.com/limitcool/blog/bootstrap"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

func Run() {

	router := router.NewRouter()
	//func() {
	//	for {
	//		fmt.Println(global.ServerSetting.HttpPort)
	//		time.Sleep(1 * time.Second)
	//	}
	//
	//}()
	s := &http.Server{
		Addr:           fmt.Sprint("0.0.0.0:", global.ServerSetting.HttpPort),
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		// 服务连接 监听
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen:%s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("开始关闭服务...")
	//(设置5秒超时时间)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("服务器关闭失败:", err)
	}
	log.Println("服务已关闭")
}
