/*
File: define_color.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-05-29 15:46:55

Description: 定义输出颜色
*/

package general

import (
	"math/rand"

	"github.com/gookit/color"
	"github.com/olekukonko/tablewriter"
)

var (
	FgBlackText        = color.FgBlack.Render        // 前景色 - 黑色
	FgWhiteText        = color.FgWhite.Render        // 前景色 - 白色
	FgLightWhiteText   = color.FgLightWhite.Render   // 前景色 - 亮白色
	FgGrayText         = color.FgGray.Render         // 前景色 - 灰色
	FgRedText          = color.FgRed.Render          // 前景色 - 红色
	FgLightRedText     = color.FgLightRed.Render     // 前景色 - 亮红色
	FgGreenText        = color.FgGreen.Render        // 前景色 - 绿色
	FgLightGreenText   = color.FgLightGreen.Render   // 前景色 - 亮绿色
	FgYellowText       = color.FgYellow.Render       // 前景色 - 黄色
	FgLightYellowText  = color.FgLightYellow.Render  // 前景色 - 亮黄色
	FgBlueText         = color.FgBlue.Render         // 前景色 - 蓝色
	FgLightBlueText    = color.FgLightBlue.Render    // 前景色 - 亮蓝色
	FgMagentaText      = color.FgMagenta.Render      // 前景色 - 品红
	FgLightMagentaText = color.FgLightMagenta.Render // 前景色 - 亮品红
	FgCyanText         = color.FgCyan.Render         // 前景色 - 青色
	FgLightCyanText    = color.FgLightCyan.Render    // 前景色 - 亮青色

	BgBlackText        = color.BgBlack.Render        // 背景色 - 黑色
	BgWhiteText        = color.BgWhite.Render        // 背景色 - 白色
	BgLightWhiteText   = color.BgLightWhite.Render   // 背景色 - 亮白色
	BgGrayText         = color.BgGray.Render         // 背景色 - 灰色
	BgRedText          = color.BgRed.Render          // 背景色 - 红色
	BgLightRedText     = color.BgLightRed.Render     // 背景色 - 亮红色
	BgGreenText        = color.BgGreen.Render        // 背景色 - 绿色
	BgLightGreenText   = color.BgLightGreen.Render   // 背景色 - 亮绿色
	BgYellowText       = color.BgYellow.Render       // 背景色 - 黄色
	BgLightYellowText  = color.BgLightYellow.Render  // 背景色 - 亮黄色
	BgBlueText         = color.BgBlue.Render         // 背景色 - 蓝色
	BgLightBlueText    = color.BgLightBlue.Render    // 背景色 - 亮蓝色
	BgMagentaText      = color.BgMagenta.Render      // 背景色 - 品红
	BgLightMagentaText = color.BgLightMagenta.Render // 背景色 - 亮品红
	BgCyanText         = color.BgCyan.Render         // 背景色 - 青色
	BgLightCyanText    = color.BgLightCyan.Render    // 背景色 - 亮青色

	InfoText      = color.Info.Render      // Info 文本
	NoteText      = color.Note.Render      // Note 文本
	LightText     = color.Light.Render     // Light 文本
	ErrorText     = color.Error.Render     // Error 文本
	DangerText    = color.Danger.Render    // Danger 文本
	NoticeText    = color.Notice.Render    // Notice 文本
	SuccessText   = color.Success.Render   // Success 文本
	CommentText   = color.Comment.Render   // Comment 文本
	PrimaryText   = color.Primary.Render   // Primary 文本
	WarnText      = color.Warn.Render      // Warn 文本
	QuestionText  = color.Question.Render  // Question 文本
	SecondaryText = color.Secondary.Render // Secondary 文本
)

// 可选的颜色列表 - Table 专用
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

// 上一次随机颜色的索引 - Table 专用
var previousColor int

// GetColor 随机获取一个颜色 - Table 专用
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
