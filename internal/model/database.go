package model

import (
	"fmt"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	var logMode logger.Interface
	if global.DatabaseSetting.LogMode == "info" {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = nil
	}
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)), &gorm.Config{Logger: logMode})

	db.Callback().Create().Register("upTimestampForCreateCallback", UpTimestampForCreateCallback)
	// 为 UP流程注册一个回调
	db.Callback().Update().Register("UpdateCallback", TimeStampForUpdateCallback)
	if err != nil {
		log.Println("err!", err)
	}
	if err != nil {
		return nil, err
	}
	sqlDb, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db, nil
}
