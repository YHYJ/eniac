//go:build linux

/*
File: define_system_linux.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-04 11:03:33

Description: 系统相关方法
*/

package general

import (
	"os"
	"strings"
)

// GetLatestKernelVersion 获取本地最新内核版本
//
// 返回：
//   - 最新内核版本号
func GetLatestKernelVersion() string {
	var latestKernelVersion string

	releaseInfoFile := "/etc/os-release"
	id, _ := readOSRelease(releaseInfoFile)
	switch id {
	case "arch":
		latestKernelVersion = getLatestKernelVersionForArch()
	case "debian":
		latestKernelVersion = getLatestKernelVersionForDebian()
	default:
		latestKernelVersion = getLatestKernelVersionForUnknown()
	}

	return latestKernelVersion
}

// readOSRelease 获取系统 ID_LIKE 或 ID
//
// 参数：
//   - 记录有系统发行信息的文件
//
// 返回：
//   - 系统 ID_LIKE 或 ID
//   - 错误信息
func readOSRelease(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var id string

	id = ReadFileKey(filePath, "ID_LIKE=")
	if id != "" {
		id = strings.TrimPrefix(id, "ID_LIKE=")
	} else {
		id = ReadFileKey(filePath, "ID=")
		id = strings.TrimPrefix(id, "ID=")
	}
	id = strings.Trim(id, `"`)

	return id, nil
}

// getLatestKernelVersionForArch 获取本地最新内核版本，Arch 系专用
//
// 返回：
//   - 最新内核版本号
func getLatestKernelVersionForArch() string {
	var latestKernelVersion string

	command := "pacman"
	args := []string{"-Q", "linux"}
	kernelVersion, _ := RunCommandGetResult(command, args)
	if len(kernelVersion) != 0 {
		latestKernelVersion = strings.Split(kernelVersion, " ")[1]
	}

	return latestKernelVersion
}

// getLatestKernelVersionForDebian 获取本地最新内核版本，Debian 系专用
//
// 返回：
//   - 最新内核版本号
func getLatestKernelVersionForDebian() string {
	var latestKernelVersion string

	// TODO: 待实现 <07-06-24, YJ> //

	return latestKernelVersion
}

// getLatestKernelVersionForUnknown 获取本地最新内核版本，未支持系统专用
//
// 返回：
//   - 最新内核版本号
func getLatestKernelVersionForUnknown() string {
	var latestKernelVersion string = "Unknown"

	return latestKernelVersion
}
