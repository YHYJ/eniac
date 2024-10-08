/*
File: define_filemanager.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-23 15:25:26

Description: 文件管理
*/

package general

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
)

// ReadFileKey 读取文件包含关键字的行
//
// 参数：
//   - file: 文件路径
//   - key: 关键字
//
// 返回：
//   - 包含关键字的行的内容
func ReadFileKey(file, key string) string {
	// 打开文件
	text, err := os.Open(file)
	if err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Printf("%s %s %s\n", DangerText(ErrorInfoFlag), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}
	defer text.Close()

	// 创建一个扫描器对象按行遍历
	scanner := bufio.NewScanner(text)
	// 逐行读取，输出指定行
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), key) {
			return scanner.Text()
		}
	}
	return ""
}

// FileExist 判断文件是否存在
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 文件存在返回 true，否则返回 false
func FileExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

// ReadFileLink 如果文件是软链接文件，返回其指向的文件路径
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 软链接文件所指向文件的路径
//   - 报错信息
func ReadFileLink(file string) (string, error) {
	if !FileExist(file) {
		return "", fmt.Errorf("File %s not exist", file)
	}

	fileinfo, err := os.Lstat(file)
	if err != nil {
		return "", err
	}

	if fileinfo.Mode()&os.ModeSymlink == 0 {
		return "", fmt.Errorf("File %s is not a symlink", file)
	}
	link, err := os.Readlink(file)
	if err != nil {
		return "", err
	}
	return link, nil
}

// CreateFile 创建文件，包括其父目录
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 错误信息
func CreateFile(file string) error {
	if FileExist(file) {
		return nil
	}
	// 创建父目录
	parentPath := filepath.Dir(file)
	if err := os.MkdirAll(parentPath, os.ModePerm); err != nil {
		return err
	}
	// 创建文件
	if _, err := os.Create(file); err != nil {
		return err
	}

	return nil
}

// DeleteFile 删除文件，如果目标是文件夹则包括其下所有文件
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 错误信息
func DeleteFile(file string) error {
	if !FileExist(file) {
		return nil
	}
	return os.RemoveAll(file)
}

// GetFileModTime 获取文件的最后修改时间
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 最后修改时间
func GetFileModTime(file string) string {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return ""
	}

	modTime := fileInfo.ModTime()
	return modTime.Format("2006-01-02 15:04:05")
}
