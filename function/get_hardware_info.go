/*
File: get_hardware_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-09-28 12:39:05

Description: 获取硬件信息
*/

package function

import (
	"encoding/json"
	"fmt"
)

// GetNicInfo 获取网卡信息
func GetNicInfo(address string) (nicInfo map[string]interface{}) {
	nicInfo = make(map[string]interface{})
	nicInfo["NicAddress"] = pciData.GetDevice(address).Address
	nicInfo["NicDriver"] = pciData.GetDevice(address).Driver
	nicInfo["NicVendor"] = pciData.GetDevice(address).Vendor.Name
	nicInfo["NicProduct"] = pciData.GetDevice(address).Product.Name

	return nicInfo
}

// GetStorageInfo 获取存储设备信息
func GetStorageInfo() (storageInfo map[string]interface{}) {
	storageInfo = make(map[string]interface{})
	for index, disk := range blockData.Disks {
		storageValue := make(map[string]interface{})
		if disk.SizeBytes > 0 {
			storageValue["StorageName"] = disk.Name
			storageValue["StorageDriver"] = disk.StorageController
			storageValue["StorageVendor"] = disk.Vendor
			storageValue["StorageModel"] = disk.Model
			storageValue["StorageType"] = disk.DriveType
			storageValue["StorageRemovable"] = disk.IsRemovable
			storageValue["StorageSerial"] = disk.SerialNumber
			storageSize, storageSizeUnit := DataUnitConvert("B", "TB", float64(disk.SizeBytes))
			storageValue["StorageSize"] = fmt.Sprintf("%.1f %s", storageSize, storageSizeUnit)
			storageInfo[fmt.Sprintf("%s%d", "Storage", index)] = storageValue
		}
	}

	return storageInfo
}

// GetGPUInfo 获取显卡信息
func GetGPUInfo() (gpuInfo map[string]interface{}) {
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
	err = json.Unmarshal([]byte(gpuDataJson), &gpuDataJ2S)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	gpuInfo = make(map[string]interface{})
	gpuInfo["GPUDriver"] = gpuDataJ2S.GPU.Cards[0].PCI.Driver
	gpuInfo["GPUAddress"] = gpuDataJ2S.GPU.Cards[0].PCI.Address
	gpuInfo["GPUVendor"] = gpuDataJ2S.GPU.Cards[0].PCI.Vendor.NAME
	gpuInfo["GPUProduct"] = gpuDataJ2S.GPU.Cards[0].PCI.Product.NAME

	return gpuInfo
}
