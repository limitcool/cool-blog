package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/model"
	"net/http"
)

// 创建结构体类型 Article
type Article struct{}

// 新建Article 实例
func NewArticle() Article {
	return Article{}
}

// 实现Get方法
func (a Article) Get(c *gin.Context) {
	c.String(http.StatusOK, "调用Get成功")
}

// 实现List方法
func (a Article) List(c *gin.Context) {
	//query := model.CreateArticleQuery{}
	//err := c.BindJSON(&query)
	//if err != nil {
	//	c.String(http.StatusBadRequest, err.Error())
	//} else {
	//	c.JSON(http.StatusOK, query)
	//}
	ArticleId := c.Param("article_id")
	article := model.Articles{}
	global.DB.Find(&article, ArticleId)
	c.JSON(http.StatusOK, article)
}

// 实现Create方法 新建文章
//func (a Article) Create(c *gin.Context) {
//	//json := model.CreateArticleQuery{}
//	//err := c.BindJSON(&json)
//	if err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//	} else {
//		c.JSON(http.StatusOK, json)
//	}
//}

// 实现Creates方法 批量新建文章
func Creates(c *gin.Context) {
	json := model.Articles{}
	err := c.BindJSON(&json)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, json)
	}
}

// 实现Update方法
func (a Article) Update(c *gin.Context) {

}

// 实现Delete方法
func (a Article) Delete(c *gin.Context) {

}
