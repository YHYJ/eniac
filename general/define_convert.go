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
)

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
