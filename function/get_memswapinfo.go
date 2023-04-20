/*
File: get_memswapinfo.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 11:35:40

Description: 获取内存和交换分区信息
*/

package function

import "github.com/shirou/gopsutil/mem"

// MemoryInfoStruct 内存信息结构体
type MemoryInfoStruct struct {
	MemTotal           float64 `json:"mem_total"`             // 内存总量
	MemTotalUnit       string  `json:"mem_total_unit"`        // 内存总量单位
	MemUsed            float64 `json:"mem_used"`              // 已用内存
	MemUsedUnit        string  `json:"mem_used_unit"`         // 已用内存单位
	MemUsedPercent     float64 `json:"mem_used_percent"`      // 内存使用率
	MemUsedPercentUnit string  `json:"mem_used_percent_unit"` // 内存使用率单位
	MemFree            float64 `json:"mem_free"`              // 空闲内存
	MemFreeUnit        string  `json:"mem_free_unit"`         // 空闲内存单位
	MemShared          float64 `json:"mem_shared"`            // 共享内存
	MemSharedUnit      string  `json:"mem_shared_unit"`       // 共享内存单位
	MemBuffCache       float64 `json:"mem_buff_cache"`        // 缓存内存
	MemBuffCacheUnit   string  `json:"mem_buff_cache_unit"`   // 缓存内存单位
	MemAvail           float64 `json:"mem_avail"`             // 可用内存
	MemAvailUnit       string  `json:"mem_avail_unit"`        // 可用内存单位
}

// SwapInfoStruct 交换分区信息结构体
type SwapInfoStruct struct {
	SwapTotal     float64 `json:"swap_total"`      // 交换区总量
	SwapTotalUnit string  `json:"swap_total_unit"` // 交换区总量单位
	SwapFree      float64 `json:"swap_free"`       // 空闲交换区
	SwapFreeUnit  string  `json:"swap_free_unit"`  // 空闲交换区单位
}

var memInfo, _ = mem.VirtualMemory()

// GetMemoryInfo 获取内存信息
func GetMemoryInfo(dataUnit string, percentUnit string) (memoryInfo MemoryInfoStruct, err error) {
	memoryInfo.MemTotal, memoryInfo.MemTotalUnit = dataUnitConvert("B", dataUnit, float64(memInfo.Total))
	memoryInfo.MemUsed, memoryInfo.MemUsedUnit = dataUnitConvert("B", dataUnit, float64(memInfo.Used))
	memoryInfo.MemUsedPercent = memInfo.UsedPercent
	memoryInfo.MemUsedPercentUnit = percentUnit
	memoryInfo.MemFree, memoryInfo.MemFreeUnit = dataUnitConvert("B", dataUnit, float64(memInfo.Free))
	memoryInfo.MemShared, memoryInfo.MemSharedUnit = dataUnitConvert("B", dataUnit, float64(memInfo.Shared))
	memoryInfo.MemBuffCache, memoryInfo.MemBuffCacheUnit = dataUnitConvert("B", dataUnit, float64(memInfo.Buffers+memInfo.Cached))
	memoryInfo.MemAvail, memoryInfo.MemAvailUnit = dataUnitConvert("B", dataUnit, float64(memInfo.Available))
	return memoryInfo, err
}

// GetSwapInfo 获取交换分区信息
func GetSwapInfo(dataUnit string) (swapInfo SwapInfoStruct, err error) {
	swapInfo.SwapTotal, swapInfo.SwapTotalUnit = dataUnitConvert("B", dataUnit, float64(memInfo.SwapTotal))
	swapInfo.SwapFree, swapInfo.SwapFreeUnit = dataUnitConvert("B", dataUnit, float64(memInfo.SwapFree))
	return swapInfo, err
}
