/*
File: get_memory_swap_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 11:35:40

Description: 子命令 'get' 的实现，获取内存和交换分区信息
*/

package cli

import (
	"github.com/gookit/color"
	"github.com/yhyj/eniac/general"
)

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
	memTotal, memTotalUnit := general.Human(float64(memData.Total), "B")
	memUsed, memUsedUnit := general.Human(float64(memData.Used), "B")
	memUsedPercent, _ := general.Human(float64(memData.UsedPercent), percentUnit)
	memFree, memFreeUnit := general.Human(float64(memData.Free), "B")
	memShared, memSharedUnit := general.Human(float64(memData.Shared), "B")
	memBuffCache, memBuffCacheUnit := general.Human(float64(memData.Buffers+memData.Cached), "B")
	memAvail, memAvailUnit := general.Human(float64(memData.Available), "B")

	// 使用冒泡排序找出最大值用以组装格式字符串
	memData := []float64{memTotal, memUsed, memUsedPercent, memFree, memShared, memBuffCache, memAvail}
	general.BubbleSort(memData)
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
	swapTotal, swapTotalUnit := general.Human(float64(memData.SwapTotal), "B")
	swapFree, swapFreeUnit := general.Human(float64(memData.SwapFree), "B")

	// 使用冒泡排序找出最大值用以组装格式字符串
	swapData := []float64{swapTotal, swapFree}
	general.BubbleSort(swapData)
	formatString := "%.2f %s"

	swapInfo := make(map[string]interface{})
	swapInfo["SwapDisabled"] = false
	if swapTotal == 0 {
		swapInfo["SwapDisabled"] = true
	}
	swapInfo["SwapTotal"] = color.Sprintf(formatString, swapTotal, swapTotalUnit) // 交换分区总量
	swapInfo["SwapFree"] = color.Sprintf(formatString, swapFree, swapFreeUnit)    // 交换分区空闲量

	return swapInfo
}
