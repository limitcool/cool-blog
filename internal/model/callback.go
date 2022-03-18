package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

// 新建时自动添加时间 回调
func UpTimestampForCreateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		nowtime := time.Now()
		for _, field := range db.Statement.Schema.Fields {
			if field.Name == "created_at" {
				var ctx context.Context
				// 设置字段的 value
				field.Set(ctx, db.Statement.ReflectValue, nowtime)
			}
		}
	}
}

// 修改内容时回调
func TimeStampForUpdateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		nowtime := time.Now()
		for _, field := range db.Statement.Schema.Fields {
			if field.Name == "updated_at" {
				var ctx context.Context
				// 设置字段的 value
				field.Set(ctx, db.Statement.ReflectValue, nowtime)
			}
		}
	}
}
