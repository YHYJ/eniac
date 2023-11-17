/*
File: get_hardware_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-09-28 12:39:05

Description: 获取硬件信息
*/

package cli

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/yhyj/eniac/general"
)

// GetStorageInfo 获取存储设备信息
func GetStorageInfo() map[string]interface{} {
	storageInfo := make(map[string]interface{})
	index := 1 // 排除虚拟设备影响的编号
	for _, disk := range blockData.Disks {
		storageValue := make(map[string]interface{})
		if disk.SizeBytes > 0 && disk.DriveType.String() != "virtual" {
			storageValue["StorageName"] = disk.Name
			storageValue["StorageDriver"] = disk.StorageController.String()
			storageValue["StorageVendor"] = func() string {
				if disk.Vendor == "unknown" {
					// 检测是否符合PCI地址格式
					pciPattern := "^[0-9A-Fa-f]{4}:[0-9A-Fa-f]{2}:[0-9A-Fa-f]{2}\\.[0-9A-Fa-f]$"
					diskPciAddress := func() string {
						if len(strings.Split(disk.BusPath, "-")) < 2 {
							return ""
						}
						return strings.Split(disk.BusPath, "-")[1]
					}()
					matched, err := regexp.MatchString(pciPattern, diskPciAddress)
					if err != nil {
						return "<unknown>"
					}
					if matched {
						return pciData.GetDevice(diskPciAddress).Vendor.Name
					} else {
						return "<unknown>"
					}
				}
				return disk.Vendor
			}()
			storageValue["StorageModel"] = disk.Model
			storageValue["StorageType"] = disk.DriveType.String()
			storageValue["StorageRemovable"] = strconv.FormatBool(disk.IsRemovable)
			storageValue["StorageSerial"] = disk.SerialNumber
			storageSize, storageSizeUnit := general.Human(float64(disk.SizeBytes), "B")
			storageValue["StorageSize"] = fmt.Sprintf("%.2f %s", storageSize, storageSizeUnit)
			storageInfo[fmt.Sprintf("%d", index)] = storageValue
			index += 1
		}
	}

	return storageInfo
}

// GetGPUInfo 获取显卡信息
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

	// 获取JSON类型的显卡信息
	gpuDataJson := gpuData.JSONString(false)

	// 解析JSON
	var gpuDataJ2S GPUDataJ2S
	if err := json.Unmarshal([]byte(gpuDataJson), &gpuDataJ2S); err != nil {
		fmt.Println("Error:", err)
	}

	gpuInfo := make(map[string]interface{})
	gpuInfo["GPUDriver"] = gpuDataJ2S.GPU.Cards[0].PCI.Driver
	gpuInfo["GPUAddress"] = gpuDataJ2S.GPU.Cards[0].PCI.Address
	gpuInfo["GPUVendor"] = gpuDataJ2S.GPU.Cards[0].PCI.Vendor.NAME
	gpuInfo["GPUProduct"] = gpuDataJ2S.GPU.Cards[0].PCI.Product.NAME

	return gpuInfo
}

// GetNicInfo 获取网卡信息
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

	// 获取JSON类型的网络信息
	networkDataJson := networkData.JSONString(false)

	// 解析JSON
	var networkDataJ2S map[string]NetworkDataJ2S
	if err := json.Unmarshal([]byte(networkDataJson), &networkDataJ2S); err != nil {
		fmt.Println("Error:", err)
	}

	// 访问解析后的数据
	networkInfo := make(map[string]interface{})
	network := networkDataJ2S["network"]
	index := 1 // 排除虚拟网卡影响的编号
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
				networkValue["NicPCIAddress"] = "<unknown>"
				networkValue["NicDriver"] = "<unknown>"
				networkValue["NicProduct"] = "<unknown>"
				networkValue["NicVendor"] = "<unknown>"
			}
			networkValue["NicMacAddress"] = func() string {
				if nic.MacAddress == "" {
					return "<unknown>"
				}
				return nic.MacAddress
			}()
			networkValue["NicSpeed"] = func() string {
				if nic.Speed == "" {
					return "<unknown>"
				}
				return nic.Speed
			}()
			networkValue["NicDuplex"] = func() string {
				if nic.Duplex == "" {
					return "<unknown>"
				}
				return nic.Duplex
			}()
			networkInfo[fmt.Sprintf("%d", index)] = networkValue
			index += 1
		}
	}

	return networkInfo
}
