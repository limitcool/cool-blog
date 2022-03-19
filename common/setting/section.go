package setting

import "time"

// 服务器配置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// 应用配置
type AppSettingS struct {
	DefaultPageSize         int
	MaxPageSize             int
	LogSavePath             string
	LogFileName             string
	LogFileExt              string
	UploadSavePath          string
	UploadServerUrl         string
	UploadImageMaxSize      int
	UploadImageAllowExts    []string
	UploadMarkdownAllowExts []string
}

// 数据库配置
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	LogMode      string
	MaxIdleConns int
	MaxOpenConns int
}

// JWT配置
type JwtSettingS struct {
	Secret string        // 密钥
	Issuer string        // JWT的签发者
	Expire time.Duration // 所签发的 JWT 过期时间，过期时间必须大于签发时间。
}

// 解码key到struct
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
