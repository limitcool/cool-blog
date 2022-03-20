package model

import (
	"github.com/limitcool/blog/global"
	"log"
)

type User struct {
	UserId    uint   `gorm:"primary_key" json:"user_id"`
	RoleId    string `json:"role_id"`
	Password  string `json:"password"  binding:"required"`
	Signature string `json:"signature"`
	Username  string `json:"username"  binding:"required"`
}

func (u User) Login() (User, error) {
	//var user User
	err := global.DB.Table("users").Where("username=?", u.Username).First(&u)
	log.Println(u.UserId)
	return u, err.Error
}

func (u User) Register() (User, error) {
	err := global.DB.Create(&User{
		Password: u.Password,
		Username: u.Username,
	}).Error
	if err != nil {
		return User{}, err
	}
	return User{}, nil
}

// CheckUserExist 判断用户是否存在 如果存在返回true,不存在返回false
func CheckUserExist(username string) (bool, error) {
	var u []User
	global.DB.Find(&u, "username=?", username)
	var count int
	count = len(u)
	return count > 0, nil
}
