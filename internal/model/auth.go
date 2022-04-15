package model

import (
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/util"
)

type Auth struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
}

func (receiver Auth) TableName() string {
	return "blog_auth"
}

// 对数据库进行查询,返回结构体和查询情况
func (a Auth) Get() (Auth, error) {
	var auth Auth
	// 对传入的密码进行MD5加密
	md5password := util.Md5(a.Password)
	db := global.DB.Table("users").Where("username = ? AND password = ?", a.Username, md5password)
	err := db.First(&auth).Error
	if err != nil {
		return auth, err
	}
	return auth, nil
}

// 对数据库进行写入
func (a Auth) Create(u User) {
	//auth := Auth{ID: u.ID, Username: u.Username, Password: u.Password}
	//global.DB.Select("id", "app_key", "app_secret").Create(&auth)
	//global.DB.Save(&auth)
}
