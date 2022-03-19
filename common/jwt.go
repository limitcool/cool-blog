package common

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/limitcool/blog/global"
	"github.com/limitcool/blog/internal/util"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

// GetJwtSecret 获取jwtSecret 并以byte数组返回
func GetJwtSecret() []byte {
	return []byte(global.JwtSetting.Secret)
}

// GenerateToken 生成token
func GenerateToken(username, password string) (string, error) {
	// 生成当前时间
	nowTime := time.Now()
	// 生成过期时间: 当前时间+配置参数
	expireTime := nowTime.Add(global.JwtSetting.Expire * time.Second)
	fmt.Println(expireTime)
	claims := Claims{
		// MD5加密Appkey
		Username: util.Md5(username),
		// MD5加密appSecret
		Password: util.Md5(password),
		// RegisteredClaims结构体
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间
			ExpiresAt: jwt.NewNumericDate(expireTime),
			// 设置签发人
			Issuer: global.JwtSetting.Issuer,
		},
	}
	// 使用HS256算法生成token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用签名密钥进行加签
	token, err := tokenClaims.SignedString(GetJwtSecret())
	return token, err
}

// ParseToken 解析和校验token
func ParseToken(token string) (*Claims, error) {
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})
	if err != nil {
		fmt.Println("55:", err)
		return nil, err
	}
	if tokenClaims != nil {
		// 校验token
		if Claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return Claims, nil
		}
	}
	return nil, err
}
