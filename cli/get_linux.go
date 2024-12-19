//go:build linux

/*
File: get_linux.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-11 10:37:20

Description: 子命令 'get' 的实现
*/

package cli

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/gookit/color"
	"github.com/pelletier/go-toml"
	"github.com/yhyj/eniac/general"
)

// GrabInformationToTable 抓取信息，各种信息分别输出为表格
//
// 参数：
//   - configTree: 解析 toml 配置文件得到的配置树
//   - flags: 系统信息各部分的开关
func GrabInformationToTable(configTree *toml.Tree, flags map[string]bool) {
	// 获取配置项
	config, err := general.LoadConfigToStruct(configTree)
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	// 设置配置项默认值
	var (
		colorful             bool   = config.Main.Colorful
		cpuCacheUnit         string = "KB"
		memoryDataUnit       string = "GB"
		memoryPercentUnit    string = "%"
		swapDataUnit         string = "GB"
		basis                string = config.Genealogy.Update.Basis
		owner                string = "user"
		archUpdateRecordFile string = "/tmp/checker-arch.log"
		aurUpdateRecordFile  string = "/tmp/checker-aur.log"
	)

	// 系统信息分配到不同的参数
	sysInfo.GetSysInfo()

	// 计算有多少个 Flag 要显示
	viewQuantity := general.MapBoolCounter(flags, true)

	// 获取随机颜色切片
	if colorful {
		colors = general.GetColor(viewQuantity * 2) // 因为分奇数/偶数行，所以要乘2
	}

	// 执行对应函数
	if flags["productFlag"] {
		productPart := func() string {
			partName := general.PartName["Product"][general.Language]
			if partName == "" {
				partName = "Product"
			}
			return partName
		}()

		// 获取数据
		productInfo := general.GetProductInfo(sysInfo) // 原始数据
		items = config.Genealogy.Product.Items         // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "Product items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}      // 表头
			tableData = [][]string{}        // 表数据
			rowData = []string{productPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, productInfo[item].(string))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["boardFlag"] {
		boardPart := func() string {
			partName := general.PartName["Board"][general.Language]
			if partName == "" {
				partName = "Board"
			}
			return partName
		}()

		// 获取数据
		boardInfo := general.GetBoardInfo(sysInfo) // 原始数据
		items = config.Genealogy.Board.Items       // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "Board items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}    // 表头
			tableData = [][]string{}      // 表数据
			rowData = []string{boardPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, boardInfo[item].(string))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["biosFlag"] {
		biosPart := func() string {
			partName := general.PartName["BIOS"][general.Language]
			if partName == "" {
				partName = "BIOS"
			}
			return partName
		}()

		// 获取数据
		biosInfo := general.GetBIOSInfo(sysInfo) // 原始数据
		items = config.Genealogy.Bios.Items      // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "BIOS items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}   // 表头
			tableData = [][]string{}     // 表数据
			rowData = []string{biosPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, biosInfo[item].(string))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["cpuFlag"] {
		cpuPart := func() string {
			partName := general.PartName["CPU"][general.Language]
			if partName == "" {
				partName = "CPU"
			}
			return partName
		}()

		// 获取 CPU 配置项
		if config.Genealogy.CPU.CacheUnit != "" {
			cpuCacheUnit = config.Genealogy.CPU.CacheUnit
		} else {
			color.Warn.Println("Config file is missing 'cpu.cache_unit' item, using default value")
		}

		// 获取数据
		cpuInfo := general.GetCPUInfo(sysInfo, cpuCacheUnit) // 原始数据
		items = config.Genealogy.CPU.Items                   // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "CPU items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}  // 表头
			tableData = [][]string{}    // 表数据
			rowData = []string{cpuPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, color.Sprintf("%v", cpuInfo[item]))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["gpuFlag"] {
		gpuPart := func() string {
			partName := general.PartName["GPU"][general.Language]
			if partName == "" {
				partName = "GPU"
			}
			return partName
		}()

		// 获取数据
		gpuInfo := general.GetGPUInfo()    // 原始数据
		items = config.Genealogy.GPU.Items // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "GPU items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}  // 表头
			tableData = [][]string{}    // 表数据
			rowData = []string{gpuPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, gpuInfo[item].(string))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["memoryFlag"] {
		memoryPart := func() string {
			partName := general.PartName["Memory"][general.Language]
			if partName == "" {
				partName = "Memory"
			}
			return partName
		}()

		// 获取 Memory 配置项
		if config.Genealogy.Memory.DataUnit != "" {
			memoryDataUnit = config.Genealogy.Memory.DataUnit
		} else {
			color.Warn.Println("Config file is missing 'memory.data_unit' item, using default value")
		}
		if config.Genealogy.Memory.PercentUnit != "" {
			memoryPercentUnit = config.Genealogy.Memory.PercentUnit
		} else {
			color.Warn.Println("Config file is missing 'memory.percent_unit' item, using default value")
		}

		// 获取数据
		memoryInfo := general.GetMemoryInfo(memoryDataUnit, memoryPercentUnit) // 原始数据
		items = config.Genealogy.Memory.Items                                  // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "Memory items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}     // 表头
			tableData = [][]string{}       // 表数据
			rowData = []string{memoryPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, memoryInfo[item].(string))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["swapFlag"] {
		swapPart := func() string {
			partName := general.PartName["Swap"][general.Language]
			if partName == "" {
				partName = "Swap"
			}
			return partName
		}()

		// 获取 Swap 配置项
		if config.Genealogy.Swap.DataUnit != "" {
			swapDataUnit = config.Genealogy.Swap.DataUnit
		} else {
			color.Warn.Println("Config file is missing 'swap.data_unit' item, using default value")
		}

		// 获取数据
		swapInfo := general.GetSwapInfo(swapDataUnit) // 原始数据
		if swapInfo["SwapStatus"] == "Unavailable" {
			items = config.Genealogy.Swap.Items.Unavailable // 原始表头
		} else {
			items = config.Genealogy.Swap.Items.Available // 原始表头
		}

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "Swap items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}   // 表头
			tableData = [][]string{}     // 表数据
			rowData = []string{swapPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, color.Sprintf("%v", swapInfo[item]))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["storageFlag"] {
		diskPart := func() string {
			partName := general.PartName["Disk"][general.Language]
			if partName == "" {
				partName = "Disk"
			}
			return partName
		}()

		// 获取数据
		storageInfo := general.GetStorageInfo() // 原始数据
		items = config.Genealogy.Storage.Items  // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "Storage items is empty")
		} else {
			// 组装表
			tableHeader = []string{""} // 表头
			tableData = [][]string{}   // 表数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
			}
			for index := 1; index <= len(storageInfo); index++ {
				rowData = []string{diskPart + strconv.Itoa(index)} // 行数据
				for _, item := range items {
					rowData = append(rowData, storageInfo[strconv.Itoa(index)].(map[string]interface{})[item].(string))
				}
				tableData = append(tableData, rowData)
			}

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["nicFlag"] {
		nicPart := func() string {
			partName := general.PartName["NIC"][general.Language]
			if partName == "" {
				partName = "NIC"
			}
			return partName
		}()

		// 获取数据
		nicInfo := general.GetNicInfo()    // 原始数据
		items = config.Genealogy.Nic.Items // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "Nic items is empty")
		} else {
			// 组装表
			tableHeader = []string{""} // 表头
			tableData = [][]string{}   // 表数据
			for _, item := range items {
				itemI18n := func() string {
					itemNname := general.GenealogyName[item][general.Language]
					if itemNname == "" {
						itemNname = item
					}
					return itemNname
				}()
				tableHeader = append(tableHeader, itemI18n)
			}
			for index := 1; index <= len(nicInfo); index++ {
				rowData = []string{nicPart + strconv.Itoa(index)} // 行数据
				for _, item := range items {
					rowData = append(rowData, nicInfo[strconv.Itoa(index)].(map[string]interface{})[item].(string))
				}
				tableData = append(tableData, rowData)
			}

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["osFlag"] {
		osPart := func() string {
			partNname := general.PartName["OS"][general.Language]
			if partNname == "" {
				partNname = "OS"
			}
			return partNname
		}()

		// 获取数据
		osInfo := general.GetOSInfo(sysInfo) // 原始数据
		items = config.Genealogy.OS.Items    // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "OS items is empty")
		} else {
			// 组装表
			tableHeader = []string{""} // 表头
			tableData = [][]string{}   // 表数据
			rowData = []string{osPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, osInfo[item].(string))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["loadFlag"] {
		loadPart := func() string {
			partName := general.PartName["Load"][general.Language]
			if partName == "" {
				partName = "Load"
			}
			return partName
		}()

		// 获取数据
		loadInfo := general.GetLoadInfo()   // 原始数据
		items = config.Genealogy.Load.Items // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "Load items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}   // 表头
			tableData = [][]string{}     // 表数据
			rowData = []string{loadPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)

				var cellData string
				switch info := loadInfo[item].(type) {
				case uint64:
					cellData = color.Sprintf("%d", info)
				case float64:
					cellData = color.Sprintf("%.2f", info)
				default:
					cellData = color.Sprintf("%v", info)
				}
				rowData = append(rowData, cellData)
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["timeFlag"] {
		timePart := func() string {
			partName := general.PartName["Time"][general.Language]
			if partName == "" {
				partName = "Time"
			}
			return partName
		}()

		// 获取数据
		timeInfo, _ := general.GetTimeInfo() // 原始数据
		items = config.Genealogy.Time.Items  // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "Time items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}   // 表头
			tableData = [][]string{}     // 表数据
			rowData = []string{timePart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, timeInfo[item].(string))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["userFlag"] {
		userPart := func() string {
			partName := general.PartName["User"][general.Language]
			if partName == "" {
				partName = "User"
			}
			return partName
		}()

		// 获取数据
		userInfo := general.GetUserInfo()   // 原始数据
		items = config.Genealogy.User.Items // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "User items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}   // 表头
			tableData = [][]string{}     // 表数据
			rowData = []string{userPart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)
				rowData = append(rowData, userInfo[item].(string))
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["packageFlag"] {
		packagePart := func() string {
			partName := general.PartName["Package"][general.Language]
			if partName == "" {
				partName = "Package"
			}
			return partName
		}()

		// 获取数据
		packageInfo, _ := general.GetPackageInfo() // 原始数据
		items = config.Genealogy.Package.Items     // 原始表头

		// 未配置表头时不显示该项，发送通知
		if len(items) == 0 {
			general.Notifier = append(general.Notifier, "Package items is empty")
		} else {
			// 组装表
			tableHeader = []string{""}      // 表头
			tableData = [][]string{}        // 表数据
			rowData = []string{packagePart} // 行数据
			for _, item := range items {
				itemI18n := func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, itemI18n)

				var cellData string
				switch info := packageInfo[item].(type) {
				case int:
					cellData = color.Sprintf("%d", info)
				default:
					cellData = color.Sprintf("%v", info)
				}
				rowData = append(rowData, cellData)
			}
			tableData = append(tableData, rowData)

			// 获取随机颜色
			oddRowColor = colors[0]
			evenRowColor = colors[1]
			colors = colors[2:]

			oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
			evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

			dataTable = table.New()                                 // 创建一个表格
			dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
			dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = evenRowStyle // 偶数行
				default:
					style = oddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			dataTable.Headers(tableHeader...) // 设置表头
			dataTable.Rows(tableData...)      // 设置单元格

			color.Println(dataTable)
		}
	}

	if flags["updateFlag"] {
		// 获取 update 配置项
		if config.Genealogy.Update.ArchRecordFile != "" {
			archUpdateRecordFile = config.Genealogy.Update.ArchRecordFile
		} else {
			color.Warn.Println("Config file is missing 'update.arch_record_file' item, using default value")
		}
		if config.Genealogy.Update.AurRecordFile != "" {
			aurUpdateRecordFile = config.Genealogy.Update.AurRecordFile
		} else {
			color.Warn.Println("Config file is missing 'update.aur_record_file' item, using default value")
		}

		if flags["onlyFlag"] {
			// 仅输出不带额外格式的可更新包信息，专为第三方更新检测插件服务
			updatablePackageInfo, _ := general.GetUpdatablePackageInfo(archUpdateRecordFile, aurUpdateRecordFile, 0)
			color.Println("Arch Official Repository:")
			for num, info := range updatablePackageInfo["UpdatablePackageList"].([]string) {
				color.Printf("%v: %v\n", num+1, info)
			}
			color.Println()
			color.Println("Arch User Repository:")
			for num, info := range updatablePackageInfo["AurPackageList"].([]string) {
				color.Printf("%v: %v\n", num+1, info)
			}
		} else {
			updatePart := func() string {
				partName := general.PartName["Update"][general.Language]
				if partName == "" {
					partName = "Update"
				}
				return partName
			}()

			// 获取数据
			checkUpdateDaemonInfo, _ := general.GetCheckUpdateDaemonInfo(basis, owner)                               // 原始数据
			updatablePackageInfo, _ := general.GetUpdatablePackageInfo(archUpdateRecordFile, aurUpdateRecordFile, 0) // 原始数据
			updateInfo := make(map[string]interface{})
			// 合并两部分数据
			for key, value := range checkUpdateDaemonInfo {
				updateInfo[key] = value
			}
			for key, value := range updatablePackageInfo {
				updateInfo[key] = value
			}
			items = config.Genealogy.Update.Items // 原始表头

			// 未配置表头时不显示该项，发送通知
			if len(items) == 0 {
				general.Notifier = append(general.Notifier, "Update items is empty")
			} else {
				// 组装表
				tableHeader = []string{""}     // 表头
				tableData = [][]string{}       // 表数据
				rowData = []string{updatePart} // 行数据
				for _, item := range items {
					itemI18n := func() string {
						itemName := general.GenealogyName[item][general.Language]
						if itemName == "" {
							itemName = item
						}
						return itemName
					}()
					tableHeader = append(tableHeader, itemI18n)

					var cellData string
					switch info := updateInfo[item].(type) {
					case []string:
						cellData = strings.Join(info, "\n")
					default:
						cellData = color.Sprintf("%v", info)
					}
					rowData = append(rowData, cellData)
				}
				tableData = append(tableData, rowData)

				// 获取随机颜色
				oddRowColor = colors[0]
				evenRowColor = colors[1]
				colors = colors[2:]

				oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
				evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

				dataTable = table.New()                                 // 创建一个表格
				dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
				dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
				dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
					var style lipgloss.Style

					switch {
					case row == 0:
						return general.HeaderStyle // 第一行为表头
					case row%2 == 0:
						style = evenRowStyle // 偶数行
					default:
						style = oddRowStyle // 奇数行
					}

					// 设置特定列格式
					switch col {
					case 0:
						style = style.Foreground(general.ColumnOneColor)
					case len(items):
						style = style.Align(lipgloss.Left)
					}

					return style
				})

				dataTable.Headers(tableHeader...) // 设置表头
				dataTable.Rows(tableData...)      // 设置单元格

				color.Println(dataTable)
			}
		}
	}
}

// GrabInformationToTab 抓取信息，各种信息通过标签交互展示
//
// 参数：
//   - configTree: 解析 toml 配置文件得到的配置树
func GrabInformationToTab(configTree *toml.Tree) {
	// 获取配置项
	config, err := general.LoadConfigToStruct(configTree)
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	// 设置配置项默认值
	var (
		colorful             bool   = config.Main.Colorful
		cycle                bool   = config.Main.Cycle
		cpuCacheUnit         string = "KB"
		memoryDataUnit       string = "GB"
		memoryPercentUnit    string = "%"
		swapDataUnit         string = "GB"
		basis                string = config.Genealogy.Update.Basis
		owner                string = "user"
		archUpdateRecordFile string = "/tmp/checker-arch.log"
		aurUpdateRecordFile  string = "/tmp/checker-aur.log"
	)

	// Tab 参数
	var (
		tabs        []string // 标签内容
		tabContents []string // 标签内容
	)

	// 系统信息分配到不同的参数
	sysInfo.GetSysInfo()

	// 计算有多少个 Flag 要显示
	viewQuantity := reflect.TypeOf(config.Genealogy).NumField()

	// 获取随机颜色切片
	if colorful {
		colors = general.GetColor(viewQuantity * 2) // 因为分奇数/偶数行，所以要乘2
	}

	// ---------- Product
	productPart := func() string {
		partName := general.PartName["Product"][general.Language]
		if partName == "" {
			partName = "Product"
		}
		return partName
	}()

	// 获取数据
	productInfo := general.GetProductInfo(sysInfo) // 原始数据
	items = config.Genealogy.Product.Items         // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, productInfo[item].(string))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, productPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- Board
	boardPart := func() string {
		partName := general.PartName["Board"][general.Language]
		if partName == "" {
			partName = "Board"
		}
		return partName
	}()

	// 获取数据
	boardInfo := general.GetBoardInfo(sysInfo) // 原始数据
	items = config.Genealogy.Board.Items       // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, boardInfo[item].(string))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, boardPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- Bios
	biosPart := func() string {
		partName := general.PartName["BIOS"][general.Language]
		if partName == "" {
			partName = "BIOS"
		}
		return partName
	}()

	// 获取数据
	biosInfo := general.GetBIOSInfo(sysInfo) // 原始数据
	items = config.Genealogy.Bios.Items      // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, biosInfo[item].(string))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, biosPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- CPU
	cpuPart := func() string {
		partName := general.PartName["CPU"][general.Language]
		if partName == "" {
			partName = "CPU"
		}
		return partName
	}()

	// 获取 CPU 配置项
	if config.Genealogy.CPU.CacheUnit != "" {
		cpuCacheUnit = config.Genealogy.CPU.CacheUnit
	} else {
		color.Warn.Println("Config file is missing 'cpu.cache_unit' item, using default value")
	}

	// 获取数据
	cpuInfo := general.GetCPUInfo(sysInfo, cpuCacheUnit) // 原始数据
	items = config.Genealogy.CPU.Items                   // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, color.Sprintf("%v", cpuInfo[item]))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, cpuPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- GPU
	gpuPart := func() string {
		partName := general.PartName["GPU"][general.Language]
		if partName == "" {
			partName = "GPU"
		}
		return partName
	}()

	// 获取数据
	gpuInfo := general.GetGPUInfo()    // 原始数据
	items = config.Genealogy.GPU.Items // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, gpuInfo[item].(string))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, gpuPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- Memory
	memoryPart := func() string {
		partName := general.PartName["Memory"][general.Language]
		if partName == "" {
			partName = "Memory"
		}
		return partName
	}()

	// 获取 Memory 配置项
	if config.Genealogy.Memory.DataUnit != "" {
		memoryDataUnit = config.Genealogy.Memory.DataUnit
	} else {
		color.Warn.Println("Config file is missing 'memory.data_unit' item, using default value")
	}
	if config.Genealogy.Memory.PercentUnit != "" {
		memoryPercentUnit = config.Genealogy.Memory.PercentUnit
	} else {
		color.Warn.Println("Config file is missing 'memory.percent_unit' item, using default value")
	}

	// 获取数据
	memoryInfo := general.GetMemoryInfo(memoryDataUnit, memoryPercentUnit) // 原始数据
	items = config.Genealogy.Memory.Items                                  // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, memoryInfo[item].(string))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, memoryPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- Swap
	swapPart := func() string {
		partName := general.PartName["Swap"][general.Language]
		if partName == "" {
			partName = "Swap"
		}
		return partName
	}()

	// 获取 Swap 配置项
	if config.Genealogy.Swap.DataUnit != "" {
		swapDataUnit = config.Genealogy.Swap.DataUnit
	} else {
		color.Warn.Println("Config file is missing 'swap.data_unit' item, using default value")
	}

	// 获取数据
	swapInfo := general.GetSwapInfo(swapDataUnit) // 原始数据
	if swapInfo["SwapStatus"] == "Unavailable" {
		items = config.Genealogy.Swap.Items.Unavailable // 原始表头
	} else {
		items = config.Genealogy.Swap.Items.Available // 原始表头
	}

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, color.Sprintf("%v", swapInfo[item]))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, swapPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- Storage
	diskPart := func() string {
		partName := general.PartName["Disk"][general.Language]
		if partName == "" {
			partName = "Disk"
		}
		return partName
	}()

	// 获取数据
	storageInfo := general.GetStorageInfo() // 原始数据
	items = config.Genealogy.Storage.Items  // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
		}
		for index := 1; index <= len(storageInfo); index++ {
			rowData = []string{} // 行数据
			for _, item := range items {
				rowData = append(rowData, storageInfo[strconv.Itoa(index)].(map[string]interface{})[item].(string))
			}
			tableData = append(tableData, rowData)
		}

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, diskPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- NIC
	nicPart := func() string {
		partName := general.PartName["NIC"][general.Language]
		if partName == "" {
			partName = "NIC"
		}
		return partName
	}()

	// 获取数据
	nicInfo := general.GetNicInfo()    // 原始数据
	items = config.Genealogy.Nic.Items // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		for _, item := range items {
			itemI18n := func() string {
				itemNname := general.GenealogyName[item][general.Language]
				if itemNname == "" {
					itemNname = item
				}
				return itemNname
			}()
			tableHeader = append(tableHeader, itemI18n)
		}
		for index := 1; index <= len(nicInfo); index++ {
			rowData = []string{} // 行数据
			for _, item := range items {
				rowData = append(rowData, nicInfo[strconv.Itoa(index)].(map[string]interface{})[item].(string))
			}
			tableData = append(tableData, rowData)
		}

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, nicPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- OS
	osPart := func() string {
		partNname := general.PartName["OS"][general.Language]
		if partNname == "" {
			partNname = "OS"
		}
		return partNname
	}()

	// 获取数据
	osInfo := general.GetOSInfo(sysInfo) // 原始数据
	items = config.Genealogy.OS.Items    // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, osInfo[item].(string))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, osPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- Load
	loadPart := func() string {
		partName := general.PartName["Load"][general.Language]
		if partName == "" {
			partName = "Load"
		}
		return partName
	}()

	// 获取数据
	loadInfo := general.GetLoadInfo()   // 原始数据
	items = config.Genealogy.Load.Items // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)

			var cellData string
			switch info := loadInfo[item].(type) {
			case uint64:
				cellData = color.Sprintf("%d", info)
			case float64:
				cellData = color.Sprintf("%.2f", info)
			default:
				cellData = color.Sprintf("%v", info)
			}
			rowData = append(rowData, cellData)
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, loadPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- Time
	timePart := func() string {
		partName := general.PartName["Time"][general.Language]
		if partName == "" {
			partName = "Time"
		}
		return partName
	}()

	// 获取数据
	timeInfo, _ := general.GetTimeInfo() // 原始数据
	items = config.Genealogy.Time.Items  // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, timeInfo[item].(string))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, timePart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- User
	userPart := func() string {
		partName := general.PartName["User"][general.Language]
		if partName == "" {
			partName = "User"
		}
		return partName
	}()

	// 获取数据
	userInfo := general.GetUserInfo()   // 原始数据
	items = config.Genealogy.User.Items // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			rowData = append(rowData, userInfo[item].(string))
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, userPart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- Package
	packagePart := func() string {
		partName := general.PartName["Package"][general.Language]
		if partName == "" {
			partName = "Package"
		}
		return partName
	}()

	// 获取数据
	packageInfo, _ := general.GetPackageInfo() // 原始数据
	items = config.Genealogy.Package.Items     // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)
			var cellData string
			switch info := packageInfo[item].(type) {
			case int:
				cellData = color.Sprintf("%d", info)
			default:
				cellData = color.Sprintf("%v", info)
			}
			rowData = append(rowData, cellData)
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, packagePart)
		tabContents = append(tabContents, dataTable.String())
	}

	// ---------- Update
	updatePart := func() string {
		partName := general.PartName["Update"][general.Language]
		if partName == "" {
			partName = "Update"
		}
		return partName
	}()

	// 获取 update 配置项
	if config.Genealogy.Update.ArchRecordFile != "" {
		archUpdateRecordFile = config.Genealogy.Update.ArchRecordFile
	} else {
		color.Warn.Println("Config file is missing 'update.arch_record_file' item, using default value")
	}
	if config.Genealogy.Update.AurRecordFile != "" {
		aurUpdateRecordFile = config.Genealogy.Update.AurRecordFile
	} else {
		color.Warn.Println("Config file is missing 'update.aur_record_file' item, using default value")
	}

	// 获取数据
	checkUpdateDaemonInfo, _ := general.GetCheckUpdateDaemonInfo(basis, owner)                               // 原始数据
	updatablePackageInfo, _ := general.GetUpdatablePackageInfo(archUpdateRecordFile, aurUpdateRecordFile, 0) // 原始数据
	updateInfo := make(map[string]interface{})
	// 合并两部分数据
	for key, value := range checkUpdateDaemonInfo {
		updateInfo[key] = value
	}
	for key, value := range updatablePackageInfo {
		updateInfo[key] = value
	}
	items = config.Genealogy.Update.Items // 原始表头

	// 未配置表头时不显示该项
	if len(items) != 0 {
		// 组装表
		tableHeader = []string{} // 表头
		tableData = [][]string{} // 表数据
		rowData = []string{}     // 行数据
		for _, item := range items {
			itemI18n := func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, itemI18n)

			var cellData string
			switch info := updateInfo[item].(type) {
			case []string:
				_, height, _ := general.GetTerminalSize()                                       // 终端尺寸
				viewRows := height - (3 + 1) - (3 + 1) - 1 - 1 - 1 - general.TableExPaddingUD*2 // 表格行数（终端行数 -（标签头行数+标签尾行数）-（表格头行数+表格尾行数）- 为省略号留的行数 - 命令行数 - 预留行数 - 数据表外部上下边距）
				if len(info) > viewRows {
					cellData = strings.Join(info[:viewRows], "\n")
					cellData = cellData + "\n......"
				} else {
					cellData = strings.Join(info, "\n")
				}
			default:
				cellData = color.Sprintf("%v", info)
			}
			rowData = append(rowData, cellData)
		}
		tableData = append(tableData, rowData)

		// 获取随机颜色
		oddRowColor = colors[0]
		evenRowColor = colors[1]
		colors = colors[2:]

		oddRowStyle = general.CellStyle.Foreground(oddRowColor)   // 奇数行样式
		evenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式

		dataTable = table.New()
		dataTable.Border(lipgloss.RoundedBorder())              // 设置表格边框
		dataTable.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		dataTable.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = evenRowStyle // 偶数行
			default:
				style = oddRowStyle // 奇数行
			}

			// 设置特定列格式
			if col == len(items)-1 {
				style = style.Align(lipgloss.Left)
			}

			return style
		})

		dataTable.Headers(tableHeader...) // 设置表头
		dataTable.Rows(tableData...)      // 设置单元格

		tabs = append(tabs, updatePart)
		tabContents = append(tabContents, dataTable.String())
	}

	// 输出 Tab
	if err := general.TabSelector(tabs, tabContents, cycle); err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}
}
