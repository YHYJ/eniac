/*
File: get_systeminfo.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 13:37:40

Description: 获取系统信息
*/

package function

import (
	"github.com/shirou/gopsutil/host"
	"github.com/zcalusic/sysinfo"
)

// BIOSInfoStruct BIOS信息结构体
type BIOSInfoStruct struct {
	BIOSVendor  string `json:"bios_vendor"`  // bios厂商
	BIOSVersion string `json:"bios_version"` // bios版本
	BIOSDate    string `json:"bios_date"`    // bios日期
}

// BoardInfoStruct 主板信息结构体
type BoardInfoStruct struct {
	BoardVendor  string `json:"board_vendor"`  // 主板厂商
	BoardName    string `json:"board_name"`    // 主板名称
	BoardVersion string `json:"board_version"` // 主板版本
}

// CPUInfoStruct CPU信息结构体
type CPUInfoStruct struct {
	CPUModel   string `json:"cpu_model"`   // cpu型号
	CPUNumber  uint   `json:"cpu_number"`  // cpu数量
	CPUCores   uint   `json:"cpu_cores"`   // cpu核心数
	CPUThreads uint   `json:"cpu_threads"` // cpu线程数
	CPUCache   uint   `json:"cpu_cache"`   // cpu缓存
}

// OSInfoStruct 系统信息结构体
type OSInfoStruct struct {
	OS       string `json:"os"`        // 操作系统
	Arch     string `json:"arch"`      // 系统架构
	Kernel   string `json:"kernel"`    // 内核版本
	Platform string `json:"platform"`  // 平台
	Hostname string `json:"hostname"`  // 主机名
	TimeZone string `json:"time_zone"` // 时区
}

// ProcsInfoStruct 进程信息结构体
type ProcsInfoStruct struct {
	Procs uint64 `json:"procs"` // 进程数
}

// ProductInfoStruct 产品信息结构体
type ProductInfoStruct struct {
	ProductVendor string `json:"product_vendor"` // 产品厂商
	ProductName   string `json:"product_name"`   // 产品名称
}

// StorageInfoStruct 存储设备信息结构体
type StorageInfoStruct struct {
	StorageList []sysinfo.StorageDevice `json:"storage_list"` // 存储设备列表
}

// TimeInfoStruct 时间信息结构体
type TimeInfoStruct struct {
	BootTime uint64 `json:"boot_time"` // 系统启动时间
	Uptime   uint64 `json:"uptime"`    // 系统运行时间
}

var hostInfo, _ = host.Info()

// GetBIOSInfo 获取BIOS信息
func GetBIOSInfo(sysInfo sysinfo.SysInfo) (biosInfo BIOSInfoStruct, err error) {
	biosInfo.BIOSVendor = sysInfo.BIOS.Vendor
	biosInfo.BIOSVersion = sysInfo.BIOS.Version
	biosInfo.BIOSDate = sysInfo.BIOS.Date

	return biosInfo, err
}

// GetBoardInfo 获取主板信息
func GetBoardInfo(sysInfo sysinfo.SysInfo) (boardInfo BoardInfoStruct, err error) {
	boardInfo.BoardVendor = sysInfo.Board.Vendor
	boardInfo.BoardName = sysInfo.Board.Name
	boardInfo.BoardVersion = sysInfo.Board.Version

	return boardInfo, err
}

// GetCPUInfo 获取CPU信息
func GetCPUInfo(sysInfo sysinfo.SysInfo) (cpuInfo CPUInfoStruct, err error) {
	cpuInfo.CPUModel = sysInfo.CPU.Model
	cpuInfo.CPUNumber = sysInfo.CPU.Cpus
	cpuInfo.CPUCores = sysInfo.CPU.Cores
	cpuInfo.CPUThreads = sysInfo.CPU.Threads
	cpuInfo.CPUCache = sysInfo.CPU.Cache

	return cpuInfo, err
}

// GetOSInfo 获取系统信息
func GetOSInfo(sysInfo sysinfo.SysInfo) (osInfo OSInfoStruct, err error) {
	osInfo.OS = upperStringFirstChar(sysInfo.OS.Name)
	osInfo.Arch = sysInfo.OS.Architecture
	osInfo.Kernel = sysInfo.Kernel.Release
	osInfo.Platform = upperStringFirstChar(sysInfo.OS.Vendor)
	osInfo.Hostname = hostInfo.Hostname
	osInfo.TimeZone = sysInfo.Node.Timezone

	return osInfo, err
}

// GetProcsInfo 获取进程信息
func GetProcsInfo() (procsInfo ProcsInfoStruct, err error) {
	procsInfo.Procs = hostInfo.Procs

	return procsInfo, err
}

// GetProductInfo 获取产品信息
func GetProductInfo(sysInfo sysinfo.SysInfo) (productInfo ProductInfoStruct, err error) {
	productInfo.ProductVendor = sysInfo.Product.Vendor
	productInfo.ProductName = sysInfo.Product.Name

	return productInfo, err
}

// GetStorageInfo 获取存储设备信息
func GetStorageInfo(sysInfo sysinfo.SysInfo) (storageInfo StorageInfoStruct, err error) {
	storageInfo.StorageList = sysInfo.Storage

	return storageInfo, err
}

// GetTimeInfo 获取时间信息
func GetTimeInfo() (timeInfo TimeInfoStruct, err error) {
	timeInfo.BootTime = hostInfo.BootTime
	timeInfo.Uptime = hostInfo.Uptime

	return timeInfo, err
}
