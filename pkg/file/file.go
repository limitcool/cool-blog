package file

import (
	"fmt"
	"os"
)

// 检查文件是否不存在,不存在返回true
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 检查目录是否存在,不存在就创建目录
func IsNotExistMkdir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}
func MustOpen(filename, filepath string) (*os.File, error) {
	// 获取当前文件路径并将其存入dir
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// 文件最终保存的绝对路径
	src := dir + "/" + filepath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("文件没有权限:%s", err)
	}
	// 创建目录
	err = IsNotExistMkdir(src)
	if err != nil {
		return nil, fmt.Errorf("创建目录失败:%s", err)
	}
	// 创建文件
	f, err := Open(src+filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("创建文件失败:%s", err)
	}
	return f, err
}

// 检查是否具有文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

// 打开一个文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
