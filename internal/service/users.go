package service

import (
	"github.com/limitcool/blog/internal/model"
	"github.com/limitcool/blog/internal/pkg"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (svc *Service) CheckLogin(param *LoginRequest) error {
	log.Println(param.Username, param.Password)
	login, err := svc.dao.Login(param.Username, param.Password)
	log.Println(login.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户名不存在!")
		}
		return err
	}
	if login.UserId > 0 {
		// 判断密码是否正确
		if param.Password == login.Password {
			token, _ := pkg.GenerateToken(param.Username, param.Password)
			log.Println("生成的Token为:", token)
			return nil
		} else {
			return errors.New("密码错误!")
		}

	}
	return errors.New("用户名不存在")
}

func (svc *Service) CheckRegister(param *RegisterRequest) error {
	if ok, _ := model.CheckUserExist(param.Username); ok {
		return errors.New("用户名已存在")
	} else {
		user, err := svc.dao.Register(param.Username, param.Password)
		log.Println(user)
		if err != nil {
			return errors.New("注册失败")
			return err
		} else {
			return errors.New("注册成功")
		}
	}

}
