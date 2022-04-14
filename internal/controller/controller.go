package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/limitcool/blog/common/errcode"
	response2 "github.com/limitcool/blog/common/response"
	"github.com/limitcool/blog/internal/model"
	"github.com/limitcool/blog/internal/service"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// 定义一个控制器结构体
type ArticleController struct {
	engie *gorm.DB
}

type Articles struct {
	PageOffset int `json:"page_offset"`
	PageSize   int `json:"page_size"`
}

// 新建文章 控制器
func (a ArticleController) Create(c *gin.Context) {
	userIp := c.ClientIP()
	query := model.Articles{}
	err := c.BindJSON(&query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "error!")
	} else {
		fmt.Println(userIp)
		result := query.Create()
		fmt.Println(result)
		if result == false {
			c.String(http.StatusOK, "标题已存在！")
		} else {
			c.JSON(http.StatusOK, query)
		}

	}

}

// newCreate

func (a ArticleController) NewCreate(c *gin.Context) {
	query := &model.Articles{}
	err := c.ShouldBindJSON(&query)
	if err != nil {
		fmt.Println(err)
	} else {
		err = query.NewCreate()
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println(query.ArticleId)
		c.JSON(http.StatusOK, query)
	}
}
func (a ArticleController) Count(c *gin.Context) {
	authorName := c.Param("author_name")
	fmt.Println(authorName)
	articles := model.Articles{Author: authorName}
	c.JSON(http.StatusOK, articles.Count())
}

// 查询方法
func (a ArticleController) List(c *gin.Context) {
	var query Articles
	query.PageOffset, _ = strconv.Atoi(c.Query("page_offset"))
	query.PageSize, _ = strconv.Atoi(c.Query("page_size"))

	newArticles := model.NewArticles()
	articles := newArticles.List(query.PageOffset, query.PageSize)
	if articles != nil {
		c.JSON(http.StatusOK, articles)
	} else {
		c.JSON(http.StatusBadRequest, "输入错误,请检查参数!")
	}

}

// 更新方法

func (a ArticleController) Update(c *gin.Context) {
	newArticles := model.NewArticles()
	c.BindJSON(&newArticles)
	newArticles.Update()
	c.JSON(http.StatusOK, newArticles)
}

// 删除方法

func (a ArticleController) Delete(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("article_id"))
	if articleId == 0 {
		c.JSON(http.StatusBadRequest, "error!")
		return
	}
	newarticle := &model.Articles{ArticleId: uint(articleId)}
	newarticle.Delete()
	c.JSON(http.StatusOK, newarticle)
}
func NewArticleController() ArticleController {
	return ArticleController{}
}

func (a ArticleController) GetHTML(c *gin.Context) {
	response := response2.NewResponse(c)
	iArticleId, _ := strconv.Atoi(c.Param("article_id"))

	articleId := uint(iArticleId)
	svc := service.New(c)
	content, err := svc.GetHtml(articleId)
	if err != nil {
		response.ToResponse(err.Error())
		return
	}
	if content == "" {
		response.ToErrorResponse(errcode.NotFound)
	} else {
		c.String(http.StatusOK, content)
	}

}
