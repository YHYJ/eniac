/*
File: define_convert.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 12:45:37

Description: 单位和格式数据转换
*/

package general

import (
	"strings"
	"time"
)

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
func UnixTime2DayHourMinuteSecond(unixTime int64) (day, hour, minute, second int64) {
	day = unixTime / 86400
	hour = (unixTime - day*86400) / 3600
	minute = (unixTime - day*86400 - hour*3600) / 60
	second = unixTime - day*86400 - hour*3600 - minute*60
	return day, hour, minute, second
}

// UnixTime2TimeString Unix 时间戳转换为字符串格式
//
// 参数：
//   - timeStamp: Unix 时间戳
//
// 返回：
//   - 格式化的 Unix 时间戳字符串
func UnixTime2TimeString(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")
}

// UpperFirstChar 最大化字符串的第一个字母
//
// 参数：
//   - str: 需要处理的字符串
//
// 返回：
//   - 处理后的字符串
func UpperFirstChar(str string) string {
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
