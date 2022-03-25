package service

import (
	"github.com/limitcool/blog/common/markdown"
	"github.com/limitcool/blog/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (svc *Service) GetHtml(aid uint) (string, error) {
	a := model.NewArticles()
	a.ArticleId = aid
	err := a.Info()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("查询的记录不存在")
	}
	if a.Content == "" {
		return "", errors.New("查询的内容为空")
	}
	return markdown.Render(a.Content), nil
}
