package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	Vp := viper.New()
	Vp.SetConfigName("config")   // 配置文件名
	Vp.AddConfigPath("configs/") // 配置文件路径
	Vp.SetConfigType("yaml")     // 配置文件类型
	err := Vp.ReadInConfig()     // 读取配置文件
	if err != nil {
		return nil, err
	}
	return &Setting{Vp}, nil
}
