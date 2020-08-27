package util

import (
	"os"
	"strings"
)

var (
	// 系统路径分隔符
	FileSep = string(os.PathSeparator)
)

// 获取一个路径的父目录地址
func GetParentDir(path string) string {
	path = strings.Trim(path, " ")
	if strings.HasSuffix(path, "/") || strings.HasSuffix(path, FileSep) {
		path = path[0 : len(path)-1]
	}
	index := strings.LastIndex(path, "/")
	if -1 == index {
		index = strings.LastIndex(path, FileSep)
	}
	return path[0:index]
}

// 根据路径创建文件夹
func MkdirAll(path string) {
	_ = os.MkdirAll(path, 0777)
}

// 创建一个文件的父目录
func MkdirParent(path string) {
	parent := GetParentDir(path)
	if !PathExists(parent) {
		MkdirAll(parent)
	}
}

// 判断一个路径是否存在
func PathExists(path string) bool {
	stat, _ := os.Stat(path)
	return stat != nil
}
