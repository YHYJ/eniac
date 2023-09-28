/*
File: get_system_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 13:37:40

Description: 获取系统信息
*/

package function

import (
	"fmt"
	"strings"

	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/zcalusic/sysinfo"
)

var hostInfo, _ = host.Info()
var pciInfo, _ = ghw.PCI()

// GetBIOSInfo 获取BIOS信息
func GetBIOSInfo(sysInfo sysinfo.SysInfo) (biosInfo map[string]interface{}, err error) {
	biosInfo = make(map[string]interface{})
	biosInfo["BIOSVendor"] = sysInfo.BIOS.Vendor   // BIOS厂商
	biosInfo["BIOSVersion"] = sysInfo.BIOS.Version // BIOS版本
	biosInfo["BIOSDate"] = sysInfo.BIOS.Date       // BIOS日期

	return biosInfo, err
}

// GetBoardInfo 获取主板信息
func GetBoardInfo(sysInfo sysinfo.SysInfo) (boardInfo map[string]interface{}, err error) {
	boardInfo = make(map[string]interface{})
	boardInfo["BoardVendor"] = sysInfo.Board.Vendor   // 主板厂商
	boardInfo["BoardName"] = sysInfo.Board.Name       // 主板名称
	boardInfo["BoardVersion"] = sysInfo.Board.Version // 主板版本

	return boardInfo, err
}

// GetCPUInfo 获取CPU信息
func GetCPUInfo(sysInfo sysinfo.SysInfo, dataUnit string) (cpuInfo map[string]interface{}, err error) {
	cpuInfo = make(map[string]interface{})
	cpuInfo["CPUModel"] = sysInfo.CPU.Model                                               // cpu型号
	cpuInfo["CPUNumber"] = sysInfo.CPU.Cpus                                               // cpu数量
	cpuInfo["CPUCores"] = sysInfo.CPU.Cores                                               // cpu核心数
	cpuInfo["CPUThreads"] = sysInfo.CPU.Threads                                           // cpu线程数
	cpuCache, cpuCacheUnit := DataUnitConvert("KB", dataUnit, float64(sysInfo.CPU.Cache)) // cpu缓存
	cpuInfo["CPUCache"] = fmt.Sprintf("%.2f %s", cpuCache, cpuCacheUnit)                  // cpu缓存

	return cpuInfo, err
}

// GetOSInfo 获取系统信息
func GetOSInfo(sysInfo sysinfo.SysInfo) (osInfo map[string]interface{}, err error) {
	osInfo = make(map[string]interface{})
	osInfo["OS"] = UpperStringFirstChar(sysInfo.OS.Name)         // 操作系统
	osInfo["Arch"] = sysInfo.OS.Architecture                     // 系统架构
	osInfo["Kernel"] = sysInfo.Kernel.Release                    // 内核版本
	osInfo["Platform"] = UpperStringFirstChar(sysInfo.OS.Vendor) // 平台
	osInfo["Hostname"] = hostInfo.Hostname                       // 主机名
	osInfo["TimeZone"] = sysInfo.Node.Timezone                   // 时区

	return osInfo, err
}

// GetProcessInfo 获取进程信息
func GetProcessInfo() (procsInfo map[string]interface{}, err error) {
	procsInfo = make(map[string]interface{})
	procsInfo["Process"] = hostInfo.Procs // 进程数

	return procsInfo, err
}

// GetProductInfo 获取产品信息
func GetProductInfo(sysInfo sysinfo.SysInfo) (productInfo map[string]interface{}, err error) {
	productInfo = make(map[string]interface{})
	productInfo["ProductVendor"] = sysInfo.Product.Vendor // 产品厂商
	productInfo["ProductName"] = sysInfo.Product.Name     // 产品名称

	return productInfo, err
}

// GetStorageInfo 获取存储设备信息
func GetStorageInfo(address string) (storageInfo map[string]interface{}, err error) {
	block, err := ghw.Block()
	if err != nil {
		fmt.Println("Failed to get block storage information:", err)
		return
	}
	storageInfo = make(map[string]interface{})
	for index, disk := range block.Disks {
		storageValue := make(map[string]interface{})
		if disk.SizeBytes > 0 {
			storageValue["StorageName"] = disk.Name
			storageValue["StorageDriver"] = disk.StorageController
			storageValue["StorageVendor"] = pciInfo.GetDevice(address).Vendor.Name
			storageValue["StorageModel"] = disk.Model
			storageValue["StorageType"] = disk.DriveType
			storageValue["StorageRemovable"] = disk.IsRemovable
			storageValue["StorageSerial"] = disk.SerialNumber
			storageSize, storageSizeUnit := DataUnitConvert("B", "TB", float64(disk.SizeBytes))
			storageValue["StorageSize"] = fmt.Sprintf("%.1f %s", storageSize, storageSizeUnit)
			storageInfo[fmt.Sprintf("%s%d", "Storage", index)] = storageValue
		}
	}

	return storageInfo, err
}

// GetTimeInfo 获取时间信息
func GetTimeInfo() (timeInfo map[string]interface{}, err error) {
	timeInfo = make(map[string]interface{})
	timeInfo["BootTime"] = Uint2TimeString(hostInfo.BootTime) // 系统启动时间
	day, hour, minute, second := Second2DayHourMinuteSecond(hostInfo.Uptime)
	result := fmt.Sprintf("%vd %vh %vm %vs", day, hour, minute, second)
	timeInfo["Uptime"] = result // 系统运行时间
	starttimeArgs := []string{"time"}
	StartTime, err := RunCommandGetResult("systemd-analyze", starttimeArgs)
	if err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", err)
	}
	timeInfo["StartTime"] = strings.Split(strings.Split(StartTime, "\n")[0], "= ")[1] // 系统启动用时

	return timeInfo, err
}
