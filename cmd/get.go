/*
File: get.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 10:53:10

Description: 程序子命令'get'时执行
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
		defaultGenealogyCfg, _ := toml.TreeFromMap(map[string]interface{}{"genealogy": map[string]string{}})
		defaultPartsCfg, _ := toml.TreeFromMap(map[string]interface{}{"parts": map[string]string{}})
		var (
			cpuCacheUnit      string     = "KB"
			memoryDataUnit    string     = "GB"
			memoryPercentUnit string     = "%"
			updateRecordFile  string     = "/tmp/system-checkupdates.log"
			genealogyCfg      *toml.Tree = defaultGenealogyCfg
			partsCfg          *toml.Tree = defaultPartsCfg
		)

		// 获取genealogy配置项
		if confTree != nil {
			genealogyCfg = func() *toml.Tree {
				if confTree.Has("genealogy") {
					return confTree.Get("genealogy").(*toml.Tree)
				}
				fmt.Printf(general.InfoFormat, "Config file is missing 'genealogy' configuration item, using default value")
				return defaultGenealogyCfg
			}()
			partsCfg = func() *toml.Tree {
				if confTree.Has("parts") {
					return confTree.Get("parts").(*toml.Tree)
				}
				fmt.Printf(general.InfoFormat, "Config file is missing 'parts' configuration item, using default value")
				return defaultPartsCfg
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

		var items []string // 输出项名称参数

		// 执行对应函数
		if productFlag {
			productPart := func() string {
				if partsCfg.Has("Product") {
					return partsCfg.Get("Product").(string)
				}
				return "Product"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ Product ······")

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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("Board") {
					return partsCfg.Get("Board").(string)
				}
				return "Board"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ Board ······")

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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("BIOS") {
					return partsCfg.Get("BIOS").(string)
				}
				return "BIOS"
			}()
			fmt.Printf(general.LineHiddenFormat, "····· BIOS ······")

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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("CPU") {
					return partsCfg.Get("CPU").(string)
				}
				return "CPU"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ CPU ······")

			// 获取CPU配置项
			if confTree != nil {
				if confTree.Has("cpu.cache_unit") {
					cpuCacheUnit = confTree.Get("cpu.cache_unit").(string)
				} else {
					fmt.Printf(general.InfoFormat, "Config file is missing 'cpu.cache_unit' configuration item, using default value")
				}
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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("GPU") {
					return partsCfg.Get("GPU").(string)
				}
				return "GPU"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ GPU ······")

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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("Memory") {
					return partsCfg.Get("Memory").(string)
				}
				return "Memory"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ Memory ······")

			// 获取Memory配置项
			if confTree != nil {
				if confTree.Has("memory.data_unit") {
					memoryDataUnit = confTree.Get("memory.data_unit").(string)
				} else {
					fmt.Printf(general.InfoFormat, "Config file is missing 'memory.data_unit' configuration item, using default value")
				}
				if confTree.Has("memory.percent_unit") {
					memoryPercentUnit = confTree.Get("memory.percent_unit").(string)
				} else {
					fmt.Printf(general.InfoFormat, "Config file is missing 'memory.percent_unit' configuration item, using default value")
				}
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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("Swap") {
					return partsCfg.Get("Swap").(string)
				}
				return "Swap"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ Swap ······")

			// 获取Memory配置项
			if confTree != nil {
				if confTree.Has("memory.data_unit") {
					memoryDataUnit = confTree.Get("memory.data_unit").(string)
				} else {
					fmt.Printf(general.InfoFormat, "Config file is missing 'memory.data_unit' configuration item, using default value")
				}
			}

			// 获取数据
			swapInfo := cli.GetSwapInfo(memoryDataUnit)

			table := tablewriter.NewWriter(os.Stdout) // 初始化表格

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
					tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				)
				table.SetColumnColor( // 设置列颜色
					tablewriter.Colors{tablewriter.FgHiBlackColor},
					tablewriter.Colors{tablewriter.FgBlueColor},
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
					tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
					tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				)
				table.SetColumnColor( // 设置列颜色
					tablewriter.Colors{tablewriter.FgHiBlackColor},
					tablewriter.Colors{tablewriter.FgBlueColor},
					tablewriter.Colors{tablewriter.FgBlueColor},
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
			fmt.Printf(general.LineHiddenFormat, "······ Storage ······")

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
				if partsCfg.Has("Disk") {
					return partsCfg.Get("Disk").(string)
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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_LEFT)                                             // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("NIC") {
					return partsCfg.Get("NIC").(string)
				}
				return "NIC"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ Nic ······")

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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_LEFT)                                             // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("OS") {
					return partsCfg.Get("OS").(string)
				}
				return "OS"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ OS ······")

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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("Load") {
					return partsCfg.Get("Load").(string)
				}
				return "Load"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ Load ······")

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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("Time") {
					return partsCfg.Get("Time").(string)
				}
				return "Time"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ Time ······")

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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				if partsCfg.Has("User") {
					return partsCfg.Get("User").(string)
				}
				return "User"
			}()
			fmt.Printf(general.LineHiddenFormat, "······ User ······")

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

			table := tablewriter.NewWriter(os.Stdout)                                              // 初始化表格
			table.SetAlignment(tablewriter.ALIGN_CENTER)                                           // 设置对齐方式
			table.SetBorders(tablewriter.Border{Top: true, Bottom: true, Left: true, Right: true}) // 设置表格边框
			table.SetCenterSeparator("·")                                                          // 设置中间分隔符
			table.SetAutoWrapText(false)                                                           // 设置是否自动换行
			table.SetRowLine(false)                                                                // 设置是否显示行边框
			table.SetHeader(tableHeader)                                                           // 设置表头
			table.SetHeaderColor(                                                                  // 设置表头颜色
				tablewriter.Colors{tablewriter.BgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			)
			table.SetColumnColor( // 设置列颜色
				tablewriter.Colors{tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
				tablewriter.Colors{tablewriter.FgBlueColor},
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
				fmt.Printf(general.LineHiddenFormat, "······ Update ······")
				// 获取update配置项
				if confTree != nil {
					if confTree.Has("update.record_file") {
						updateRecordFile = confTree.Get("update.record_file").(string)
					} else {
						fmt.Printf(general.InfoFormat, "Config file is missing 'update.record_file' configuration item, using default value")
					}
				}
				textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[32;1m%v\x1b[0m\n"
				listFormat := "%8v: \x1b[32m%v\x1b[0m\n"
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
