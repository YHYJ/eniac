//go:build linux

/*
File: define_spider_linux.go
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

	"encoding/json"
	"regexp"
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
	pciData, _     = ghw.PCI()           // PCI 信息
	blockData, _   = ghw.Block()         // 存储设备信息
	networkData, _ = ghw.Network()       // 网络设备信息
	gpuData, _     = ghw.GPU()           // 显卡信息
	loadData, _    = load.Avg()          // 系统负载信息
	memData, _     = mem.VirtualMemory() // 内存信息
	hostData, _    = host.Info()         // 主机信息
	userData, _    = user.Current()      // 用户信息
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
			storageValue["StorageVendor"] = func() string {
				if disk.Vendor == "unknown" {
					// 检测是否符合 PCI 地址格式
					pciPattern := "^[0-9A-Fa-f]{4}:[0-9A-Fa-f]{2}:[0-9A-Fa-f]{2}\\.[0-9A-Fa-f]$"
					diskPciAddress := func() string {
						if len(strings.Split(disk.BusPath, "-")) < 2 {
							return ""
						}
						return strings.Split(disk.BusPath, "-")[1]
					}()
					matched, err := regexp.MatchString(pciPattern, diskPciAddress)
					if err != nil {
						return "--/--"
					}
					if matched {
						return pciData.GetDevice(diskPciAddress).Vendor.Name
					} else {
						return "--/--"
					}
				}
				return disk.Vendor
			}()
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

// GetGPUInfo 获取显卡信息
//
// 返回：
//   - 显卡信息
func GetGPUInfo() map[string]interface{} {
	type GPUDataJ2S struct {
		GPU struct {
			Cards []struct {
				Address string `json:"address"`
				Index   int    `json:"index"`
				PCI     struct {
					Driver   string `json:"driver"`
					Address  string `json:"address"`
					Revision string `json:"revision"`
					Vendor   struct {
						ID   string `json:"id"`
						NAME string `json:"name"`
					} `json:"vendor"`
					Product struct {
						ID   string `json:"id"`
						NAME string `json:"name"`
					} `json:"product"`
					Subsystem struct {
						ID   string `json:"id"`
						NAME string `json:"name"`
					} `json:"subsystem"`
					Class struct {
						ID   string `json:"id"`
						NAME string `json:"name"`
					} `json:"class"`
					Subclass struct {
						ID   string `json:"id"`
						NAME string `json:"name"`
					} `json:"subclass"`
					ProgrammingInterface struct {
						ID   string `json:"id"`
						NAME string `json:"name"`
					} `json:"programming_interface"`
				} `json:"pci"`
			} `json:"cards"`
		} `json:"gpu"`
	}

	// 获取 JSON 类型的显卡信息
	gpuDataJson := gpuData.JSONString(false)

	// 解析 JSON
	var gpuDataJ2S GPUDataJ2S
	if err := json.Unmarshal([]byte(gpuDataJson), &gpuDataJ2S); err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Danger.Printf("Parse JSON error (%s:%d): %s\n", fileName, lineNo+1, err)
	}

	gpuInfo := make(map[string]interface{})
	gpuInfo["GPUDriver"] = gpuDataJ2S.GPU.Cards[0].PCI.Driver
	gpuInfo["GPUAddress"] = gpuDataJ2S.GPU.Cards[0].PCI.Address
	gpuInfo["GPUVendor"] = gpuDataJ2S.GPU.Cards[0].PCI.Vendor.NAME
	gpuInfo["GPUProduct"] = gpuDataJ2S.GPU.Cards[0].PCI.Product.NAME

	return gpuInfo
}

// GetNicInfo 获取网卡信息
//
// 返回：
//   - 网卡信息
func GetNicInfo() map[string]interface{} {
	type NICDataJ2S struct {
		Name       string `json:"name"`
		MacAddress string `json:"mac_address"`
		IsVirtual  bool   `json:"is_virtual"`
		PCIAddress string `json:"pci_address"`
		Speed      string `json:"speed"`
		Duplex     string `json:"duplex"`
	}
	type NetworkDataJ2S struct {
		Nics []NICDataJ2S `json:"nics"`
	}

	// 获取 JSON 类型的网络信息
	networkDataJson := networkData.JSONString(false)

	// 解析 JSON
	var networkDataJ2S map[string]NetworkDataJ2S
	if err := json.Unmarshal([]byte(networkDataJson), &networkDataJ2S); err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Danger.Printf("Parse JSON error (%s:%d): %s\n", fileName, lineNo+1, err)
	}

	// 访问解析后的数据
	networkInfo := make(map[string]interface{})
	network := networkDataJ2S["network"]
	index := 1 // 排除编号为0的虚拟网卡
	for _, nic := range network.Nics {
		networkValue := make(map[string]interface{})
		if !nic.IsVirtual {
			networkValue["NicName"] = nic.Name
			if nic.PCIAddress != "" {
				networkValue["NicPCIAddress"] = nic.PCIAddress
				networkValue["NicDriver"] = pciData.GetDevice(nic.PCIAddress).Driver
				networkValue["NicProduct"] = pciData.GetDevice(nic.PCIAddress).Product.Name
				networkValue["NicVendor"] = pciData.GetDevice(nic.PCIAddress).Vendor.Name
			} else {
				networkValue["NicPCIAddress"] = "--/--"
				networkValue["NicDriver"] = "--/--"
				networkValue["NicProduct"] = "--/--"
				networkValue["NicVendor"] = "--/--"
			}
			networkValue["NicMacAddress"] = func() string {
				if nic.MacAddress == "" {
					return "--/--"
				}
				return nic.MacAddress
			}()
			networkValue["NicSpeed"] = func() string {
				if nic.Speed == "" {
					return "--/--"
				}
				return nic.Speed
			}()
			networkValue["NicDuplex"] = func() string {
				if nic.Duplex == "" {
					return "--/--"
				}
				return nic.Duplex
			}()
			networkInfo[color.Sprintf("%d", index)] = networkValue
			index += 1
		}
	}

	return networkInfo
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
	osInfo["OS"] = UpperStringFirstChar(sysInfo.OS.Name)         // 操作系统
	osInfo["Arch"] = sysInfo.OS.Architecture                     // 系统架构
	osInfo["CurrentKernel"] = sysInfo.Kernel.Release             // 当前内核版本
	osInfo["LatestKernel"] = GetLatestKernelVersion()            // 本地最新内核版本
	osInfo["Platform"] = UpperStringFirstChar(sysInfo.OS.Vendor) // 平台
	osInfo["Hostname"] = hostData.Hostname                       // 主机名
	osInfo["TimeZone"] = sysInfo.Node.Timezone                   // 时区

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
