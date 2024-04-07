/*
File: define_color.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-20 10:22:15

Description: 关于输出的 Table 的颜色
*/

package general

import (
	"math/rand"

	"github.com/olekukonko/tablewriter"
)

// 可选的颜色列表
var availableColors = []int{
	// 蓝
	tablewriter.FgBlueColor,
	tablewriter.FgHiBlueColor,
	// 青
	tablewriter.FgCyanColor,
	tablewriter.FgHiCyanColor,
	// 绿
	tablewriter.FgGreenColor,
	tablewriter.FgHiGreenColor,
	// 品红
	tablewriter.FgMagentaColor,
	tablewriter.FgHiMagentaColor,
	// 红
	// tablewriter.FgRedColor,
	tablewriter.FgHiRedColor,
	// 白
	tablewriter.FgWhiteColor,
	tablewriter.FgHiWhiteColor,
	// 黄
	tablewriter.FgYellowColor,
	tablewriter.FgHiYellowColor,
}

// 上一次随机颜色的索引
var previousColor int

// GetColor 随机获取一个颜色
//
// 返回：
//   - 颜色代码
func GetColor() int {
	if len(availableColors) == 0 {
		return tablewriter.FgWhiteColor
	}

	// 随机取一个颜色
	index := rand.Intn(len(availableColors) - 1)
	// 如果 index == previousColor，则在 availableColors 长度范围内使 index + 1，超出长度范围则使 index - 1
	if index == previousColor {
		if index <= len(availableColors)-1 {
			index++
		} else {
			index--
		}
	} else {
		previousColor = index
	}
	color := availableColors[index]

	return color
}
