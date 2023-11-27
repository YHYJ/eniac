/*
File: get.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 10:53:10

Description: 由程序子命令 get 执行
*/

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"github.com/yhyj/eniac/cli"
	"github.com/yhyj/eniac/general"
	"github.com/zcalusic/sysinfo"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get system information",
	Long:  `Get system information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置文件
		cfgFile, _ := cmd.Flags().GetString("config")
		confTree, err := cli.GetTomlConfig(cfgFile)
		if err != nil {
			fmt.Printf(general.Info2PFormat, err, ", use default configuration")
		}

		// 设置配置项默认值
		defaultMainCfg, _ := toml.TreeFromMap(map[string]interface{}{"main": map[string]string{}})
		defaultGenealogyCfg, _ := toml.TreeFromMap(map[string]interface{}{"genealogy": map[string]string{}})
		var (
			color             bool       = true
			cpuCacheUnit      string     = "KB"
			memoryDataUnit    string     = "GB"
			memoryPercentUnit string     = "%"
			updateRecordFile  string     = "/tmp/system-checkupdates.log"
			genealogyCfg      *toml.Tree = defaultGenealogyCfg
			mainCfg           *toml.Tree = defaultMainCfg
		)

		// 获取配置项
		if confTree != nil {
			genealogyCfg = func() *toml.Tree {
				if confTree.Has("genealogy") {
					return confTree.Get("genealogy").(*toml.Tree)
				}
				fmt.Printf(general.InfoFormat, "Config file is missing 'genealogy' configuration item, using default value")
				return defaultGenealogyCfg
			}()
			mainCfg = func() *toml.Tree {
				if confTree.Has("main") {
					return confTree.Get("main").(*toml.Tree)
				}
				fmt.Printf(general.InfoFormat, "Config file is missing 'main' configuration item, using default value")
				return defaultMainCfg
			}()
			color = func() bool {
				if mainCfg.Has("color") {
					return mainCfg.Get("color").(bool)
				}
				fmt.Printf(general.InfoFormat, "Config file is missing 'main.color' configuration item, using default value")
				return true
			}()
		} else {
			fmt.Printf(general.InfoFormat, "Config file is empty, using default value")
		}

		// 采集系统信息（集中采集一次后分配到不同的参数）
		var sysInfo sysinfo.SysInfo
		sysInfo.GetSysInfo()

		// 解析参数
		var biosFlag, boardFlag, cpuFlag, gpuFlag, loadFlag, memoryFlag, osFlag, productFlag, storageFlag, swapFlag, nicFlag, timeFlag, userFlag, updateFlag, onlyFlag bool
		allFlag, _ := cmd.Flags().GetBool("all")
		if allFlag {
			biosFlag, boardFlag, gpuFlag, cpuFlag, loadFlag, memoryFlag, osFlag, productFlag, storageFlag, swapFlag, nicFlag, timeFlag, userFlag, updateFlag = true, true, true, true, true, true, true, true, true, true, true, true, true, true
			onlyFlag = false
		} else {
			biosFlag, _ = cmd.Flags().GetBool("bios")
			boardFlag, _ = cmd.Flags().GetBool("board")
			gpuFlag, _ = cmd.Flags().GetBool("gpu")
			cpuFlag, _ = cmd.Flags().GetBool("cpu")
			loadFlag, _ = cmd.Flags().GetBool("load")
			memoryFlag, _ = cmd.Flags().GetBool("memory")
			osFlag, _ = cmd.Flags().GetBool("os")
			productFlag, _ = cmd.Flags().GetBool("product")
			storageFlag, _ = cmd.Flags().GetBool("storage")
			swapFlag, _ = cmd.Flags().GetBool("swap")
			nicFlag, _ = cmd.Flags().GetBool("nic")
			timeFlag, _ = cmd.Flags().GetBool("time")
			userFlag, _ = cmd.Flags().GetBool("user")
			updateFlag, _ = cmd.Flags().GetBool("update")
			onlyFlag, _ = cmd.Flags().GetBool("only")
		}

		var (
			items       []string                   // 输出项名称参数
			columnColor = tablewriter.FgWhiteColor // 默认列颜色
			headerColor = tablewriter.FgWhiteColor // 默认表头颜色
		)

		// 执行对应函数
		if productFlag {
			productPart := func() string {
				if mainCfg.Has("parts.Product") {
					return mainCfg.Get("parts.Product").(string)
				}
				return "Product"
			}()
			fmt.Printf(general.LineShownFormat, "······ Product ······")

			// 获取数据
			productInfo := cli.GetProductInfo(sysInfo)
			items = []string{"ProductVendor", "ProductName"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
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
			if color {
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

		if boardFlag {
			boardPart := func() string {
				if mainCfg.Has("parts.Board") {
					return mainCfg.Get("parts.Board").(string)
				}
				return "Board"
			}()
			fmt.Printf(general.LineShownFormat, "······ Board ······")

			// 获取数据
			boardInfo := cli.GetBoardInfo(sysInfo)
			items = []string{"BoardVendor", "BoardName", "BoardVersion"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
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
			if color {
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

		if biosFlag {
			biosPart := func() string {
				if mainCfg.Has("parts.BIOS") {
					return mainCfg.Get("parts.BIOS").(string)
				}
				return "BIOS"
			}()
			fmt.Printf(general.LineShownFormat, "····· BIOS ······")

			// 获取数据
			biosInfo := cli.GetBIOSInfo(sysInfo)
			items = []string{"BIOSVendor", "BIOSVersion", "BIOSDate"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
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
			if color {
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

		if cpuFlag {
			cpuPart := func() string {
				if mainCfg.Has("parts.CPU") {
					return mainCfg.Get("parts.CPU").(string)
				}
				return "CPU"
			}()
			fmt.Printf(general.LineShownFormat, "······ CPU ······")

			// 获取 CPU 配置项
			if genealogyCfg.Has("cpu.cache_unit") {
				cpuCacheUnit = genealogyCfg.Get("cpu.cache_unit").(string)
			} else {
				fmt.Printf(general.InfoFormat, "Config file is missing 'cpu.cache_unit' configuration item, using default value")
			}

			// 获取数据
			cpuInfo := cli.GetCPUInfo(sysInfo, cpuCacheUnit)
			items = []string{"CPUModel", "CPUNumber", "CPUCores", "CPUThreads", "CPUCache"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
				tableHeader = append(tableHeader, item)
			}

			// 组装表数据
			tableData := [][]string{}
			outputInfo := []string{cpuPart}
			for _, item := range items {
				outputValue := fmt.Sprintf("%v", cpuInfo[item])
				outputInfo = append(outputInfo, outputValue)
			}
			tableData = append(tableData, outputInfo)

			// 获取随机颜色
			if color {
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

		if gpuFlag {
			gpuPart := func() string {
				if mainCfg.Has("parts.GPU") {
					return mainCfg.Get("parts.GPU").(string)
				}
				return "GPU"
			}()
			fmt.Printf(general.LineShownFormat, "······ GPU ······")

			// 获取数据
			gpuInfo := cli.GetGPUInfo()
			items = []string{"GPUAddress", "GPUDriver", "GPUProduct", "GPUVendor"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
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
			if color {
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

		if memoryFlag {
			memoryPart := func() string {
				if mainCfg.Has("parts.Memory") {
					return mainCfg.Get("parts.Memory").(string)
				}
				return "Memory"
			}()
			fmt.Printf(general.LineShownFormat, "······ Memory ······")

			// 获取 Memory 配置项
			if genealogyCfg.Has("memory.data_unit") {
				memoryDataUnit = genealogyCfg.Get("memory.data_unit").(string)
			} else {
				fmt.Printf(general.InfoFormat, "Config file is missing 'memory.data_unit' configuration item, using default value")
			}
			if genealogyCfg.Has("memory.percent_unit") {
				memoryPercentUnit = genealogyCfg.Get("memory.percent_unit").(string)
			} else {
				fmt.Printf(general.InfoFormat, "Config file is missing 'memory.percent_unit' configuration item, using default value")
			}

			// 获取数据
			memoryInfo := cli.GetMemoryInfo(memoryDataUnit, memoryPercentUnit)
			items = []string{"MemoryUsedPercent", "MemoryTotal", "MemoryUsed", "MemoryAvail", "MemoryFree", "MemoryBuffCache", "MemoryShared"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
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
			if color {
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

		if swapFlag {
			swapPart := func() string {
				if mainCfg.Has("parts.Swap") {
					return mainCfg.Get("parts.Swap").(string)
				}
				return "Swap"
			}()
			fmt.Printf(general.LineShownFormat, "······ Swap ······")

			// 获取 Memory 配置项
			if genealogyCfg.Has("memory.data_unit") {
				memoryDataUnit = genealogyCfg.Get("memory.data_unit").(string)
			} else {
				fmt.Printf(general.InfoFormat, "Config file is missing 'memory.data_unit' configuration item, using default value")
			}

			// 获取数据
			swapInfo := cli.GetSwapInfo(memoryDataUnit)

			table := tablewriter.NewWriter(os.Stdout) // 初始化表格

			// 获取随机颜色
			if color {
				columnColor = general.GetColor()
				headerColor = tablewriter.FgCyanColor
			}

			// 组装表头
			tableHeader := []string{""}
			if swapInfo["SwapDisabled"] == true {
				items = []string{"SwapDisabled"}
				for _, item := range items {
					if genealogyCfg.Has(item) {
						item = genealogyCfg.Get(item).(string)
					}
					tableHeader = append(tableHeader, item)
				}
				table.SetHeader(tableHeader) // 设置表头
				table.SetHeaderColor(        // 设置表头颜色
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
					if genealogyCfg.Has(item) {
						item = genealogyCfg.Get(item).(string)
					}
					tableHeader = append(tableHeader, item)
				}
				table.SetHeader(tableHeader) // 设置表头
				table.SetHeaderColor(        // 设置表头颜色
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
				outputValue := fmt.Sprintf("%v", swapInfo[item])
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

		if storageFlag {
			fmt.Printf(general.LineShownFormat, "······ Storage ······")

			// 获取数据
			storageInfo := cli.GetStorageInfo()
			items = []string{"StorageName", "StorageSize", "StorageType", "StorageDriver", "StorageVendor", "StorageModel", "StorageSerial", "StorageRemovable"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
				tableHeader = append(tableHeader, item)
			}

			// 组装表数据
			tableData := [][]string{}
			diskPart := func() string {
				if mainCfg.Has("parts.Disk") {
					return mainCfg.Get("parts.Disk").(string)
				}
				return "Disk"
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
			if color {
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

		if nicFlag {
			nicPart := func() string {
				if mainCfg.Has("parts.NIC") {
					return mainCfg.Get("parts.NIC").(string)
				}
				return "NIC"
			}()
			fmt.Printf(general.LineShownFormat, "······ Nic ······")

			// 获取数据
			nicInfo := cli.GetNicInfo()
			items = []string{"NicName", "NicMacAddress", "NicDriver", "NicVendor", "NicProduct", "NicPCIAddress", "NicSpeed", "NicDuplex"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
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
			if color {
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

		if osFlag {
			osPart := func() string {
				if mainCfg.Has("parts.OS") {
					return mainCfg.Get("parts.OS").(string)
				}
				return "OS"
			}()
			fmt.Printf(general.LineShownFormat, "······ OS ······")

			// 获取数据
			osInfo := cli.GetOSInfo(sysInfo)
			items = []string{"OS", "Kernel", "Platform", "Arch", "TimeZone", "Hostname"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
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
			if color {
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

		if loadFlag {
			loadPart := func() string {
				if mainCfg.Has("parts.Load") {
					return mainCfg.Get("parts.Load").(string)
				}
				return "Load"
			}()
			fmt.Printf(general.LineShownFormat, "······ Load ······")

			// 获取数据
			loadInfo := cli.GetLoadInfo()
			items = []string{"Load1", "Load5", "Load15", "Process"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
				tableHeader = append(tableHeader, item)
			}

			// 组装表数据
			tableData := [][]string{}
			outputInfo := []string{loadPart}
			outputValue := ""
			for _, item := range items {
				if item == "Process" {
					outputValue = fmt.Sprintf("%d", loadInfo[item].(uint64))
				} else {
					outputValue = fmt.Sprintf("%.2f", loadInfo[item].(float64))
				}
				outputInfo = append(outputInfo, outputValue)
			}
			tableData = append(tableData, outputInfo)

			// 获取随机颜色
			if color {
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

		if timeFlag {
			timePart := func() string {
				if mainCfg.Has("parts.Time") {
					return mainCfg.Get("parts.Time").(string)
				}
				return "Time"
			}()
			fmt.Printf(general.LineShownFormat, "······ Time ······")

			// 获取数据
			timeInfo, _ := cli.GetTimeInfo()
			items = []string{"StartTime", "Uptime", "BootTime"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
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
			if color {
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

		if userFlag {
			userPart := func() string {
				if mainCfg.Has("parts.User") {
					return mainCfg.Get("parts.User").(string)
				}
				return "User"
			}()
			fmt.Printf(general.LineShownFormat, "······ User ······")

			// 获取数据
			userInfo := cli.GetUserInfo()
			items = []string{"UserName", "User", "UserUid", "UserGid", "UserHomeDir"}

			// 组装表头
			tableHeader := []string{""}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					item = genealogyCfg.Get(item).(string)
				}
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
			if color {
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

		if updateFlag {
			if onlyFlag {
				// 仅输出可更新包信息，专为系统更新检测插件服务
				updateInfo, err := cli.GetUpdateInfo(updateRecordFile, 0)
				if err != nil {
					fmt.Printf(general.ErrorBaseFormat, err)
				} else {
					for num, info := range updateInfo {
						fmt.Printf("%v: %v\n", num+1, info)
					}
				}
			} else {
				fmt.Printf(general.LineShownFormat, "······ Update ······")
				// 获取 update 配置项
				if genealogyCfg.Has("update.record_file") {
					updateRecordFile = genealogyCfg.Get("update.record_file").(string)
				} else {
					fmt.Printf(general.InfoFormat, "Config file is missing 'update.record_file' configuration item, using default value")
				}
				textFormat := "\x1b[37;1m%v:\x1b[0m \x1b[37;1m%v\x1b[0m\n"
				listFormat := "%8v: \x1b[37m%v\x1b[0m\n"
				if color {
					textFormat = "\x1b[30;1m%v:\x1b[0m \x1b[32;1m%v\x1b[0m\n"
					listFormat = "%8v: \x1b[32m%v\x1b[0m\n"
				}
				// 输出更新状态监测
				daemonInfo, _ := cli.GetUpdateDaemonInfo()
				items = []string{"DaemonStatus"}
				for _, item := range items {
					if genealogyCfg.Has(item) {
						fmt.Printf(textFormat, genealogyCfg.Get(item).(string), daemonInfo[item])
					} else {
						fmt.Printf(textFormat, item, daemonInfo[item])
					}
				}
				// 输出具体更新信息
				updateInfo, err := cli.GetUpdateInfo(updateRecordFile, 0)
				if err != nil {
					fmt.Printf(general.ErrorBaseFormat, err)
				} else {
					key := "UpdateList"
					if genealogyCfg.Has(key) {
						fmt.Printf(textFormat, genealogyCfg.Get(key).(string), len(updateInfo))
					} else {
						fmt.Printf(textFormat, key, len(updateInfo))
					}
					for num, info := range updateInfo {
						fmt.Printf(listFormat, num+1, info)
					}
				}
			}
		}
	},
}

func init() {
	getCmd.Flags().BoolP("all", "", false, "Get all information")
	getCmd.Flags().BoolP("bios", "", false, "Get BIOS information")
	getCmd.Flags().BoolP("board", "", false, "Get Board information")
	getCmd.Flags().BoolP("cpu", "", false, "Get CPU information")
	getCmd.Flags().BoolP("gpu", "", false, "Get GPU information")
	getCmd.Flags().BoolP("load", "", false, "Get Load information")
	getCmd.Flags().BoolP("memory", "", false, "Get Memory information")
	getCmd.Flags().BoolP("os", "", false, "Get OS information")
	getCmd.Flags().BoolP("product", "", false, "Get Product information")
	getCmd.Flags().BoolP("storage", "", false, "Get Storage information")
	getCmd.Flags().BoolP("swap", "", false, "Get Swap information")
	getCmd.Flags().BoolP("nic", "", false, "Get NIC information")
	getCmd.Flags().BoolP("time", "", false, "Get Time information")
	getCmd.Flags().BoolP("user", "", false, "Get User information")
	getCmd.Flags().BoolP("update", "", false, "Get Update information")
	getCmd.Flags().BoolP("only", "", false, "Get update package information only")

	getCmd.Flags().BoolP("help", "h", false, "help for get command")
	rootCmd.AddCommand(getCmd)
}
