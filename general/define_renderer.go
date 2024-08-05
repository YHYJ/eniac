/*
File: define_renderer.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-05-31 11:59:55

Description: 定义渲染器
*/

package general

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	TabExPaddingUD         = 0 // 标签页外部上下边距
	TabExPaddingLR         = 1 // 标签页外部左右边距
	InactiveTabInPaddingUD = 0 // 不活跃标签页内部上下边距
	InactiveTabInPaddingLR = 2 // 不活跃标签页内部左右边距
	ActiveTabInPaddingUD   = 0 // 活跃标签页内部上下边距
	ActiveTabInPaddingLR   = 3 // 活跃标签页内部左右边距
	TableInPaddingUD       = 0 // 数据表内部上下边距
	TableInPaddingLR       = 1 // 数据表内部左右边距
	TableExPaddingUD       = 1 // 数据表外部上下边距
	TableExPaddingLR       = 0 // 数据表外部左右边距

	highlightColor = lipgloss.AdaptiveColor{Light: TabLightColor, Dark: TabDarkColor} // 高亮颜色

	inactiveTabBorder  = tabBorderWithBottom("┴", "─", "┴")                                                                                                           // 不活跃标签页边框
	activeTabBorder    = tabBorderWithBottom("┘", " ", "└")                                                                                                           // 活跃标签页边框
	tabExStyle         = lipgloss.NewStyle().Padding(TabExPaddingUD, TabExPaddingLR)                                                                                  // 标签页外部样式
	inactiveTabInStyle = lipgloss.NewStyle().Border(inactiveTabBorder, true).Padding(InactiveTabInPaddingUD, InactiveTabInPaddingLR).BorderForeground(highlightColor) // 不活跃标签页内部样式
	activeTabInStyle   = inactiveTabInStyle.Border(activeTabBorder, true).Padding(ActiveTabInPaddingUD, ActiveTabInPaddingLR)                                         // 活跃标签页内部样式

	tableExStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).UnsetBorderTop().Padding(TableExPaddingUD, TableExPaddingLR).Align(lipgloss.Center).BorderForeground(highlightColor) // 窗口样式

	renderer    = lipgloss.NewRenderer(os.Stdout)                                                                                           // 创建一个 lipgloss 渲染器
	HeaderStyle = renderer.NewStyle().Align(lipgloss.Center).Padding(TableInPaddingUD, TableInPaddingLR).Bold(true).Foreground(HeaderColor) // 表头样式
	BorderStyle = renderer.NewStyle().Foreground(BorderColor)                                                                               // 边框样式
	CellStyle   = renderer.NewStyle().Align(lipgloss.Center).Padding(TableInPaddingUD, TableInPaddingLR).Bold(false)                        // 单元格样式
)

// tabBorderWithBottom 返回指定样式的边框，用于构建活跃/不活跃标签
//
// 参数：
//   - left: 左边框
//   - middle: 中间框
//   - right: 右边框
//
// 返回：
//   - 指定样式的边框
func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
