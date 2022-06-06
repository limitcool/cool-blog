package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 创建基本类型
type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	ID        uint           `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"` // 时间日期直接按照字符串处理即可
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	//gorm.Model
}
