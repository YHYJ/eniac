/*
File: define_color.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-20 10:22:15

Description:
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
	tablewriter.FgRedColor,
	tablewriter.FgHiRedColor,
	// 白
	tablewriter.FgWhiteColor,
	tablewriter.FgHiWhiteColor,
	// 黄
	tablewriter.FgYellowColor,
	tablewriter.FgHiYellowColor,
}

// 获取一个随机颜色并从可选列表中移除
func GetColor() int {
	if len(availableColors) == 0 {
		return tablewriter.FgWhiteColor
	}

	index := rand.Intn(len(availableColors))
	color := availableColors[index]

	// 从可选列表中移除已选中的颜色
	availableColors = append(availableColors[:index], availableColors[index+1:]...)

	return color
}
