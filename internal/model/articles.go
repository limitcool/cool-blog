package model

import (
	"fmt"
	"github.com/limitcool/blog/global"
	"log"
	"time"
)

type Article struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Articles struct {
	BaseModel
	ArticleId  uint   `gorm:"primaryKey;AUTO_INCREMENT" json:"article_id"` // 文章id
	Title      string `json:"title"`                                       // 文章标题
	ArticleTag string `json:"article_tag"`                                 // 文章标签
	Author     string `json:"author"`                                      // 作者
	Content    string `json:"content"`                                     // 文章内容
}

func NewArticle(id int, title, author string) Article {
	return Article{Id: id, Title: title, Author: author}
}

func NewArticles() Articles {
	return Articles{}
}

func (a Articles) TableName() string {
	return "articles"
}

// 新建文章 必须传递指针参数才能改变原值
func (a *Articles) Create() bool {
	var tmp Articles
	sql := "INSERT  INTO articles(title,article_tag,author,content,created_at,updated_at) SELECT ?,?,?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM articles WHERE  title=?)"
	result := global.DB.Model(&tmp).Exec(sql, a.Title, a.ArticleTag, a.Author, a.Content, a.CreatedAt, a.UpdatedAt, a.Title)
	// 判断是否插入数据成功
	if result.RowsAffected > 0 {
		result = global.DB.Where("title = ?", a.Title).First(&a)
		return true
	} else {
		result = global.DB.Where("title = ?", a.Title).First(&a)
		fmt.Println(a.ArticleId)
		return false
	}
}

// new 新建文章 必须传递指针才能改变原值
func (a *Articles) NewCreate() error {
	return global.DB.Create(&a).Error
}

// 文章数量计数
func (a *Articles) Count() int64 {
	var count int64
	if a.Author != "" {
		// 通过作者查询文章数量
		global.DB.Model(&a).Where("author =?", a.Author).Count(&count)
	} else {
		log.Println("查询出现错误")
		return -1
	}
	return count
}

// 文章查询
func (a *Articles) List(pageOffset, pageSize int) []*Articles {
	var articles []*Articles
	if pageOffset >= 0 && pageSize > 0 {
		global.DB.Limit(pageSize).Offset(pageOffset).Find(&articles)
	} else {
		log.Println("输入的值不符合规范")
		return articles
	}
	if a.Title != "" {
		global.DB.Where("title = ?", a.Title)
	}

	return articles
}

// 更新文章
func (a *Articles) Update() {
	var t time.Time
	fmt.Println(t)
	if a.CreatedAt == t {
		return
	}
	global.DB.Where("article_id = ?", a.ArticleId).Save(&a)
}

func (a *Articles) Delete() {
	global.DB.Where("article_id = ?", a.ArticleId).Delete(&a)
}
