package cache_service

import (
	"strconv"
	"strings"
)

const (
	cacheArticle = "article"
	cacheTag     = "tag"
)

type Article struct {
	ID       uint
	Title    string
	Author   string
	Tags     []string
	Category string
}

func (a *Article) GetArticleKey() string {
	return cacheArticle + "_" + strconv.Itoa(int(a.ID))
}

// 获取文章缓存key的方法
func (a *Article) GetArticlesKey() string {
	keys := []string{
		cacheArticle,
		"List",
	}
	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(int(a.ID)))
	}
	if len(a.Tags) > 0 {
		// 遍历a.tags解构,并将数据存入keys
		for _, v := range a.Tags {
			keys = append(keys, v)
		}
	}
	if a.Title != "" {
		keys = append(keys, a.Title)
	}
	if a.Author != "" {
		keys = append(keys, a.Author)
	}
	if a.Category != "" {
		keys = append(keys, a.Category)
	}
	return strings.Join(keys, "_")
}
