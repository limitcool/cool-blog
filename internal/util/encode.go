package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// Md5 对字符串进行MD5算法
func Md5(inputStr string) string {
	h := md5.New()
	h.Write([]byte(inputStr))
	str := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return str
}

// Base64EncodeWithString base64字符加密
func Base64EncodeWithString(input string) string {

	return base64.StdEncoding.EncodeToString([]byte(input))
}

// Base64DecodeWithString base64字符解密
func Base64DecodeWithString(input string) (string, error) {

	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(data), err
}
