package model

import (
	"github.com/limitcool/blog/common/snowflake"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/util"
	"log"
)

type User struct {
	BaseModel
	RoleId      string `json:"role_id"`
	Password    string `json:"password"  binding:"required"`
	Signature   string `json:"signature"`
	Username    string `gorm:"unique" json:"username"  binding:"required"`
	SnowFlakeId int64
	Profile     Profile    `json:"profile"`
	Articles    []Articles // 文章
}

func (u User) name() {

}
func (u User) TableName() string {
	return "users"
}
func (u User) Login() (User, error) {
	//var user User
	err := global.DB.Table("users").Where("username=?", u.Username).First(&u)
	log.Println(u.ID)
	return u, err.Error
}

func (u User) Register() (User, error) { // 对密码进行MD5加密
	md5Password := util.Md5(u.Password)
	// 生成雪花id
	snowflakeId := snowflake.GenerateSnowFlakeId()
	profile := Profile{
		Desc: "default",
		Img:  "default",
	}
	log.Println("41", u.Profile.Desc, u.Profile.Img)
	user := User{
		Password:    md5Password,
		Username:    u.Username,
		SnowFlakeId: snowflakeId,
		Profile:     profile,
	}

	err := global.DB.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// CheckUserExist 判断用户是否存在 如果存在返回true,不存在返回false
func CheckUserExist(username string) (bool, error) {
	var u []User
	global.DB.Find(&u, "username=?", username)
	var count int
	count = len(u)
	return count > 0, nil
}

// GetProfile 获取用户个人资料
//func (u *User) GetProfile() Profile {
//	global.DB.Preload("Profile").Where("id = ?", u.ID).Find(&u)
//	return u.Profile
//}
