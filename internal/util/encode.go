package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

//对字符串进行MD5加密
func Md5(inputStr string) string {
	h := md5.New()
	h.Write([]byte(inputStr)) // 需要加密的字符串为
	str := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return str
}

//base64字符加密
func Base64EncodeWithString(input string) string {

	return base64.StdEncoding.EncodeToString([]byte(input))
}

//base64字符解密
func Base64DecodeWithString(input string) (string, error) {

	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(data), err
}
