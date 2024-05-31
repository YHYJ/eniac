/*
File: get.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-11 10:37:20

Description: 子命令 'get' 的实现
*/

package cli

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/gookit/color"
	"github.com/pelletier/go-toml"
	"github.com/yhyj/eniac/general"
	"github.com/zcalusic/sysinfo"
)

// GrabSystemInformation 抓取系统信息
// 参数：
//   - configTree: 解析 toml 配置文件得到的配置树
//   - flags: 系统信息各部分的开关
func GrabSystemInformation(configTree *toml.Tree, flags map[string]bool) {
	// 获取配置项
	config, err := general.LoadConfigToStruct(configTree)
	if err != nil {
		color.Danger.Println(err)
		return
	}

	// 设置配置项默认值
	var (
		colorful          bool   = true // TODO: 未读取配置项
		cpuCacheUnit      string = "KB"
		memoryDataUnit    string = "GB"
		memoryPercentUnit string = "%"
		updateRecordFile  string = "/tmp/system-checkupdates.log"
	)

	// 采集系统信息（集中采集一次后分配到不同的参数）
	var sysInfo sysinfo.SysInfo
	sysInfo.GetSysInfo()

	// 表格参数
	var (
		items        []string               // 输出项名称参数
		oddRowColor  = general.DefaultColor // 奇数行颜色
		evenRowColor = general.DefaultColor // 偶数行颜色
	)

	// 执行对应函数
	if flags["productFlag"] {
		productPart := func() string {
			partName := general.PartName["Product"][general.Language]
			if partName == "" {
				partName = "Product"
			}
			return partName
		}()
		// TODO: 是否删除：color.Printf("%s\n", general.FgGrayText("······ Product ······"))

		// 获取数据
		productInfo := general.GetProductInfo(sysInfo)
		items = []string{"ProductVendor", "ProductName"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{productPart}
		for _, item := range items {
			outputValue := productInfo[item].(string)
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["boardFlag"] {
		boardPart := func() string {
			partName := general.PartName["Board"][general.Language]
			if partName == "" {
				partName = "Board"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ Board ······"))

		// 获取数据
		boardInfo := general.GetBoardInfo(sysInfo)
		items = []string{"BoardVendor", "BoardName", "BoardVersion"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{boardPart}
		for _, item := range items {
			outputValue := boardInfo[item].(string)
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["biosFlag"] {
		biosPart := func() string {
			partName := general.PartName["BIOS"][general.Language]
			if partName == "" {
				partName = "BIOS"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ BIOS ······"))

		// 获取数据
		biosInfo := general.GetBIOSInfo(sysInfo)
		items = []string{"BIOSVendor", "BIOSVersion", "BIOSDate"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{biosPart}
		for _, item := range items {
			outputValue := biosInfo[item].(string)
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["cpuFlag"] {
		cpuPart := func() string {
			partName := general.PartName["CPU"][general.Language]
			if partName == "" {
				partName = "CPU"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ CPU ······"))

		// 获取 CPU 配置项
		if config.Genealogy.Cpu.CacheUnit != "" {
			cpuCacheUnit = config.Genealogy.Cpu.CacheUnit
		} else {
			color.Danger.Println("Config file is missing 'cpu.cache_unit' item, using default value")
		}

		// 获取数据
		cpuInfo := general.GetCPUInfo(sysInfo, cpuCacheUnit)
		items = []string{"CPUModel", "CPUNumber", "CPUCores", "CPUThreads", "CPUCache"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{cpuPart}
		for _, item := range items {
			outputValue := color.Sprintf("%v", cpuInfo[item])
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["gpuFlag"] {
		gpuPart := func() string {
			partName := general.PartName["GPU"][general.Language]
			if partName == "" {
				partName = "GPU"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ GPU ······"))

		// 获取数据
		gpuInfo := general.GetGPUInfo()
		items = []string{"GPUAddress", "GPUDriver", "GPUProduct", "GPUVendor"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{gpuPart}
		for _, item := range items {
			outputValue := gpuInfo[item].(string)
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["memoryFlag"] {
		memoryPart := func() string {
			partName := general.PartName["Memory"][general.Language]
			if partName == "" {
				partName = "Memory"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ Memory ······"))

		// 获取 Memory 配置项
		if config.Genealogy.Memory.DataUnit != "" {
			memoryDataUnit = config.Genealogy.Memory.DataUnit
		} else {
			color.Danger.Println("Config file is missing 'memory.data_unit' item, using default value")
		}
		if config.Genealogy.Memory.PercentUnit != "" {
			memoryPercentUnit = config.Genealogy.Memory.PercentUnit
		} else {
			color.Danger.Println("Config file is missing 'memory.percent_unit' item, using default value")
		}

		// 获取数据
		memoryInfo := general.GetMemoryInfo(memoryDataUnit, memoryPercentUnit)
		items = []string{"MemoryUsedPercent", "MemoryTotal", "MemoryUsed", "MemoryAvail", "MemoryFree", "MemoryBuffCache", "MemoryShared"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{memoryPart}
		for _, item := range items {
			outputValue := memoryInfo[item].(string)
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["swapFlag"] {
		swapPart := func() string {
			partName := general.PartName["Swap"][general.Language]
			if partName == "" {
				partName = "Swap"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ Swap ······"))

		// 获取 Memory 配置项
		if config.Genealogy.Memory.DataUnit != "" {
			memoryDataUnit = config.Genealogy.Memory.DataUnit
		} else {
			color.Danger.Println("Config file is missing 'memory.data_unit' item, using default value")
		}

		// 获取数据
		swapInfo := general.GetSwapInfo(memoryDataUnit)

		// 组装表头
		tableHeader := []string{""}
		if swapInfo["SwapStatus"] == "Disabled" {
			items = []string{"SwapStatus"}
			for _, item := range items {
				item = func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, item)
			}
		} else {
			items = []string{"SwapTotal", "SwapFree"}
			for _, item := range items {
				item = func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, item)
			}
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{swapPart}
		for _, item := range items {
			outputValue := color.Sprintf("%v", swapInfo[item])
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["storageFlag"] {
		diskPart := func() string {
			partName := general.PartName["Disk"][general.Language]
			if partName == "" {
				partName = "Disk"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ Storage ······"))

		// 获取数据
		storageInfo := general.GetStorageInfo()
		items = []string{"StorageName", "StorageSize", "StorageType", "StorageDriver", "StorageVendor", "StorageModel", "StorageSerial", "StorageRemovable"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		for index := 1; index <= len(storageInfo); index++ {
			outputInfo := []string{diskPart + strconv.Itoa(index)}
			for _, item := range items {
				outputValue := storageInfo[strconv.Itoa(index)].(map[string]interface{})[item].(string)
				outputInfo = append(outputInfo, outputValue)
			}
			tableData = append(tableData, outputInfo)
		}

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["nicFlag"] {
		nicPart := func() string {
			partName := general.PartName["NIC"][general.Language]
			if partName == "" {
				partName = "NIC"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ Nic ······"))

		// 获取数据
		nicInfo := general.GetNicInfo()
		items = []string{"NicName", "NicMacAddress", "NicDriver", "NicVendor", "NicProduct", "NicPCIAddress", "NicSpeed", "NicDuplex"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemNname := general.GenealogyName[item][general.Language]
				if itemNname == "" {
					itemNname = item
				}
				return itemNname
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		for index := 1; index <= len(nicInfo); index++ {
			outputInfo := []string{nicPart + strconv.Itoa(index)}
			for _, item := range items {
				outputValue := nicInfo[strconv.Itoa(index)].(map[string]interface{})[item].(string)
				outputInfo = append(outputInfo, outputValue)
			}
			tableData = append(tableData, outputInfo)
		}

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["osFlag"] {
		osPart := func() string {
			partNname := general.PartName["OS"][general.Language]
			if partNname == "" {
				partNname = "OS"
			}
			return partNname
		}()
		// color.Printf("%s\n", general.FgGrayText("······ OS ······"))

		// 获取数据
		osInfo := general.GetOSInfo(sysInfo)
		items = []string{"OS", "Kernel", "Platform", "Arch", "TimeZone", "Hostname"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{osPart}
		for _, item := range items {
			outputValue := osInfo[item].(string)
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["loadFlag"] {
		loadPart := func() string {
			partName := general.PartName["Load"][general.Language]
			if partName == "" {
				partName = "Load"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ Load ······"))

		// 获取数据
		loadInfo := general.GetLoadInfo()
		items = []string{"Load1", "Load5", "Load15", "Process"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{loadPart}
		outputValue := ""
		for _, item := range items {
			if item == "Process" {
				outputValue = color.Sprintf("%d", loadInfo[item].(uint64))
			} else {
				outputValue = color.Sprintf("%.2f", loadInfo[item].(float64))
			}
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["timeFlag"] {
		timePart := func() string {
			partName := general.PartName["Time"][general.Language]
			if partName == "" {
				partName = "Time"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ Time ······"))

		// 获取数据
		timeInfo, _ := general.GetTimeInfo()
		items = []string{"StartTime", "Uptime", "BootTime"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{timePart}
		for _, item := range items {
			outputValue := timeInfo[item].(string)
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["userFlag"] {
		userPart := func() string {
			partName := general.PartName["User"][general.Language]
			if partName == "" {
				partName = "User"
			}
			return partName
		}()
		// color.Printf("%s\n", general.FgGrayText("······ User ······"))

		// 获取数据
		userInfo := general.GetUserInfo()
		items = []string{"UserName", "User", "UserUid", "UserGid", "UserHomeDir"}

		// 组装表头
		tableHeader := []string{""}
		for _, item := range items {
			item = func() string {
				itemName := general.GenealogyName[item][general.Language]
				if itemName == "" {
					itemName = item
				}
				return itemName
			}()
			tableHeader = append(tableHeader, item)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{userPart}
		for _, item := range items {
			outputValue := userInfo[item].(string)
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		// 获取随机颜色
		if colorful {
			oddRowColor = general.GetColor()
			evenRowColor = general.GetColor()
		}

		var (
			OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
			EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
		)

		table := table.New()                                // 创建一个表格
		table.Border(lipgloss.RoundedBorder())              // 设置表格边框
		table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
		table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
			var style lipgloss.Style

			switch {
			case row == 0:
				return general.HeaderStyle // 第一行为表头
			case row%2 == 0:
				style = EvenRowStyle // 偶数行
			default:
				style = OddRowStyle // 奇数行
			}

			// 设置第一列格式
			if col == 0 {
				style = style.Foreground(general.ColumnOneColor)
			}

			return style
		})

		table.Headers(tableHeader...) // 设置表头
		table.Rows(tableData...)      // 设置单元格

		color.Println(table)
	}

	if flags["updateFlag"] {
		if flags["onlyFlag"] {
			// 仅输出不带额外格式的可更新包信息，专为第三方更新检测插件服务
			packageInfo, _ := general.GetPackageInfo(updateRecordFile, 0)
			for num, info := range packageInfo["PackageList"].([]string) {
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
			// color.Printf("%s\n", general.FgGrayText("······ Update ······"))

			// 获取 update 配置项
			if config.Genealogy.Update.RecordFile != "" {
				updateRecordFile = config.Genealogy.Update.RecordFile
			} else {
				color.Danger.Println("Config file is missing 'update.record_file' item, using default value")
			}

			// 获取数据
			daemonInfo, _ := general.GetUpdateDaemonInfo()
			packageInfo, _ := general.GetPackageInfo(updateRecordFile, 0)
			updateInfo := make(map[string]interface{})
			// 合并两部分数据
			for key, value := range daemonInfo {
				updateInfo[key] = value
			}
			for key, value := range packageInfo {
				updateInfo[key] = value
			}
			items = []string{"UpdateDaemonStatus", "PackageQuantity", "PackageList"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				item = func() string {
					itemName := general.GenealogyName[item][general.Language]
					if itemName == "" {
						itemName = item
					}
					return itemName
				}()
				tableHeader = append(tableHeader, item)
			}

			// 组装表数据
			tableData := [][]string{}
			outputInfo := []string{updatePart}
			var outputValue string
			for _, item := range items {
				if item == "PackageList" {
					packageList := updateInfo[item].([]string)
					outputValue = strings.Join(packageList, "\n")
				} else {
					outputValue = updateInfo[item].(string)
				}
				outputInfo = append(outputInfo, outputValue)
			}
			tableData = append(tableData, outputInfo)

			// 获取随机颜色
			if colorful {
				oddRowColor = general.GetColor()
				evenRowColor = general.GetColor()
			}

			var (
				OddRowStyle  = general.CellStyle.Foreground(oddRowColor)  // 奇数行样式
				EvenRowStyle = general.CellStyle.Foreground(evenRowColor) // 偶数行样式
			)

			table := table.New()                                // 创建一个表格
			table.Border(lipgloss.RoundedBorder())              // 设置表格边框
			table.BorderStyle(general.BorderStyle)              // 设置表格边框样式
			table.StyleFunc(func(row, col int) lipgloss.Style { // 按位置设置单元格样式
				var style lipgloss.Style

				switch {
				case row == 0:
					return general.HeaderStyle // 第一行为表头
				case row%2 == 0:
					style = EvenRowStyle // 偶数行
				default:
					style = OddRowStyle // 奇数行
				}

				// 设置第一列格式
				if col == 0 {
					style = style.Foreground(general.ColumnOneColor)
				}

				return style
			})

			table.Headers(tableHeader...) // 设置表头
			table.Rows(tableData...)      // 设置单元格

			color.Println(table)
		}
	}
}
