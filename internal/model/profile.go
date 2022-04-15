package model

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	ID        uint           `gorm:"primary_key;  json:"id"`
	CreatedAt time.Time      `json:"created_at"` // 时间日期直接按照字符串处理即可
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Desc      string         `gorm:"default:'';" json:"desc"`
	Img       string         `gorm:"default:'';" json:"img"`
	UserId    uint           `json:"user_id"`
}

func (p Profile) TableName() string {
	return "profile"
}
