/*
File: define_math.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-03 14:20:52

Description: 处理数学计算
*/

package general

// Max 返回两个数中的最大值
//
// 参数：
//   - x: 第一个数
//   - y: 第二个数
//
// 返回：
//   - 两个数中的最大值
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min 返回两个数中的最小值
//
// 参数：
//   - x: 第一个数
//   - y: 第二个数
//
// 返回：
//   - 两个数中的最小值
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MapBoolCounter 统计指定布尔值在映射中出现的次数
//
// 参数：
//   - elements: 目标映射
//   - target: 目标元素
//
// 返回：
//   - 出现次数
func MapBoolCounter(elements map[string]bool, target bool) int {
	count := 0
	for _, element := range elements {
		if element == target {
			count++
		}
	}
	return count
}
