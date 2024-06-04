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
	"bufio"
	"os"
	"os/user"
	"path"

	"strconv"
	"strings"

	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/gookit/color"
	"github.com/zcalusic/sysinfo"
)

var (
	blockData, _ = ghw.Block()         // 存储设备信息
	loadData, _  = load.Avg()          // 系统负载信息
	memData, _   = mem.VirtualMemory() // 内存信息
	hostData, _  = host.Info()         // 主机信息
	userData, _  = user.Current()      // 用户信息
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
			storageValue["StorageSize"] = color.Sprintf("%.2f %s", storageSize, storageSizeUnit)
			storageInfo[color.Sprintf("%d", index)] = storageValue
			index += 1
		}
	}

	return storageInfo
}

// GetLoadInfo 获取负载信息
//
// 返回：
//   - 系统负载信息
func GetLoadInfo() map[string]interface{} {
	loadInfo := make(map[string]interface{})
	loadInfo["Load1"] = loadData.Load1   // 1分钟内的负载
	loadInfo["Load5"] = loadData.Load5   // 5分钟内的负载
	loadInfo["Load15"] = loadData.Load15 // 15分钟内的负载
	loadInfo["Process"] = hostData.Procs // 进程数

	return loadInfo
}

// GetMemoryInfo 获取内存信息
//
// 参数：
//   - dataUnit: 存储数据单位
//   - percentUnit: 百分比数据单位
//
// 返回：
//   - 内存信息
func GetMemoryInfo(dataUnit string, percentUnit string) map[string]interface{} {
	// 内存数据
	memTotal, memTotalUnit := Human(float64(memData.Total), "B")
	memUsed, memUsedUnit := Human(float64(memData.Used), "B")
	memUsedPercent, _ := Human(float64(memData.UsedPercent), percentUnit)
	memFree, memFreeUnit := Human(float64(memData.Free), "B")
	memShared, memSharedUnit := Human(float64(memData.Shared), "B")
	memBuffCache, memBuffCacheUnit := Human(float64(memData.Buffers+memData.Cached), "B")
	memAvail, memAvailUnit := Human(float64(memData.Available), "B")

	// 使用冒泡排序找出最大值用以组装格式字符串
	memData := []float64{memTotal, memUsed, memUsedPercent, memFree, memShared, memBuffCache, memAvail}
	BubbleSort(memData)
	formatString := "%.2f %s"

	memoryInfo := make(map[string]interface{})
	memoryInfo["MemoryTotal"] = color.Sprintf(formatString, memTotal, memTotalUnit)             // 内存总量
	memoryInfo["MemoryUsed"] = color.Sprintf(formatString, memUsed, memUsedUnit)                // 已用内存
	memoryInfo["MemoryUsedPercent"] = color.Sprintf("%.1f%s", memUsedPercent, percentUnit)      // 内存使用率
	memoryInfo["MemoryFree"] = color.Sprintf(formatString, memFree, memFreeUnit)                // 空闲内存
	memoryInfo["MemoryShared"] = color.Sprintf(formatString, memShared, memSharedUnit)          // 共享内存
	memoryInfo["MemoryBuffCache"] = color.Sprintf(formatString, memBuffCache, memBuffCacheUnit) // 缓存内存
	memoryInfo["MemoryAvail"] = color.Sprintf(formatString, memAvail, memAvailUnit)             // 可用内存

	return memoryInfo
}

// GetSwapInfo 获取交换分区信息
//
// 参数：
//   - dataUnit: 存储数据单位
//
// 返回：
//   - 交换分区信息
func GetSwapInfo(dataUnit string) map[string]interface{} {
	swapTotal, swapTotalUnit := Human(float64(memData.SwapTotal), "B")
	swapFree, swapFreeUnit := Human(float64(memData.SwapFree), "B")

	// 使用冒泡排序找出最大值用以组装格式字符串
	swapData := []float64{swapTotal, swapFree}
	BubbleSort(swapData)
	formatString := "%.2f %s"

	swapInfo := make(map[string]interface{})
	swapInfo["SwapStatus"] = func() string {
		if swapTotal == 0 {
			return "Unavailable"
		}
		return "Available"
	}()
	swapInfo["SwapTotal"] = color.Sprintf(formatString, swapTotal, swapTotalUnit) // 交换分区总量
	swapInfo["SwapFree"] = color.Sprintf(formatString, swapFree, swapFreeUnit)    // 交换分区空闲量

	return swapInfo
}

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
	cpuInfo["CPUModel"] = sysInfo.CPU.Model                            // CPU 型号
	cpuInfo["CPUNumber"] = sysInfo.CPU.Cpus                            // CPU 数量
	cpuInfo["CPUCores"] = sysInfo.CPU.Cores                            // CPU 核心数
	cpuInfo["CPUThreads"] = sysInfo.CPU.Threads                        // CPU 线程数
	cpuCache, cpuCacheUnit := Human(float64(sysInfo.CPU.Cache), "KiB") // CPU 缓存
	cpuInfo["CPUCache"] = color.Sprintf("%.0f %s", cpuCache, cpuCacheUnit)

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

	// 需要额外步骤获取的信息
	osCode := FindSystemCode(hostData.PlatformVersion) // 系统代号
	timeZone := GetTimeZoneOriginal()                  // 时区

	osInfo["OS"] = color.Sprintf("%s %s", osCode, hostData.PlatformVersion) // 操作系统
	osInfo["Arch"] = hostData.KernelArch                                    // 系统架构
	osInfo["Kernel"] = hostData.KernelVersion                               // 内核版本
	osInfo["Platform"] = UpperStringFirstChar(hostData.Platform)            // 平台
	osInfo["Hostname"] = hostData.Hostname                                  // 主机名
	osInfo["TimeZone"] = timeZone                                           // 时区

	return osInfo
}

// GetTimeZoneOriginal 获取时区信息的原始方法（检测 /etc/localtime 实际指向的文件）
//
// 返回：
//   - 时区信息
func GetTimeZoneOriginal() string {
	var localtimeFile = "/etc/localtime"
	var timeZone string

	filePath, err := ReadFileLink(localtimeFile)
	if err != nil {
		timeZone = ""
	}

	dir, file := path.Split(filePath)
	if len(dir) == 0 || len(file) == 0 {
		timeZone = ""
	}

	_, fname := path.Split(dir[:len(dir)-1])
	if fname == "zoneinfo" {
		timeZone = file
	} else {
		timeZone = path.Join(fname, file)
	}

	return timeZone
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
	timeInfo["BootTime"] = UnixTime2TimeString(hostData.BootTime) // 系统启动时间
	day, hour, minute, second := UnixTime2DayHourMinuteSecond(hostData.Uptime)
	result := color.Sprintf("%vd %vh %vm %vs", day, hour, minute, second)
	timeInfo["Uptime"] = result // 系统运行时间
	starttimeArgs := []string{"time"}
	StartTime, err := RunCommandGetResult("systemd-analyze", starttimeArgs)
	if err != nil {
		return nil, err
	}
	timeInfo["StartTime"] = strings.Split(strings.Split(StartTime, "\n")[0], "= ")[1] // 系统启动用时

	return timeInfo, nil
}

// GetPackageInfo 读取可更新包信息
//
// 参数：
//   - filePath: 更新信息记录文件路径
//   - line: 读取指定行，等于 0 时读取全部行
//
// 返回：
//   - 可更新包信息
//   - 错误信息
func GetPackageInfo(filePath string, line int) (map[string]interface{}, error) {
	var packageSlice []string
	if filePath != "" && FileExist(filePath) {
		// 打开文件
		text, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer text.Close()

		// 创建一个扫描器对象按行遍历
		scanner := bufio.NewScanner(text)
		// 行计数
		count := 1
		// 逐行读取，输出指定行
		for scanner.Scan() {
			if line == count {
				packageSlice = append(packageSlice, scanner.Text())
				break
			}
			packageSlice = append(packageSlice, scanner.Text())
			count++
		}
	}
	updateInfo := make(map[string]interface{})
	updateInfo["PackageList"] = packageSlice
	updateInfo["PackageQuantity"] = strconv.Itoa(len(packageSlice))

	return updateInfo, nil
}

// GetUpdateDaemonInfo 获取更新检测服务的信息
//
// 返回：
//   - 更新检测服务的信息
//   - 错误信息
func GetUpdateDaemonInfo() (map[string]interface{}, error) {
	daemonInfo := make(map[string]interface{})
	daemonArgs := []string{"is-active", "system-checkupdates.timer"}
	updateDaemonStatus, err := RunCommandGetResult("systemctl", daemonArgs)
	if err != nil {
		return nil, err
	}
	daemonInfo["UpdateDaemonStatus"] = updateDaemonStatus

	return daemonInfo, nil
}

// GetUserInfo 获取用户信息
//
// 返回：
//   - 用户信息
func GetUserInfo() map[string]interface{} {
	userInfo := make(map[string]interface{})
	userInfo["User"] = userData.Name           // 用户名称
	userInfo["UserName"] = userData.Username   // 用户昵称
	userInfo["UserUid"] = userData.Uid         // 用户 ID
	userInfo["UserGid"] = userData.Gid         // 用户组 ID
	userInfo["UserHomeDir"] = userData.HomeDir // 用户主目录

	return userInfo
}
