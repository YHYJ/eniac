/*
File: data_convert.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 12:45:37

Description: 数据转换（包括单位和格式）
*/

package function

import "strings"

// 最大化字符串的第一个字母
func upperStringFirstChar(str string) string {
	if len(str) == 0 {
		return str
	}

	return strings.ToUpper(str[:1]) + str[1:]
}

// 数据单位转换
func dataUnitConvert(oldUnit string, newUnit string, data float64) (float64, string) {
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
		} // TODO
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
