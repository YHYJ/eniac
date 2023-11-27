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

// GetBIOSInfo 获取 BIOS 信息
//
// 参数：
//   - sysInfo: 总的系统信息
//
// 返回：
//   - BIOS 信息
func GetBIOSInfo(sysInfo sysinfo.SysInfo) map[string]interface{} {
	biosInfo := make(map[string]interface{})
	biosInfo["BIOSVendor"] = sysInfo.BIOS.Vendor   // BIOS 厂商
	biosInfo["BIOSVersion"] = sysInfo.BIOS.Version // BIOS 版本
	biosInfo["BIOSDate"] = sysInfo.BIOS.Date       // BIOS 日期

	return biosInfo
}

// GetBoardInfo 获取主板信息
//
// 参数：
//   - sysInfo: 总的系统信息
//
// 返回：
//   - 主板信息
func GetBoardInfo(sysInfo sysinfo.SysInfo) map[string]interface{} {
	boardInfo := make(map[string]interface{})
	boardInfo["BoardVendor"] = sysInfo.Board.Vendor   // 主板厂商
	boardInfo["BoardName"] = sysInfo.Board.Name       // 主板名称
	boardInfo["BoardVersion"] = sysInfo.Board.Version // 主板版本

	return boardInfo
}

// GetCPUInfo 获取 CPU 信息
//
// 参数：
//   - sysInfo: 总的系统信息
//   - dataUnit: 存储数据单位
//
// 返回：
//   - CPU 信息
func GetCPUInfo(sysInfo sysinfo.SysInfo, dataUnit string) map[string]interface{} {
	cpuInfo := make(map[string]interface{})
	cpuInfo["CPUModel"] = sysInfo.CPU.Model                                    // CPU 型号
	cpuInfo["CPUNumber"] = sysInfo.CPU.Cpus                                    // CPU 数量
	cpuInfo["CPUCores"] = sysInfo.CPU.Cores                                    // CPU 核心数
	cpuInfo["CPUThreads"] = sysInfo.CPU.Threads                                // CPU 线程数
	cpuCache, cpuCacheUnit := general.Human(float64(sysInfo.CPU.Cache), "KiB") // CPU 缓存
	cpuInfo["CPUCache"] = fmt.Sprintf("%.0f %s", cpuCache, cpuCacheUnit)

	return cpuInfo
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
	osInfo["OS"] = general.UpperStringFirstChar(sysInfo.OS.Name)         // 操作系统
	osInfo["Arch"] = sysInfo.OS.Architecture                             // 系统架构
	osInfo["Kernel"] = sysInfo.Kernel.Release                            // 内核版本
	osInfo["Platform"] = general.UpperStringFirstChar(sysInfo.OS.Vendor) // 平台
	osInfo["Hostname"] = hostData.Hostname                               // 主机名
	osInfo["TimeZone"] = sysInfo.Node.Timezone                           // 时区

	return osInfo
}

// GetProductInfo 获取产品信息
//
// 参数：
//   - sysInfo: 总的系统信息
//
// 返回：
//   - 产品信息
func GetProductInfo(sysInfo sysinfo.SysInfo) map[string]interface{} {
	productInfo := make(map[string]interface{})
	productInfo["ProductVendor"] = sysInfo.Product.Vendor // 产品厂商
	productInfo["ProductName"] = sysInfo.Product.Name     // 产品名称

	return productInfo
}

// GetTimeInfo 获取时间信息
//
// 返回：
//   - 时间信息
//   - 错误信息
func GetTimeInfo() (map[string]interface{}, error) {
	timeInfo := make(map[string]interface{})
	timeInfo["BootTime"] = general.UnixTime2TimeString(hostData.BootTime) // 系统启动时间
	day, hour, minute, second := general.UnixTime2DayHourMinuteSecond(hostData.Uptime)
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
