//go:build linux

/*
File: define_system_linux.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-07-19 15:27:08

Description: 操作系统相关方法
*/
package general

import (
	"os"
	"strings"
)

const releaseFile = "/etc/os-release"

// GetSystemID 获取系统 ID_LIKE 或 ID
//
// 返回：
//   - 系统 ID_LIKE 或 ID
//   - 错误信息
func GetSystemID() (string, error) {
	file, err := os.Open(releaseFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var id string

	id = ReadFileKey(releaseFile, "ID_LIKE=")
	if id != "" {
		id = strings.TrimPrefix(id, "ID_LIKE=")
	} else {
		id = ReadFileKey(releaseFile, "ID=")
		id = strings.TrimPrefix(id, "ID=")
	}
	id = strings.Trim(id, `"`)

	return id, nil
}
