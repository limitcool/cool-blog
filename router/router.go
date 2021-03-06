package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/common/captcha"
	_ "github.com/limitcool/blog/docs"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/controller"
	"github.com/limitcool/blog/internal/middleware"
	"github.com/limitcool/blog/router/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	var (
		articles = controller.NewArticleController()
		user     = controller.NewUserController()
		profile  = controller.NewProfileController()
	)
	logfile, err := os.Create(global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt)
	log.Println("logfile path :", logfile.Name())
	if err != nil {
		fmt.Println("无法创建log文件:", err)
	}
	r := gin.New()
	gin.SetMode("debug")
	//gin.DefaultWriter = io.MultiWriter(logfile)
	r.Use(gin.Logger(), gin.Recovery(), middleware.AppInfo())
	// 使用prometheus中间件
	p := middleware.NewPrometheus(r)
	r.Use(p.Middleware())
	r.Use(middleware.MaxAllowed(200)) //限制每秒最多允许200个请求
	r.Use(middleware.Cors())
	r.Use(middleware.ContextTimeout(5 * time.Second))
	// prometheus监控 采集指标
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	apiV1 := r.Group("/api/v1/articles/")
	//apiV1.Use(middleware.JWT())
	//apiV1.Use(middleware.CheckCasbinAuth())

	r.POST("/auth", api.GetAuth)

	r.POST("/login", user.Login)
	r.POST("/register", user.Register)
	r.GET("/getCaptcha", captcha.GenerateCaptcha)
	r.GET("/verifyCaptcha", captcha.CaptchaVerify)
	r.POST("/upload/file", controller.NewUpload().UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.StaticFS("/qrcode", http.Dir(global.AppSetting.QrCodeSavePath))
	{
		apiV1.POST("/create", articles.Create)
		apiV1.POST("/new_create", articles.NewCreate)
		// 获取指定文章
		apiV1.GET("/:article_id", articles.Get)
		apiV1.GET("/author/:author_name", articles.Count)
		// 获取文章列表
		apiV1.GET("", articles.List)
		// 获取HTML
		apiV1.GET("/html/:article_id", articles.GetHTML)
		// 更新指定文章
		apiV1.PUT("/:article_id", articles.Update)
		// 通过article_id 删除指定文章
		apiV1.DELETE("/:article_id", articles.Delete)
	}
	// profile
	{
		r.POST("/profile/create", profile.Create)
	}
	// sleep
	r.GET("/sleep", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.JSON(200, "OK!!")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "Nginx:1.20.1:Page not found",
		})
	})
	return r
}
