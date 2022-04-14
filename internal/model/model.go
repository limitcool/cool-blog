package model

import (
	"gorm.io/gorm"
)

// BaseModel 创建基本类型
type BaseModel struct {
	*gorm.DB `gorm:"-" json:"-"`
	//Id        uint           `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	//CreatedAt time.Time      `json:"created_at"` // 时间日期直接按照字符串处理即可
	//UpdatedAt time.Time      `json:"updated_at"`
	//DeletedAt gorm.DeletedAt `gorm:"index"`
	gorm.Model
}
