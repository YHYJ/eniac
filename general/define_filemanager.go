/*
File: define_filemanager.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-23 15:25:26

Description: 文件操作
*/

package general

import (
	"os"
	"strings"
)

// FileExist 判断文件是否存在
func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// CreateFile 创建文件，如果其父目录不存在则创建父目录
func CreateFile(filePath string) error {
	if FileExist(filePath) {
		return nil
	}
	// 截取filePath的父目录
	parentPath := filePath[:strings.LastIndex(filePath, "/")]
	if err := os.MkdirAll(parentPath, os.ModePerm); err != nil {
		return err
	}
	_, err := os.Create(filePath)
	return err
}

// DeleteFile 删除文件
func DeleteFile(filePath string) error {
	if !FileExist(filePath) {
		return nil
	}
	return os.Remove(filePath)
}
