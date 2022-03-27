package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")   // 配置文件名
	vp.AddConfigPath("configs/") // 配置文件路径
	vp.SetConfigType("yaml")     // 配置文件类型
	err := vp.ReadInConfig()     // 读取配置文件
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
