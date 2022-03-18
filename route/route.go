package route

import (
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/common/captcha"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/controller"
	dao2 "github.com/limitcool/blog/internal/dao"
	"github.com/limitcool/blog/internal/middleware"
	"github.com/limitcool/blog/route/api"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
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
	return r
}
