package model

import (
	"fmt"
	"github.com/limitcool/blog/common"
	"github.com/limitcool/blog/global"
	"log"
	"time"
)

//type Article struct {
//	Id     int    `json:"id"`
//	Title  string `json:"title"`
//	Author string `json:"author"`
//}

type Articles struct {
	BaseModel
	Title   string `json:"title"`   // 文章标题
	Author  string `json:"author"`  // 作者
	Content string `json:"content"` // 文章内容
	//MarkdownUrl string `json:"markdown_url"`                                // markdown上传后得到的url
	Tags     []Tag     `gorm:"many2many:articles_tags" json:"tags"` // 标签
	Category Category  `json:"category"`                            // 分类
	UserID   uint      `json:"user_id"`
	Comments []Comment `json:"comments"` // 评论
}

//func NewArticle(id int, title, author string) Article {
//	return Article{Id: id, Title: title, Author: author}
//}

func NewArticles() Articles {
	return Articles{}
}

func (a Articles) TableName() string {
	return "articles"
}

// 新建文章 必须传递指针参数才能改变原值
func (a *Articles) Create() bool {
	var tmp Articles
	sql := "INSERT  INTO articles(title,author,content,created_at,updated_at) SELECT ?,?,?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM articles WHERE  title=?)"
	result := global.DB.Model(&tmp).Exec(sql, a.Title, a.Author, a.Content, a.CreatedAt, a.UpdatedAt, a.Title)
	// 判断是否插入数据成功
	if result.RowsAffected > 0 {
		result = global.DB.Where("title = ?", a.Title).First(&a)
		return true
	} else {
		result = global.DB.Where("title = ?", a.Title).First(&a)
		fmt.Println(a.ID)
		return false
	}
}

// new 新建文章 必须传递指针才能改变原值
func (a *Articles) NewCreate(token string) error {
	//err := global.DB.Find(&Category{}, "category_name", &a.Category.CategoryName).Error
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return global.DB.Create(&a).Error
	//	}
	//	return err
	//}
	//
	claim, err := common.ParseToken(token)
	if err != nil {
		return err
	}
	a.UserID = GetIdByUsername(claim.Username)
	//a.UserID =
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

// 批量文章查询
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
	global.DB.Where("article_id = ?", a.ID).Save(&a)
}

// 删除文章
func (a *Articles) Delete() {
	global.DB.Where("article_id = ?", a.ID).Delete(&a)
}

// 获取文章信息
func (a *Articles) Info() (err error) {
	if a.ID < 0 {
		log.Println("输入的ID不正确")
		return err
	}
	// Preload里面要填结构体名称
	err = global.DB.Preload("Tags").Preload("Category").Preload("Comments").First(&a, a.ID).Error
	if err != nil {
		log.Println(err)
		log.Println("查询失败")
		return err
	}
	return nil
}

// Markdown渲染为Html
func (a *Articles) MarkdownToHtml() {

}

// GetTag 获取文章拥有的标签列表
func (a *Articles) GetTag() []Tag {
	//err := global.DB.Preload("Tags").Preload(clause.Associations).Where("article_id = ?", a.ArticleId).First(&a).Error
	err := global.DB.Preload("Tags").Where("article_id = ?", a.ID).First(&a).Error
	if err != nil {
		fmt.Println(err)
	}
	return a.Tags
}

// GetIdByUsername  通过username获取userid
func GetIdByUsername(username string) uint {
	var user User
	global.DB.Where("username = ?", username).Find(&user)
	return user.ID
}
