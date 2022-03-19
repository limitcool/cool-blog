package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/common/captcha"
	_ "github.com/limitcool/blog/docs"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/controller"
	dao2 "github.com/limitcool/blog/internal/dao"
	"github.com/limitcool/blog/internal/middleware"
	"github.com/limitcool/blog/route/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
)

func NewRouter() *gin.Engine {

	logfile, err := os.Create(global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt)
	log.Println(logfile)
	if err != nil {
		fmt.Println("无法创建log文件:", err)
	}
	r := gin.New()
	gin.SetMode("debug")
	//gin.DefaultWriter = io.MultiWriter(logfile)
	r.Use(gin.Logger(), gin.Recovery(), middleware.AppInfo())
	r.Use(middleware.Cors())
	article := dao2.NewArticle()
	apiV1 := r.Group("/api/v1/articles/")
	apiV1.Use(middleware.JWT())
	articles := controller.NewArticleController()
	r.POST("/auth", api.GetAuth)
	user := controller.NewUserController()
	r.POST("/login", user.Login)
	r.POST("/register", user.Register)
	r.GET("/getCaptcha", captcha.GenerateCaptcha)
	r.GET("/verifyCaptcha", captcha.CaptchaVerify)
	r.POST("/upload/file", controller.NewUpload().UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	{
		apiV1.GET("/articles/cookie", dao2.Cookie)
		apiV1.POST("/create", articles.Create)
		apiV1.POST("/new_create", articles.NewCreate)
		// 获取指定文章
		apiV1.GET("/:article_id", article.List)
		apiV1.GET("/author/:author_name", articles.Count)
		// 获取文章列表
		apiV1.GET("", articles.List)
		// 更新指定文章
		apiV1.PUT("/:article_id", articles.Update)
		// 通过article_id 删除指定文章
		apiV1.DELETE("/:article_id", articles.Delete)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
