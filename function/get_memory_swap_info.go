/*
File: get_memory_swap_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 11:35:40

Description: 获取内存和交换分区信息
*/

package function

import "fmt"

// GetMemoryInfo 获取内存信息
func GetMemoryInfo(dataUnit string, percentUnit string) map[string]interface{} {
	memoryInfo := make(map[string]interface{})
	memTotal, memTotalUnit := DataUnitConvert("B", dataUnit, float64(memData.Total))
	memoryInfo["MemoryTotal"] = fmt.Sprintf("%6.1f%s", memTotal, memTotalUnit) // 内存总量
	memUsed, memUsedUnit := DataUnitConvert("B", dataUnit, float64(memData.Used))
	memoryInfo["MemoryUsed"] = fmt.Sprintf("%6.1f%s", memUsed, memUsedUnit) // 已用内存
	memUsedPercent, _ := DataUnitConvert("B", percentUnit, float64(memData.UsedPercent))
	memoryInfo["MemoryUsedPercent"] = fmt.Sprintf("%6.1f%s", memUsedPercent, percentUnit) // 内存使用率
	memFree, memFreeUnit := DataUnitConvert("B", dataUnit, float64(memData.Free))
	memoryInfo["MemoryFree"] = fmt.Sprintf("%6.1f%s", memFree, memFreeUnit) // 空闲内存
	memShared, memSharedUnit := DataUnitConvert("B", dataUnit, float64(memData.Shared))
	memoryInfo["MemoryShared"] = fmt.Sprintf("%6.1f%s", memShared, memSharedUnit) // 共享内存
	memBuffCache, memBuffCacheUnit := DataUnitConvert("B", dataUnit, float64(memData.Buffers+memData.Cached))
	memoryInfo["MemoryBuffCache"] = fmt.Sprintf("%6.1f%s", memBuffCache, memBuffCacheUnit) // 缓存内存
	memAvail, memAvailUnit := DataUnitConvert("B", dataUnit, float64(memData.Available))
	memoryInfo["MemoryAvail"] = fmt.Sprintf("%6.1f%s", memAvail, memAvailUnit) // 可用内存

	return memoryInfo
}

// GetSwapInfo 获取交换分区信息
func GetSwapInfo(dataUnit string) map[string]interface{} {
	swapInfo := make(map[string]interface{})
	swapTotal, swapTotalUnit := DataUnitConvert("B", dataUnit, float64(memData.SwapTotal))
	swapInfo["SwapDisabled"] = false
	if swapTotal == 0 {
		swapInfo["SwapDisabled"] = true
	}
	swapInfo["SwapTotal"] = fmt.Sprintf("%6.1f%s", swapTotal, swapTotalUnit) // 交换分区总量
	swapFree, swapFreeUnit := DataUnitConvert("B", dataUnit, float64(memData.SwapFree))
	swapInfo["SwapFree"] = fmt.Sprintf("%6.1f%s", swapFree, swapFreeUnit) // 交换分区空闲量

	return swapInfo
}
