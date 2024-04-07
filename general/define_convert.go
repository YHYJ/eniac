/*
File: define_convert.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 12:45:37

Description: 单位和格式数据转换
*/

package general

import (
	"fmt"
	"strings"
	"time"
)

// BubbleSort 冒泡排序
//
// 参数：
//   - arr: 需要排序的切片
func BubbleSort(arr []float64) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换 arr[j] 和 arr[j+1] 的位置
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// FormatFloat 动态计算浮点数长度并输出合适的格式字符串
//
// 参数：
//   - value: 需要处理的浮点数
//   - precision: 期望的小数位数
//
// 返回：
//   - 格式化后的字符串
func FormatFloat(value float64, precision int) string {
	digits := len(fmt.Sprint(int(value))) + 1 + precision // 整数部分长度 + 小数点 + 小数部分长度
	formatString := "%" + fmt.Sprintf("%d.1f ", digits) + "%s"
	return formatString
}

// UnixTime2DayHourMinuteSecond Unix 时间戳转换为天、小时、分钟、秒
//
// 参数：
//   - totalSeconds: Unix 时间戳
//
// 返回：
//   - day: 天
//   - hour: 小时
//   - minute: 分钟
//   - second: 秒
func UnixTime2DayHourMinuteSecond(unixTime uint64) (day, hour, minute, second uint64) {
	day = unixTime / 86400
	hour = (unixTime - day*86400) / 3600
	minute = (unixTime - day*86400 - hour*3600) / 60
	second = unixTime - day*86400 - hour*3600 - minute*60
	return day, hour, minute, second
}

// UnixTime2TimeString uint64 格式的 Unix 时间戳转换为字符串格式
//
// 参数：
//   - timeStamp: Unix 时间戳
//
// 返回：
//   - 格式化的 Unix 时间戳字符串
func UnixTime2TimeString(unixTime uint64) string {
	return time.Unix(int64(unixTime), 0).Format("2006-01-02 15:04:05")
}

// UpperStringFirstChar 最大化字符串的第一个字母
//
// 参数：
//   - str: 需要处理的字符串
//
// 返回：
//   - 处理后的字符串
func UpperStringFirstChar(str string) string {
	if len(str) == 0 {
		return str
	}

	return strings.ToUpper(str[:1]) + str[1:]
}

// Human 存储数据转换为人类可读的格式
//
// 参数：
//   - size: 需要转换的存储数据
//   - initialUnit: 初始单位
//
// 返回：
//   - 转换后的数据
//   - 转换后的单位
func Human(size float64, initialUnit string) (float64, string) {
	allUnits := [...]string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}

	// 重构单位切片，从 initialUnit 开始
	var units []string
	for i, unit := range allUnits {
		if unit == initialUnit {
			units = allUnits[i:]
			break
		}
	}

	// 数据及转换
	for _, unit := range units {
		if size < 1024 {
			return size, unit
		}
		size /= 1024
	}

	return size, initialUnit
}
