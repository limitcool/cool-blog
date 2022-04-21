package model

import (
	"github.com/limitcool/blog/global"
	"log"
)

type Comment struct {
	*BaseModel
	UserID     uint   `json:"user_id"` // 发布评论的用户id
	Content    string `json:"content"` // 评论内容
	Root       int    `json:"root"`    // 根评论id,不为0即是回复评论
	Parent     uint   `json:"parent"`  // 父评论id 为0 是root评论
	ArticlesID uint   `json:"articles_id"`
}

func (c *Comment) Create() error {
	if err := global.DB.Create(&c).Error; err != nil {
		return err
	}
	return nil
}

// 查询文章下的所有评论
func (c *Comment) List() {
	var comments []Comment
	global.DB.Where("articles_id = ?", c.ArticlesID).Find(&comments)
	log.Fatal(comments)
}

// 只查询父评论
func (c *Comment) ListParent() {
	var comments []Comment
	global.DB.Where("articles_id = ? AND parent = ?", c.ArticlesID, 0).Find(&comments)
	log.Fatal(comments)
}

// 查询指定父评论下的子评论
func (c *Comment) ListSon() {
	var comments []Comment
	global.DB.Where("articles_id = ? AND parent = ?", c.ArticlesID, c.Parent).Find(&comments)
	//log.Fatal(comments)
}
