/*
File: get_memory_swap_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 11:35:40

Description: 获取内存和交换分区信息
*/

package function

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

var memInfo, _ = mem.VirtualMemory()

// GetMemoryInfo 获取内存信息
func GetMemoryInfo(dataUnit string, percentUnit string) (memoryInfo map[string]interface{}, err error) {
	memoryInfo = make(map[string]interface{})
	memTotal, memTotalUnit := DataUnitConvert("B", dataUnit, float64(memInfo.Total))
	memoryInfo["MemTotal"] = fmt.Sprintf("%.2f %s", memTotal, memTotalUnit) // 内存总量
	memUsed, memUsedUnit := DataUnitConvert("B", dataUnit, float64(memInfo.Used))
	memoryInfo["MemUsed"] = fmt.Sprintf("%.2f %s", memUsed, memUsedUnit) // 已用内存
	memUsedPercent, _ := DataUnitConvert("B", percentUnit, float64(memInfo.UsedPercent))
	memoryInfo["MemUsedPercent"] = fmt.Sprintf("%.2f %s", memUsedPercent, percentUnit) // 内存使用率
	memFree, memFreeUnit := DataUnitConvert("B", dataUnit, float64(memInfo.Free))
	memoryInfo["MemFree"] = fmt.Sprintf("%.2f %s", memFree, memFreeUnit) // 空闲内存
	memShared, memSharedUnit := DataUnitConvert("B", dataUnit, float64(memInfo.Shared))
	memoryInfo["MemShared"] = fmt.Sprintf("%.2f %s", memShared, memSharedUnit) // 共享内存
	memBuffCache, memBuffCacheUnit := DataUnitConvert("B", dataUnit, float64(memInfo.Buffers+memInfo.Cached))
	memoryInfo["MemBuffCache"] = fmt.Sprintf("%.2f %s", memBuffCache, memBuffCacheUnit) // 缓存内存
	memAvail, memAvailUnit := DataUnitConvert("B", dataUnit, float64(memInfo.Available))
	memoryInfo["MemAvail"] = fmt.Sprintf("%.2f %s", memAvail, memAvailUnit) // 可用内存

	return memoryInfo, err
}

// GetSwapInfo 获取交换分区信息
func GetSwapInfo(dataUnit string) (swapInfo map[string]interface{}, err error) {
	swapInfo = make(map[string]interface{})
	swapTotal, swapTotalUnit := DataUnitConvert("B", dataUnit, float64(memInfo.SwapTotal))
	swapInfo["SwapTotal"] = fmt.Sprintf("%.2f %s", swapTotal, swapTotalUnit) // 交换分区总量
	swapFree, swapFreeUnit := DataUnitConvert("B", dataUnit, float64(memInfo.SwapFree))
	swapInfo["SwapFree"] = fmt.Sprintf("%.2f %s", swapFree, swapFreeUnit) // 交换分区空闲量

	return swapInfo, err
}
