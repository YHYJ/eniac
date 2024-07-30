//go:build darwin

/*
File: define_spider_darwin.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-11 14:50:52

Description: 信息抓取器
*/

package general

import (
	"strconv"

	"github.com/gookit/color"
	"github.com/zcalusic/sysinfo"
)

// GetStorageInfo 获取存储设备信息
//
// 返回：
//   - 存储设备信息
func GetStorageInfo() map[string]interface{} {
	storageInfo := make(map[string]interface{})
	index := 1 // 排除编号为0的虚拟设备
	for _, disk := range blockData.Disks {
		storageValue := make(map[string]interface{})
		if disk.SizeBytes > 0 && disk.DriveType.String() != "virtual" {
			storageValue["StorageName"] = disk.Name
			storageValue["StorageDriver"] = disk.StorageController.String()
			storageValue["StorageVendor"] = disk.Vendor
			storageValue["StorageModel"] = disk.Model
			storageValue["StorageType"] = disk.DriveType.String()
			storageValue["StorageRemovable"] = strconv.FormatBool(disk.IsRemovable)
			storageValue["StorageSerial"] = disk.SerialNumber
			storageSize, storageSizeUnit := Human(float64(disk.SizeBytes), "B")
			storageValue["StorageSize"] = color.Sprintf("%.1f %s", storageSize, storageSizeUnit)
			storageInfo[color.Sprintf("%d", index)] = storageValue
			index += 1
		}
	}

	return storageInfo
}

// GetOSInfo 获取系统信息
//
// 参数：
//   - sysInfo: 总的系统信息 (System Info)
//
// 返回：
//   - 系统信息 (OS Info)
func GetOSInfo(sysInfo sysinfo.SysInfo) map[string]interface{} {
	osInfo := make(map[string]interface{})

	// 需要额外步骤获取的信息
	osCode := FindOSCode(hostData.PlatformVersion) // 系统代号
	timeZone := GetTimeZoneOriginal()              // 时区

	osInfo["OS"] = color.Sprintf("%s %s", osCode, hostData.PlatformVersion) // 操作系统
	osInfo["Arch"] = hostData.KernelArch                                    // 系统架构
	osInfo["CurrentKernel"] = hostData.KernelVersion                        // 当前内核版本
	osInfo["Platform"] = UpperFirstChar(hostData.Platform)                  // 平台
	osInfo["Hostname"] = hostData.Hostname                                  // 主机名
	osInfo["TimeZone"] = timeZone                                           // 时区

	return osInfo
}
