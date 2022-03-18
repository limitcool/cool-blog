package dao

import (
	"github.com/limitcool/blog/internal/model"
)

type Dao struct {
}

// 调用 auth.Get()
func (d *Dao) GetAuth(username, password string) (model.Auth, error) {
	auth := model.Auth{Username: username, Password: password}
	return auth.Get()
}
