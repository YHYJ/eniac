/*
File: get_memswapinfo.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 11:35:40

Description: 获取内存和交换分区信息
*/

package function

import "github.com/shirou/gopsutil/mem"

var memInfo, _ = mem.VirtualMemory()

// GetMemoryInfo 获取内存信息
func GetMemoryInfo(dataUnit string, percentUnit string) (memoryInfo map[string]interface{}, err error) {
	memoryInfo = make(map[string]interface{})
	memoryInfo["MemTotal"], memoryInfo["MemTotalUnit"] = dataUnitConvert("B", dataUnit, float64(memInfo.Total))                          // 内存总量，内存总量单位
	memoryInfo["MemUsed"], memoryInfo["MemUsedUnit"] = dataUnitConvert("B", dataUnit, float64(memInfo.Used))                             // 已用内存，已用内存单位
	memoryInfo["MemUsedPercent"] = memInfo.UsedPercent                                                                                   // 内存使用率
	memoryInfo["MemUsedPercentUnit"] = percentUnit                                                                                       // 内存使用率单位
	memoryInfo["MemFree"], memoryInfo["MemFreeUnit"] = dataUnitConvert("B", dataUnit, float64(memInfo.Free))                             // 空闲内存，空闲内存单位
	memoryInfo["MemShared"], memoryInfo["MemSharedUnit"] = dataUnitConvert("B", dataUnit, float64(memInfo.Shared))                       // 共享内存，共享内存单位
	memoryInfo["MemBuffCache"], memoryInfo["MemBuffCacheUnit"] = dataUnitConvert("B", dataUnit, float64(memInfo.Buffers+memInfo.Cached)) // 缓存内存，缓存内存单位
	memoryInfo["MemAvail"], memoryInfo["MemAvailUnit"] = dataUnitConvert("B", dataUnit, float64(memInfo.Available))                      // 可用内存，可用内存单位

	return memoryInfo, err
}

// GetSwapInfo 获取交换分区信息
func GetSwapInfo(dataUnit string) (swapInfo map[string]interface{}, err error) {
	swapInfo = make(map[string]interface{})
	swapInfo["SwapTotal"], swapInfo["SwapTotalUnit"] = dataUnitConvert("B", dataUnit, float64(memInfo.SwapTotal)) // 交换分区总量，交换分区总量单位
	swapInfo["SwapFree"], swapInfo["SwapFreeUnit"] = dataUnitConvert("B", dataUnit, float64(memInfo.SwapFree))    // 交换分区空闲量，交换分区空闲量单位

	return swapInfo, err
}
