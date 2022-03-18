package model

import (
	"gorm.io/gorm"
	"time"
)

// 创建基本类型
type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	CreatedAt time.Time `json:"created_at"` // 时间日期直接按照字符串处理即可
	UpdatedAt time.Time `json:"updated_at"`
}
