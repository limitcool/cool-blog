package global

import (
	"github.com/limitcool/blog/internal/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JwtSetting      *setting.JwtSettingS
	PaySetting      *setting.Pay
	RedisSetting    *setting.Redis
)
