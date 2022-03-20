package dao

import (
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/model"
	"log"
)

func (d *Dao) Login(username, password string) (model.User, error) {
	user := model.User{Username: username, Password: password}
	return user.Login()
}

func (d *Dao) Register(username, password string) (model.User, error) {
	user := model.User{Username: username, Password: password}
	return user.Register()
}

func (d *Dao) GetRoleId(username string) string {
	user := model.User{Username: username}
	global.DB.Table("users").Where("username = ?", username).First(&user)
	log.Println("21:", user.RoleId)
	return user.RoleId
}
