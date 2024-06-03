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

	"github.com/charmbracelet/lipgloss"
	"github.com/gookit/color"
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

// Tab 专用
const (
	TabLightColor = "#874BFD"
	TabDarkColor  = "#7D56F4"
)

//  Table 专用

// Notice: NEW
const (
	HeaderColor    = lipgloss.Color("#CCCCCC") // 表头颜色
	BorderColor    = lipgloss.Color("#6C757D") // 边框颜色
	ColumnOneColor = lipgloss.Color("#555555") // 第一列颜色

	DefaultColor = lipgloss.Color("#FFFFFF") // 默认颜色
)

var availableColors = []lipgloss.Color{
	// 红
	lipgloss.Color("#DC143C"),
	lipgloss.Color("#E74C3C"),
	// 橙
	lipgloss.Color("#FFA500"),
	lipgloss.Color("#FF8C00"),
	lipgloss.Color("#F39C76"),
	// 黄
	lipgloss.Color("#F5B041"),
	lipgloss.Color("#FFA500"),
	lipgloss.Color("#F9E79F"),
	// 绿
	lipgloss.Color("#008000"),
	lipgloss.Color("#70AD47"),
	lipgloss.Color("#2ECC71"),
	// 青
	lipgloss.Color("#00CED1"),
	lipgloss.Color("#20B2AA"),
	// 蓝
	lipgloss.Color("#1E90FF"),
	lipgloss.Color("#87CEEB"),
	lipgloss.Color("#5499C7"),
	// 紫
	lipgloss.Color("#855EFC"),
	lipgloss.Color("#6A5ACD"),
	lipgloss.Color("#8E44AD"),
	lipgloss.Color("#9B59B6"),
	// 品红
	lipgloss.Color("#FF69B4"),
	lipgloss.Color("#DB7093"),
}

var previousColor int // 上一次随机颜色的索引

// GetColor 随机获取多个颜色
//
// 参数：
//   - count: 颜色数量
//
// 返回：
//   - 颜色代码
func GetColor(count int) []lipgloss.Color {
	var colors []lipgloss.Color

	enabledColorLength := len(availableColors)

	if enabledColorLength == 0 { // 没有可用的颜色，则返回默认颜色
		for i := 0; i < count; i++ {
			colors = append(colors, DefaultColor)
		}
	} else {
		for i := 0; i < count; i++ {
			// 随机取一个颜色
			index := rand.Intn(enabledColorLength - 1)
			if index == previousColor { // 如果 index == previousColor
				if index < enabledColorLength { // 且未到 availableColors 最后一个元素，则 index + 1
					index++
				} else { // 否则 index 等于 availableColors 最后一个元素的下标
					index = enabledColorLength - 1
				}
			} else {
				previousColor = index
			}
			colors = append(colors, availableColors[index])
		}
	}

	return colors
}
