/*
File: get.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-07-29 14:05:12

Description: 子命令 'get' 的实现
*/

package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/yhyj/eniac/general"
	"github.com/zcalusic/sysinfo"
)

// 表格参数
var (
	items        []string                                     // 输出项名称
	oddRowColor  = general.DefaultColor                       // 奇数行颜色
	evenRowColor = general.DefaultColor                       // 偶数行颜色
	oddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
	evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

	dataTable   *table.Table // 创建一个表格
	tableHeader []string     // 表头
	tableData   [][]string   // 表数据
	rowData     []string     // 行数据

	colors []lipgloss.Color // 随机颜色切片
)

// 采集系统信息
var sysInfo sysinfo.SysInfo
