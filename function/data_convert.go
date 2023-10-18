/*
File: data_convert.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 12:45:37

Description: 数据转换（包括单位和格式）
*/

package function

import (
	"fmt"
	"strings"
	"time"
)

// 冒泡排序
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

// 动态计算浮点数长度并输出合适的格式字符串
func FormatFloat(value float64, precision int) string {
	digits := len(fmt.Sprint(int(value))) + 1 + precision // 整数部分长度 + 小数点 + 小数部分长度
	formatString := "%" + fmt.Sprintf("%d.1f", digits) + "%s"
	return formatString
}

// 秒转换为天、小时、分钟、秒
func Second2DayHourMinuteSecond(totalSeconds uint64) (day, hour, minute, second uint64) {
	day = totalSeconds / 86400
	hour = (totalSeconds - day*86400) / 3600
	minute = (totalSeconds - day*86400 - hour*3600) / 60
	second = totalSeconds - day*86400 - hour*3600 - minute*60
	return day, hour, minute, second
}

// uint64格式的时间戳转换为字符串格式
func Uint2TimeString(timeStamp uint64) string {
	return time.Unix(int64(timeStamp), 0).Format("2006-01-02 15:04:05")
}

// 最大化字符串的第一个字母
func UpperStringFirstChar(str string) string {
	if len(str) == 0 {
		return str
	}

	return strings.ToUpper(str[:1]) + str[1:]
}

// 数据单位转换
func DataUnitConvert(oldUnit string, newUnit string, data float64) (float64, string) {
	if oldUnit == "B" && newUnit == "KB" {
		if data < 1024 {
			newUnit = "B"
			return data, newUnit
		} else {
			return data / 1024, newUnit
		}
	} else if oldUnit == "B" && newUnit == "MB" {
		if data < 1024*1024 {
			newUnit = "KB"
			return data / 1024, newUnit
		} else {
			return data / 1024 / 1024, newUnit
		}
	} else if oldUnit == "B" && newUnit == "GB" {
		if data < 1024*1024*1024 {
			newUnit = "MB"
			return data / 1024 / 1024, newUnit
		} else {
			return data / 1024 / 1024 / 1024, newUnit
		}
	} else if oldUnit == "B" && newUnit == "TB" {
		if data < 1024*1024*1024*1024 {
			newUnit = "GB"
			return data / 1024 / 1024 / 1024, newUnit
		} else {
			return data / 1024 / 1024 / 1024 / 1024, newUnit
		}
	} else if oldUnit == "KB" && newUnit == "B" {
		return data * 1024, newUnit
	} else if oldUnit == "KB" && newUnit == "MB" {
		if data < 1024 {
			newUnit = "KB"
			return data, newUnit
		} else {
			return data / 1024, newUnit
		}
	} else if oldUnit == "KB" && newUnit == "GB" {
		if data < 1024*1024 {
			newUnit = "MB"
			return data / 1024, newUnit
		} else {
			return data / 1024 / 1024, newUnit
		}
	} else if oldUnit == "KB" && newUnit == "TB" {
		if data < 1024*1024*1024 {
			newUnit = "GB"
			return data / 1024 / 1024, newUnit
		} else {
			return data / 1024 / 1024 / 1024, newUnit
		}
	} else if oldUnit == "MB" && newUnit == "B" {
		return data * 1024 * 1024, newUnit
	} else if oldUnit == "MB" && newUnit == "KB" {
		return data * 1024, newUnit
	} else if oldUnit == "MB" && newUnit == "GB" {
		if data < 1024 {
			newUnit = "MB"
			return data, newUnit
		} else {
			return data / 1024, newUnit
		}
	} else if oldUnit == "MB" && newUnit == "TB" {
		if data < 1024*1024 {
			newUnit = "GB"
			return data / 1024, newUnit
		} else {
			return data / 1024 / 1024, newUnit
		}
	} else if oldUnit == "GB" && newUnit == "B" {
		return data * 1024 * 1024 * 1024, newUnit
	} else if oldUnit == "GB" && newUnit == "KB" {
		return data * 1024 * 1024, newUnit
	} else if oldUnit == "GB" && newUnit == "MB" {
		return data * 1024, newUnit
	} else if oldUnit == "GB" && newUnit == "TB" {
		if data < 1024 {
			newUnit = "GB"
			return data, newUnit
		} else {
			return data / 1024, newUnit
		}
	} else if oldUnit == "TB" && newUnit == "B" {
		return data * 1024 * 1024 * 1024 * 1024, newUnit
	} else if oldUnit == "TB" && newUnit == "KB" {
		return data * 1024 * 1024 * 1024, newUnit
	} else if oldUnit == "TB" && newUnit == "MB" {
		return data * 1024 * 1024, newUnit
	} else if oldUnit == "TB" && newUnit == "GB" {
		return data * 1024, newUnit
	} else {
		return data, oldUnit
	}
}
