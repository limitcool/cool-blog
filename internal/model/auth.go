package model

import (
	"github.com/limitcool/blog/global"
)

type Auth struct {
	UserId   uint   `gorm:"primarykey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (receiver Auth) TableName() string {
	return "blog_auth"
}

// 对数据库进行查询,返回结构体和查询情况
func (a Auth) Get() (Auth, error) {
	var auth Auth
	db := global.DB.Table("users").Where("username = ? AND password = ?", a.Username, a.Password)
	err := db.First(&auth).Error
	if err != nil {
		return auth, err
	}
	return auth, nil
}

// 对数据库进行写入
func (a Auth) Create(u User) {
	auth := Auth{UserId: u.UserId, Username: u.Username, Password: u.Password}
	//global.DB.Select("id", "app_key", "app_secret").Create(&auth)
	global.DB.Save(&auth)
}
