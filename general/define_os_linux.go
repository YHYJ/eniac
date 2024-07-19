//go:build linux

/*
File: define_os_linux.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-04 11:03:33

Description: 操作系统信息
*/

package general

import (
	"strings"
)

// GetLatestKernelVersion 获取本地最新内核版本
//
// 返回：
//   - 最新内核版本号
func GetLatestKernelVersion() string {
	var latestKernelVersion string

	id, _ := GetSystemID()
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

// getLatestKernelVersionForArch 获取本地最新内核版本，Arch 系专用
//
// 返回：
//   - 最新内核版本号
func getLatestKernelVersionForArch() string {
	var latestKernelVersion string

	args := []string{"-Q", "linux"}
	kernelVersion, _, _ := RunCommandToBuffer("pacman", args)
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
