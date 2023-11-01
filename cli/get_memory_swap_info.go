/*
File: get_memory_swap_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 11:35:40

Description: 获取内存和交换分区信息
*/

package cli

import (
	"fmt"

	"github.com/yhyj/eniac/general"
)

// GetMemoryInfo 获取内存信息
func GetMemoryInfo(dataUnit string, percentUnit string) map[string]interface{} {
	// 内存数据
	memTotal, memTotalUnit := general.DataUnitConvert("B", dataUnit, float64(memData.Total))
	memUsed, memUsedUnit := general.DataUnitConvert("B", dataUnit, float64(memData.Used))
	memUsedPercent, _ := general.DataUnitConvert("B", percentUnit, float64(memData.UsedPercent))
	memFree, memFreeUnit := general.DataUnitConvert("B", dataUnit, float64(memData.Free))
	memShared, memSharedUnit := general.DataUnitConvert("B", dataUnit, float64(memData.Shared))
	memBuffCache, memBuffCacheUnit := general.DataUnitConvert("B", dataUnit, float64(memData.Buffers+memData.Cached))
	memAvail, memAvailUnit := general.DataUnitConvert("B", dataUnit, float64(memData.Available))

	// 使用冒泡排序找出最大值用以组装格式字符串
	memData := []float64{memTotal, memUsed, memUsedPercent, memFree, memShared, memBuffCache, memAvail}
	general.BubbleSort(memData)
	max := memData[len(memData)-1]
	formatString := general.FormatFloat(max, 1)

	memoryInfo := make(map[string]interface{})
	memoryInfo["MemoryTotal"] = fmt.Sprintf(formatString, memTotal, memTotalUnit)             // 内存总量
	memoryInfo["MemoryUsed"] = fmt.Sprintf(formatString, memUsed, memUsedUnit)                // 已用内存
	memoryInfo["MemoryUsedPercent"] = fmt.Sprintf(formatString, memUsedPercent, percentUnit)  // 内存使用率
	memoryInfo["MemoryFree"] = fmt.Sprintf(formatString, memFree, memFreeUnit)                // 空闲内存
	memoryInfo["MemoryShared"] = fmt.Sprintf(formatString, memShared, memSharedUnit)          // 共享内存
	memoryInfo["MemoryBuffCache"] = fmt.Sprintf(formatString, memBuffCache, memBuffCacheUnit) // 缓存内存
	memoryInfo["MemoryAvail"] = fmt.Sprintf(formatString, memAvail, memAvailUnit)             // 可用内存

	return memoryInfo
}

// GetSwapInfo 获取交换分区信息
func GetSwapInfo(dataUnit string) map[string]interface{} {
	swapTotal, swapTotalUnit := general.DataUnitConvert("B", dataUnit, float64(memData.SwapTotal))
	swapFree, swapFreeUnit := general.DataUnitConvert("B", dataUnit, float64(memData.SwapFree))

	// 使用冒泡排序找出最大值用以组装格式字符串
	swapData := []float64{swapTotal, swapFree}
	general.BubbleSort(swapData)
	max := swapData[len(swapData)-1]
	formatString := general.FormatFloat(max, 1)

	swapInfo := make(map[string]interface{})
	swapInfo["SwapDisabled"] = false
	if swapTotal == 0 {
		swapInfo["SwapDisabled"] = true
	}
	swapInfo["SwapTotal"] = fmt.Sprintf(formatString, swapTotal, swapTotalUnit) // 交换分区总量
	swapInfo["SwapFree"] = fmt.Sprintf(formatString, swapFree, swapFreeUnit)    // 交换分区空闲量

	return swapInfo
}
