/*
File: define_selector.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-10 13:33:59

Description: 定义选择器

- Update, View 等方法通过 model 与用户进行交互
*/

package general

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var quitKey = "q"                 // 默认的退出键
var selectorType = "program name" // 选择器主题

// model 结构体，选择器的数据
type model struct {
	Tabs       []string // 所有标签
	TabContent []string // 标签对应的内容
	ActiveTab  int      // 当前激活的标签
	Cycle      bool     // 是否允许循环切换
}

// Init model 结构体的初始化方法，是 BubbleTea 框架中的一个特殊方法
func (m model) Init() tea.Cmd {
	// 返回 nil 意味着不需要 I/O 操作
	return nil
}

// Update model 结构体的更新方法，是 BubbleTea 框架中的一个特殊方法
//
// 参数：
//   - msg: 包含来自 I/O 操作结果的数据，出发更新功能，并以此出触发 UI 绘制
//
// 返回：
//   - model: 更新后的 model
//   - tea.Cmd: 一个 I/O 操作，完成后会返回一条消息，如果为 nil 则被视为无操作
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// 监控按键事件
	case tea.KeyMsg:
		// 对按下的相应按键做出对应反应
		switch keyPress := msg.String(); keyPress {
		case quitKey, "ctrl+c", "esc":
			return m, tea.Quit
		case "left", "h", "p", "shift+tab":
			if m.Cycle {
				m.ActiveTab--
				m.fixCursor(0, len(m.Tabs)-1)
			} else {
				m.ActiveTab = Max(m.ActiveTab-1, 0)
			}
		case "right", "l", "n", "tab":
			if m.Cycle {
				m.ActiveTab++
				m.fixCursor(0, len(m.Tabs)-1)
			} else {
				m.ActiveTab = Min(m.ActiveTab+1, len(m.Tabs)-1)
			}
		}
	}
	// 将更新后的 model 返回给 BubbleTea 进行处理
	return m, nil
}

// View model 结构体的视图方法，是 BubbleTea 框架中的一个特殊方法
//
// 返回：
//   - string: 绘制内容
func (m model) View() string {
	// 呈现的选项卡
	var renderedTabs []string

	// 构建显示内容
	s := strings.Builder{}

	for i, t := range m.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.ActiveTab
		if isActive {
			style = activeTabInStyle
		} else {
			style = inactiveTabInStyle
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	s.WriteString(row)
	s.WriteString("\n")
	s.WriteString(tableExStyle.Width((lipgloss.Width(row) - tableExStyle.GetHorizontalFrameSize())).Render(m.TabContent[m.ActiveTab]))
	return tabExStyle.Render(s.String())
}

// fixCursor 修正光标位置，防止越界
//
// 参数：
//   - minIndex: 最小光标索引
//   - maxIndex: 最大光标索引
func (m *model) fixCursor(minIndex, maxIndex int) {
	if m.ActiveTab > maxIndex {
		m.ActiveTab = 0
	} else if m.ActiveTab < minIndex {
		m.ActiveTab = maxIndex
	}
}

// TabSelector 标签选择器，接受一个标签切片和一个标签内容切片，显示选中的标签的内容
//
// 参数：
//   - tabs: 所有标签
//   - contents: 所有标签对应的内容
//   - cycle: 是否允许循环切换
//
// 返回：
//   - 错误信息
func TabSelector(tabs, contents []string, cycle bool) error {
	if len(tabs) != len(contents) {
		return fmt.Errorf("Tabs and contents must have the same length")
	}
	m := model{Tabs: tabs, TabContent: contents, Cycle: cycle}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		return fmt.Errorf("Error running program: %s", err)
	}
	return nil
}
