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

	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/jaypipes/ghw"

	"github.com/gookit/color"
	"github.com/zcalusic/sysinfo"
)

var (
	pciData, _     = ghw.PCI()     // PCI 信息
	networkData, _ = ghw.Network() // 网络设备信息
	gpuData, _     = ghw.GPU()     // 显卡信息
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
			storageValue["StorageSize"] = color.Sprintf("%.1f %s", storageSize, storageSizeUnit)
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
		color.Printf("%s %s %s\n", DangerText(ErrorInfoFlag), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
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
		color.Printf("%s %s %s\n", DangerText(ErrorInfoFlag), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
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

// GetOSInfo 获取系统信息
//
// 参数：
//   - sysInfo: 总的系统信息 (System Info)
//
// 返回：
//   - 系统信息 (OS Info)
func GetOSInfo(sysInfo sysinfo.SysInfo) map[string]interface{} {
	osInfo := make(map[string]interface{})
	osInfo["OS"] = UpperFirstChar(sysInfo.OS.Name)         // 操作系统
	osInfo["Arch"] = sysInfo.OS.Architecture               // 系统架构
	osInfo["CurrentKernel"] = sysInfo.Kernel.Release       // 当前内核版本
	osInfo["LatestKernel"] = GetLatestKernelVersion()      // 本地最新内核版本
	osInfo["Platform"] = UpperFirstChar(sysInfo.OS.Vendor) // 平台
	osInfo["Hostname"] = hostData.Hostname                 // 主机名
	osInfo["TimeZone"] = sysInfo.Node.Timezone             // 时区

	return osInfo
}

// GetPackageInfo 获取安装包信息
//
// 返回：
//   - 安装包信息
//   - 错误信息
func GetPackageInfo() (map[string]interface{}, error) {
	packageInfo := make(map[string]interface{})
	packageData, err := GetInstalledPackageData()
	if err != nil {
		return nil, err
	}
	packageInfo["PackageAsExplicitCount"] = packageData.AsExplicitCount                                                    // 单独指定安装包总数
	packageInfo["PackageAsDependencyCount"] = packageData.AsDependencyCount                                                // 作为依赖安装包总数
	packageInfo["PackageTotalCount"] = packageData.PackageTotalCount                                                       // 已安装包总数
	packageInfo["PackageTotalSize"] = color.Sprintf("%.2f %s", packageData.PackageTotalSize, packageData.PackageTotalUnit) // 已安装包总大小

	return packageInfo, nil
}

// GetUpdatablePackageInfo 读取可更新包信息
//
// 参数：
//   - filePath: 更新信息记录文件路径
//   - line: 读取指定行，等于 0 时读取全部行
//
// 返回：
//   - 可更新包信息
//   - 错误信息
func GetUpdatablePackageInfo(filePath string, line int) (map[string]interface{}, error) {
	var (
		lastCheckTime string
		packageSlice  []string
	)
	if filePath != "" && FileExist(filePath) {
		// 获取文件最后修改时间作为最新更新检查时间
		lastCheckTime = GetFileModTime(filePath)

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
	updateInfo["LastCheckTime"] = lastCheckTime
	updateInfo["UpdatablePackageList"] = packageSlice
	updateInfo["UpdatablePackageQuantity"] = strconv.Itoa(len(packageSlice))

	return updateInfo, nil
}

// GetCheckUpdateDaemonInfo 获取更新检测服务的信息
//
// 返回：
//   - 更新检测服务的信息
//   - 错误信息
func GetCheckUpdateDaemonInfo() (map[string]interface{}, error) {
	daemonInfo := make(map[string]interface{})

	// 判断更新检测服务状态的依据
	const basis = "system-checkupdates.timer"

	// 检查更新检测服务是否可用（值为 enabled, disabled 或空字符串）
	daemonIsEnabledArgs := []string{"is-enabled", basis}
	updateDaemonIsEnabled, _, _ := RunCommandToBuffer("systemctl", daemonIsEnabledArgs)

	switch updateDaemonIsEnabled {
	case "enabled":
		// 检查更新检测服务是否处于活动状态
		daemonIsActiveArgs := []string{"is-active", basis}
		updateDaemonIsActive, _, _ := RunCommandToBuffer("systemctl", daemonIsActiveArgs)

		daemonInfo["UpdateDaemonStatus"] = UpperFirstChar(updateDaemonIsActive)
	case "disabled":
		daemonInfo["UpdateDaemonStatus"] = "disabled"
	default:
		daemonInfo["UpdateDaemonStatus"] = "not-found"
	}
	return daemonInfo, nil
}
