package dao

import (
	"github.com/limitcool/blog/internal/model"
)

func (d *Dao) Login(username, password string) (model.User, error) {
	user := model.User{Username: username, Password: password}
	return user.Login()
}

func (d *Dao) Register(username, password string) (model.User, error) {
	user := model.User{Username: username, Password: password}
	return user.Register()
}
