/*
File: get_system_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 13:37:40

Description: 获取系统信息
*/

package cli

import (
	"fmt"
	"strings"

	"github.com/yhyj/eniac/general"
	"github.com/zcalusic/sysinfo"
)

// GetBIOSInfo 获取BIOS信息
func GetBIOSInfo(sysInfo sysinfo.SysInfo) map[string]interface{} {
	biosInfo := make(map[string]interface{})
	biosInfo["BIOSVendor"] = sysInfo.BIOS.Vendor   // BIOS厂商
	biosInfo["BIOSVersion"] = sysInfo.BIOS.Version // BIOS版本
	biosInfo["BIOSDate"] = sysInfo.BIOS.Date       // BIOS日期

	return biosInfo
}

// GetBoardInfo 获取主板信息
func GetBoardInfo(sysInfo sysinfo.SysInfo) map[string]interface{} {
	boardInfo := make(map[string]interface{})
	boardInfo["BoardVendor"] = sysInfo.Board.Vendor   // 主板厂商
	boardInfo["BoardName"] = sysInfo.Board.Name       // 主板名称
	boardInfo["BoardVersion"] = sysInfo.Board.Version // 主板版本

	return boardInfo
}

// GetCPUInfo 获取CPU信息
func GetCPUInfo(sysInfo sysinfo.SysInfo, dataUnit string) map[string]interface{} {
	cpuInfo := make(map[string]interface{})
	cpuInfo["CPUModel"] = sysInfo.CPU.Model                                               // cpu型号
	cpuInfo["CPUNumber"] = sysInfo.CPU.Cpus                                               // cpu数量
	cpuInfo["CPUCores"] = sysInfo.CPU.Cores                                               // cpu核心数
	cpuInfo["CPUThreads"] = sysInfo.CPU.Threads                                           // cpu线程数
	cpuCache, cpuCacheUnit := general.DataUnitConvert("KB", dataUnit, float64(sysInfo.CPU.Cache)) // cpu缓存
	cpuInfo["CPUCache"] = fmt.Sprintf("%.1f%s", cpuCache, cpuCacheUnit)

	return cpuInfo
}

// GetOSInfo 获取系统信息
func GetOSInfo(sysInfo sysinfo.SysInfo) map[string]interface{} {
	osInfo := make(map[string]interface{})
	osInfo["OS"] = general.UpperStringFirstChar(sysInfo.OS.Name)         // 操作系统
	osInfo["Arch"] = sysInfo.OS.Architecture                     // 系统架构
	osInfo["Kernel"] = sysInfo.Kernel.Release                    // 内核版本
	osInfo["Platform"] = general.UpperStringFirstChar(sysInfo.OS.Vendor) // 平台
	osInfo["Hostname"] = hostData.Hostname                       // 主机名
	osInfo["TimeZone"] = sysInfo.Node.Timezone                   // 时区

	return osInfo
}

// GetProcessInfo 获取进程信息
func GetProcessInfo() map[string]interface{} {
	procsInfo := make(map[string]interface{})
	procsInfo["Process"] = hostData.Procs // 进程数

	return procsInfo
}

// GetProductInfo 获取产品信息
func GetProductInfo(sysInfo sysinfo.SysInfo) map[string]interface{} {
	productInfo := make(map[string]interface{})
	productInfo["ProductVendor"] = sysInfo.Product.Vendor // 产品厂商
	productInfo["ProductName"] = sysInfo.Product.Name     // 产品名称

	return productInfo
}

// GetTimeInfo 获取时间信息
func GetTimeInfo() (map[string]interface{}, error) {
	timeInfo := make(map[string]interface{})
	timeInfo["BootTime"] = general.Uint2TimeString(hostData.BootTime) // 系统启动时间
	day, hour, minute, second := general.Second2DayHourMinuteSecond(hostData.Uptime)
	result := fmt.Sprintf("%vd %vh %vm %vs", day, hour, minute, second)
	timeInfo["Uptime"] = result // 系统运行时间
	starttimeArgs := []string{"time"}
	StartTime, err := general.RunCommandGetResult("systemd-analyze", starttimeArgs)
	if err != nil {
		return nil, err
	}
	timeInfo["StartTime"] = strings.Split(strings.Split(StartTime, "\n")[0], "= ")[1] // 系统启动用时

	return timeInfo, nil
}
