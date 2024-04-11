/*
File: get.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-11 10:37:20

Description: 子命令 'get' 的实现
*/

package cli

import (
	"os"
	"strconv"

	"github.com/gookit/color"
	"github.com/olekukonko/tablewriter"
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
	config, err := LoadConfigToStruct(configTree)
	if err != nil {
		color.Error.Println(err)
		return
	}

	// 设置配置项默认值
	var (
		colorful          bool   = true
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
		items       []string                   // 输出项名称参数
		columnColor = tablewriter.FgWhiteColor // 默认列颜色
		headerColor = tablewriter.FgWhiteColor // 默认表头颜色
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
		color.Printf("%s\n", general.FgGrayText("······ Product ······"))

		// 获取数据
		productInfo := GetProductInfo(sysInfo)
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["boardFlag"] {
		boardPart := func() string {
			partName := general.PartName["Board"][general.Language]
			if partName == "" {
				partName = "Board"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ Board ······"))

		// 获取数据
		boardInfo := GetBoardInfo(sysInfo)
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["biosFlag"] {
		biosPart := func() string {
			partName := general.PartName["BIOS"][general.Language]
			if partName == "" {
				partName = "BIOS"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ BIOS ······"))

		// 获取数据
		biosInfo := GetBIOSInfo(sysInfo)
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["cpuFlag"] {
		cpuPart := func() string {
			partName := general.PartName["CPU"][general.Language]
			if partName == "" {
				partName = "CPU"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ CPU ······"))

		// 获取 CPU 配置项
		if config.Genealogy.Cpu.CacheUnit != "" {
			cpuCacheUnit = config.Genealogy.Cpu.CacheUnit
		} else {
			color.Error.Println("Config file is missing 'cpu.cache_unit' item, using default value")
		}

		// 获取数据
		cpuInfo := GetCPUInfo(sysInfo, cpuCacheUnit)
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["gpuFlag"] {
		gpuPart := func() string {
			partName := general.PartName["GPU"][general.Language]
			if partName == "" {
				partName = "GPU"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ GPU ······"))

		// 获取数据
		gpuInfo := GetGPUInfo()
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["memoryFlag"] {
		memoryPart := func() string {
			partName := general.PartName["Memory"][general.Language]
			if partName == "" {
				partName = "Memory"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ Memory ······"))

		// 获取 Memory 配置项
		if config.Genealogy.Memory.DataUnit != "" {
			memoryDataUnit = config.Genealogy.Memory.DataUnit
		} else {
			color.Error.Println("Config file is missing 'memory.data_unit' item, using default value")
		}
		if config.Genealogy.Memory.PercentUnit != "" {
			memoryPercentUnit = config.Genealogy.Memory.PercentUnit
		} else {
			color.Error.Println("Config file is missing 'memory.percent_unit' item, using default value")
		}

		// 获取数据
		memoryInfo := GetMemoryInfo(memoryDataUnit, memoryPercentUnit)
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["swapFlag"] {
		swapPart := func() string {
			partName := general.PartName["Swap"][general.Language]
			if partName == "" {
				partName = "Swap"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ Swap ······"))

		// 获取 Memory 配置项
		if config.Genealogy.Memory.DataUnit != "" {
			memoryDataUnit = config.Genealogy.Memory.DataUnit
		} else {
			color.Error.Println("Config file is missing 'memory.data_unit' item, using default value")
		}

		// 获取数据
		swapInfo := GetSwapInfo(memoryDataUnit)

		table := tablewriter.NewWriter(os.Stdout) // 初始化表格

		// 获取随机颜色
		if colorful {
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		// 组装表头
		tableHeader := []string{""}
		if swapInfo["SwapDisabled"] == true {
			items = []string{"SwapDisabled"}
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
			table.SetHeader(tableHeader)      // 设置表头
			table.SetAutoFormatHeaders(false) // 设置是否自动格式化表头
			table.SetHeaderColor(             // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, headerColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{columnColor},
			)
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
			table.SetHeader(tableHeader)      // 设置表头
			table.SetAutoFormatHeaders(false) // 设置是否自动格式化表头
			table.SetHeaderColor(             // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, headerColor},
				tablewriter.Colors{tablewriter.Bold, headerColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{columnColor},
				tablewriter.Colors{columnColor},
			)
		}

		// 组装表数据
		tableData := [][]string{}
		outputInfo := []string{swapPart}
		for _, item := range items {
			outputValue := color.Sprintf("%v", swapInfo[item])
			outputInfo = append(outputInfo, outputValue)
		}
		tableData = append(tableData, outputInfo)

		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["storageFlag"] {
		color.Printf("%s\n", general.FgGrayText("······ Storage ······"))

		// 获取数据
		storageInfo := GetStorageInfo()
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
		diskPart := func() string {
			partName := general.PartName["Disk"][general.Language]
			if partName == "" {
				partName = "Disk"
			}
			return partName
		}()
		for index := 1; index <= len(storageInfo); index++ {
			outputInfo := []string{diskPart + "." + strconv.Itoa(index)}
			for _, item := range items {
				outputValue := storageInfo[strconv.Itoa(index)].(map[string]interface{})[item].(string)
				outputInfo = append(outputInfo, outputValue)
			}
			tableData = append(tableData, outputInfo)
		}

		// 获取随机颜色
		if colorful {
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_LEFT)                                             // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["nicFlag"] {
		nicPart := func() string {
			partName := general.PartName["NIC"][general.Language]
			if partName == "" {
				partName = "NIC"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ Nic ······"))

		// 获取数据
		nicInfo := GetNicInfo()
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
			outputInfo := []string{nicPart + "." + strconv.Itoa(index)}
			for _, item := range items {
				outputValue := nicInfo[strconv.Itoa(index)].(map[string]interface{})[item].(string)
				outputInfo = append(outputInfo, outputValue)
			}
			tableData = append(tableData, outputInfo)
		}

		// 获取随机颜色
		if colorful {
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_LEFT)                                             // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["osFlag"] {
		osPart := func() string {
			partNname := general.PartName["OS"][general.Language]
			if partNname == "" {
				partNname = "OS"
			}
			return partNname
		}()
		color.Printf("%s\n", general.FgGrayText("······ OS ······"))

		// 获取数据
		osInfo := GetOSInfo(sysInfo)
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["loadFlag"] {
		loadPart := func() string {
			partName := general.PartName["Load"][general.Language]
			if partName == "" {
				partName = "Load"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ Load ······"))

		// 获取数据
		loadInfo := GetLoadInfo()
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["timeFlag"] {
		timePart := func() string {
			partName := general.PartName["Time"][general.Language]
			if partName == "" {
				partName = "Time"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ Time ······"))

		// 获取数据
		timeInfo, _ := GetTimeInfo()
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["userFlag"] {
		userPart := func() string {
			partName := general.PartName["User"][general.Language]
			if partName == "" {
				partName = "User"
			}
			return partName
		}()
		color.Printf("%s\n", general.FgGrayText("······ User ······"))

		// 获取数据
		userInfo := GetUserInfo()
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
			columnColor = general.GetColor()
			headerColor = tablewriter.FgCyanColor
		}

		table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
		table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
		table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
		table.SetCenterSeparator("·")                                                          // 设置中间分隔符
		table.SetAutoWrapText(false)                                                           // 设置是否自动换行
		table.SetRowLine(false)                                                                // 设置是否显示行边框
		table.SetHeader(tableHeader)                                                           // 设置表头
		table.SetAutoFormatHeaders(false)                                                      // 设置是否自动格式化表头
		table.SetHeaderColor(                                                                  // 设置表头颜色
			tablewriter.Colors{tablewriter.BgHiBlackColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
			tablewriter.Colors{tablewriter.Bold, headerColor},
		)
		table.SetColumnColor( // 设置列颜色
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
			tablewriter.Colors{columnColor},
		)

		// 填充表格
		for _, data := range tableData {
			table.Append(data)
		}

		// 渲染表格
		table.Render()
	}

	if flags["updateFlag"] {
		updateInfo, err := GetUpdateInfo(updateRecordFile, 0)
		if err != nil {
			color.Error.Println(err)
		}
		if flags["onlyFlag"] {
			// 仅输出可更新包信息，专为第三方系统更新检测插件服务
			for num, info := range updateInfo {
				color.Println("%v: %v\n", num+1, info)
			}
		} else {
			color.Printf("%s\n", general.FgGrayText("······ Update ······"))
			// 获取 update 配置项
			if config.Genealogy.Update.RecordFile != "" {
				updateRecordFile = config.Genealogy.Update.RecordFile
			} else {
				color.Error.Println("Config file is missing 'update.record_file' item, using default value")
			}
			itemColor := general.LightText
			itemInfoColor := general.FgWhiteText
			listColor := general.FgWhiteText
			if colorful {
				itemColor = general.PrimaryText
				itemInfoColor = general.FgGreenText
				listColor = general.FgGreenText
			}
			// 更新服务状态监测
			daemonInfo, _ := GetUpdateDaemonInfo()
			daemonItem := "UpdateDaemonStatus"
			daemonItemName := general.GenealogyName[daemonItem][general.Language]
			if daemonItemName != "" {
				color.Printf("%v: %v\n", itemColor(daemonItemName), itemInfoColor(daemonInfo[daemonItem]))
			} else {
				color.Printf("%v: %v\n", itemColor(daemonItem), itemInfoColor(daemonInfo[daemonItem]))
			}
			// 更新列表计数
			packageItem := "UpdateList"
			packageItemName := general.GenealogyName[packageItem][general.Language]
			if packageItemName != "" {
				color.Printf("%v: %v\n", itemColor(packageItemName), itemInfoColor(len(updateInfo)))
			} else {
				color.Printf("%v: %v\n", itemColor(packageItem), itemInfoColor(len(updateInfo)))
			}
			// 输出可更新包信息
			for num, info := range updateInfo {
				color.Printf("%8v: %v\n", num+1, listColor(info))
			}
		}
	}
}