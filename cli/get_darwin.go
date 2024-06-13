//go:build darwin

/*
File: get_darwin.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-11 10:37:20

Description: 子命令 'get' 的实现
*/

package cli

import (
	"reflect"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/gookit/color"
	"github.com/pelletier/go-toml"
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
		color.Danger.Printf("Load config error (%s:%d): %s\n", fileName, lineNo+1, err)
		return
	}

	// 设置配置项默认值
	var (
		colorful          bool   = config.Main.Colorful
		cpuCacheUnit      string = "KB"
		MemoryDataUnit    string = "GB"
		memoryPercentUnit string = "%"
		SwapDataUnit      string = "GB"
	)

	// 系统信息分配到不同的参数
	sysInfo.GetSysInfo()

	// 计算有多少个 Flag 要显示
	viewQuantity := general.MapBoolCounter(flags, true)

	// 获取随机颜色切片
	if colorful {
		colors = general.GetColor(viewQuantity * 2) //因为分奇数/偶数行，所以要乘2
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
		productInfo := general.GetProductInfo(sysInfo)
		items = config.Genealogy.Product.Items

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

	if flags["boardFlag"] {
		boardPart := func() string {
			partName := general.PartName["Board"][general.Language]
			if partName == "" {
				partName = "Board"
			}
			return partName
		}()

		// 获取数据
		boardInfo := general.GetBoardInfo(sysInfo)
		items = config.Genealogy.Board.Items

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

	if flags["biosFlag"] {
		biosPart := func() string {
			partName := general.PartName["BIOS"][general.Language]
			if partName == "" {
				partName = "BIOS"
			}
			return partName
		}()

		// 获取数据
		biosInfo := general.GetBIOSInfo(sysInfo)
		items = config.Genealogy.Bios.Items

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
		cpuInfo := general.GetCPUInfo(sysInfo, cpuCacheUnit)
		items = config.Genealogy.CPU.Items

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
			MemoryDataUnit = config.Genealogy.Memory.DataUnit
		} else {
			color.Warn.Println("Config file is missing 'memory.data_unit' item, using default value")
		}
		if config.Genealogy.Memory.PercentUnit != "" {
			memoryPercentUnit = config.Genealogy.Memory.PercentUnit
		} else {
			color.Warn.Println("Config file is missing 'memory.percent_unit' item, using default value")
		}

		// 获取数据
		memoryInfo := general.GetMemoryInfo(MemoryDataUnit, memoryPercentUnit)
		items = config.Genealogy.Memory.Items

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
			SwapDataUnit = config.Genealogy.Swap.DataUnit
		} else {
			color.Warn.Println("Config file is missing 'swap.data_unit' item, using default value")
		}

		// 获取数据
		swapInfo := general.GetSwapInfo(SwapDataUnit)
		if swapInfo["SwapStatus"] == "Unavailable" {
			items = config.Genealogy.Swap.Items.Unavailable
		} else {
			items = config.Genealogy.Swap.Items.Available
		}

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

	if flags["storageFlag"] {
		diskPart := func() string {
			partName := general.PartName["Disk"][general.Language]
			if partName == "" {
				partName = "Disk"
			}
			return partName
		}()

		// 获取数据
		storageInfo := general.GetStorageInfo()
		items = config.Genealogy.Storage.Items

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

	if flags["osFlag"] {
		osPart := func() string {
			partNname := general.PartName["OS"][general.Language]
			if partNname == "" {
				partNname = "OS"
			}
			return partNname
		}()

		// 获取数据
		osInfo := general.GetOSInfo(sysInfo)
		items = config.Genealogy.OS.Items

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

	if flags["loadFlag"] {
		loadPart := func() string {
			partName := general.PartName["Load"][general.Language]
			if partName == "" {
				partName = "Load"
			}
			return partName
		}()

		// 获取数据
		loadInfo := general.GetLoadInfo()
		items = config.Genealogy.Load.Items

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
			if item == "Process" {
				cellData = color.Sprintf("%d", loadInfo[item].(uint64))
			} else {
				cellData = color.Sprintf("%.2f", loadInfo[item].(float64))
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

	if flags["userFlag"] {
		userPart := func() string {
			partName := general.PartName["User"][general.Language]
			if partName == "" {
				partName = "User"
			}
			return partName
		}()

		// 获取数据
		userInfo := general.GetUserInfo()
		items = config.Genealogy.User.Items

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

// GrabInformationToTab 抓取信息，各种信息通过标签交互展示
//
// 参数：
//   - configTree: 解析 toml 配置文件得到的配置树
func GrabInformationToTab(configTree *toml.Tree) {
	// 获取配置项
	config, err := general.LoadConfigToStruct(configTree)
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Danger.Printf("Load config error (%s:%d): %s\n", fileName, lineNo+1, err)
		return
	}

	// 设置配置项默认值
	var (
		colorful          bool   = config.Main.Colorful
		cpuCacheUnit      string = "KB"
		MemoryDataUnit    string = "GB"
		memoryPercentUnit string = "%"
		SwapDataUnit      string = "GB"
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
		colors = general.GetColor(viewQuantity * 2) //因为分奇数/偶数行，所以要乘2
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
	productInfo := general.GetProductInfo(sysInfo)
	items = config.Genealogy.Product.Items

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

	// ---------- Board
	boardPart := func() string {
		partName := general.PartName["Board"][general.Language]
		if partName == "" {
			partName = "Board"
		}
		return partName
	}()

	// 获取数据
	boardInfo := general.GetBoardInfo(sysInfo)
	items = config.Genealogy.Board.Items

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

	// ---------- Bios
	biosPart := func() string {
		partName := general.PartName["BIOS"][general.Language]
		if partName == "" {
			partName = "BIOS"
		}
		return partName
	}()

	// 获取数据
	biosInfo := general.GetBIOSInfo(sysInfo)
	items = config.Genealogy.Bios.Items

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
	cpuInfo := general.GetCPUInfo(sysInfo, cpuCacheUnit)
	items = config.Genealogy.CPU.Items

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
		MemoryDataUnit = config.Genealogy.Memory.DataUnit
	} else {
		color.Warn.Println("Config file is missing 'memory.data_unit' item, using default value")
	}
	if config.Genealogy.Memory.PercentUnit != "" {
		memoryPercentUnit = config.Genealogy.Memory.PercentUnit
	} else {
		color.Warn.Println("Config file is missing 'memory.percent_unit' item, using default value")
	}

	// 获取数据
	memoryInfo := general.GetMemoryInfo(MemoryDataUnit, memoryPercentUnit)
	items = config.Genealogy.Memory.Items

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
		SwapDataUnit = config.Genealogy.Swap.DataUnit
	} else {
		color.Warn.Println("Config file is missing 'swap.data_unit' item, using default value")
	}

	// 获取数据
	swapInfo := general.GetSwapInfo(SwapDataUnit)
	if swapInfo["SwapStatus"] == "Unavailable" {
		items = config.Genealogy.Swap.Items.Unavailable
	} else {
		items = config.Genealogy.Swap.Items.Available
	}

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

	// ---------- Storage
	diskPart := func() string {
		partName := general.PartName["Disk"][general.Language]
		if partName == "" {
			partName = "Disk"
		}
		return partName
	}()

	// 获取数据
	storageInfo := general.GetStorageInfo()
	items = config.Genealogy.Storage.Items

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

	// ---------- OS
	osPart := func() string {
		partNname := general.PartName["OS"][general.Language]
		if partNname == "" {
			partNname = "OS"
		}
		return partNname
	}()

	// 获取数据
	osInfo := general.GetOSInfo(sysInfo)
	items = config.Genealogy.OS.Items

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

	// ---------- Load
	loadPart := func() string {
		partName := general.PartName["Load"][general.Language]
		if partName == "" {
			partName = "Load"
		}
		return partName
	}()

	// 获取数据
	loadInfo := general.GetLoadInfo()
	items = config.Genealogy.Load.Items

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
		if item == "Process" {
			cellData = color.Sprintf("%d", loadInfo[item].(uint64))
		} else {
			cellData = color.Sprintf("%.2f", loadInfo[item].(float64))
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

	// ---------- User
	userPart := func() string {
		partName := general.PartName["User"][general.Language]
		if partName == "" {
			partName = "User"
		}
		return partName
	}()

	// 获取数据
	userInfo := general.GetUserInfo()
	items = config.Genealogy.User.Items

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

	if err := general.TabSelector(tabs, tabContents); err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Danger.Printf("Tab selector error (%s:%d): %s\n", fileName, lineNo+1, err)
	}
}
