package dao

import (
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/model"
)

func (d Dao) ProfileCreate(desc, img string) (uint, error) {
	profile := model.Profile{Desc: desc, Img: img}
	err := global.DB.Create(&profile).Error
	if err != nil {
		return 0, err
	}
	return profile.ID, err
}
