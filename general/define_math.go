/*
File: define_math.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-03 14:20:52

Description: 处理数学计算
*/

package general

// max 返回两个数中的最大值
//
// 参数：
//   - x: 第一个数
//   - y: 第二个数
//
// 返回：
//   - 两个数中的最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// min 返回两个数中的最小值
//
// 参数：
//   - x: 第一个数
//   - y: 第二个数
//
// 返回：
//   - 两个数中的最小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
