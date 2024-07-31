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
	renderer    = lipgloss.NewRenderer(os.Stdout)                                                                   // 创建一个 lipgloss 渲染器
	HeaderStyle = renderer.NewStyle().Align(lipgloss.Center).Padding(0, 1, 0, 1).Bold(true).Foreground(HeaderColor) // 表头样式
	BorderStyle = renderer.NewStyle().Foreground(BorderColor)                                                       // 边框样式
	CellStyle   = renderer.NewStyle().Align(lipgloss.Center).Padding(0, 1, 0, 1).Bold(false)                        // 单元格样式

	highlightColor    = lipgloss.AdaptiveColor{Light: TabLightColor, Dark: TabDarkColor}
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")                                                                                                               // 不活跃标签页边框
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")                                                                                                               // 活跃标签页边框
	tabStyle          = lipgloss.NewStyle().Padding(0, 1, 0, 1)                                                                                                          // 标签整体样式
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).Padding(0, 2, 0, 2).BorderForeground(highlightColor)                                         // 不活跃标签页样式
	activeTabStyle    = inactiveTabStyle.Border(activeTabBorder, true).Padding(0, 3, 0, 3)                                                                               // 活跃标签页样式
	windowStyle       = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).UnsetBorderTop().Padding(1, 0, 1, 0).Align(lipgloss.Center).BorderForeground(highlightColor) // 窗口样式
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
