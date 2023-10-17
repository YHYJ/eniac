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
	"github.com/yhyj/eniac/function"
	"github.com/zcalusic/sysinfo"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get system information",
	Long:  `Get system information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置文件
		cfgFile, _ := cmd.Flags().GetString("config")
		confTree, err := function.GetTomlConfig(cfgFile)
		if err != nil {
			fmt.Printf("\x1b[36;1m%s, %s\x1b[0m\n", err, "use default configuration")
		}
		// 设置配置项默认值
		defaultGenealogyCfg, _ := toml.TreeFromMap(map[string]interface{}{"genealogy": map[string]string{}})
		var (
			cpuCacheUnit      string     = "KB"
			memoryDataUnit    string     = "GB"
			memoryPercentUnit string     = "%"
			genealogyCfg      *toml.Tree = defaultGenealogyCfg
			updateRecordFile  string     = "/tmp/system-checkupdates.log"
		)
		// 获取genealogy配置项
		if confTree != nil {
			if confTree.Has("genealogy") {
				genealogyCfg = confTree.Get("genealogy").(*toml.Tree)
			} else {
				fmt.Printf("\x1b[34;1mConfig file is missing '%s' configuration item, using default value\x1b[0m\n", "genealogy")
			}
		}
		// 采集系统信息（集中采集一次后分配到不同的参数）
		var sysInfo sysinfo.SysInfo
		sysInfo.GetSysInfo()
		// 解析参数
		var biosFlag, boardFlag, cpuFlag, gpuFlag, loadFlag, memoryFlag, osFlag, processFlag, productFlag, storageFlag, swapFlag, netFlag, timeFlag, userFlag, updateFlag, onlyFlag bool
		allFlag, _ := cmd.Flags().GetBool("all")
		if allFlag {
			biosFlag, boardFlag, gpuFlag, cpuFlag, loadFlag, memoryFlag, osFlag, processFlag, productFlag, storageFlag, swapFlag, netFlag, timeFlag, userFlag, updateFlag = true, true, true, true, true, true, true, true, true, true, true, true, true, true, true
			onlyFlag = false
		} else {
			biosFlag, _ = cmd.Flags().GetBool("bios")
			boardFlag, _ = cmd.Flags().GetBool("board")
			gpuFlag, _ = cmd.Flags().GetBool("gpu")
			cpuFlag, _ = cmd.Flags().GetBool("cpu")
			loadFlag, _ = cmd.Flags().GetBool("load")
			memoryFlag, _ = cmd.Flags().GetBool("memory")
			osFlag, _ = cmd.Flags().GetBool("os")
			processFlag, _ = cmd.Flags().GetBool("process")
			productFlag, _ = cmd.Flags().GetBool("product")
			storageFlag, _ = cmd.Flags().GetBool("storage")
			swapFlag, _ = cmd.Flags().GetBool("swap")
			netFlag, _ = cmd.Flags().GetBool("net")
			timeFlag, _ = cmd.Flags().GetBool("time")
			userFlag, _ = cmd.Flags().GetBool("user")
			updateFlag, _ = cmd.Flags().GetBool("update")
			onlyFlag, _ = cmd.Flags().GetBool("only")
		}

		var items []string // 输出项名称参数

		// 执行对应函数
		if productFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Product")
			productInfo := function.GetProductInfo(sysInfo)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[33;1m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"ProductVendor", "ProductName"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), productInfo[item])
				} else {
					fmt.Printf(textFormat, item, productInfo[item])
				}
			}
		}
		if boardFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Board")
			boardInfo := function.GetBoardInfo(sysInfo)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[33;1m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"BoardVendor", "BoardName", "BoardVersion"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), boardInfo[item])
				} else {
					fmt.Printf(textFormat, item, boardInfo[item])
				}
			}
		}
		if biosFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "BIOS")
			biosInfo := function.GetBIOSInfo(sysInfo)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[33;1m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"BIOSVendor", "BIOSVersion", "BIOSDate"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), biosInfo[item])
				} else {
					fmt.Printf(textFormat, item, biosInfo[item])
				}
			}
		}
		if cpuFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "CPU")
			// 获取CPU配置项
			if confTree != nil {
				if confTree.Has("cpu.cache_unit") {
					cpuCacheUnit = confTree.Get("cpu.cache_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1mConfig file is missing '%s' item, using default value\x1b[0m\n", "cpu.cache_unit")
				}
			}
			cpuInfo := function.GetCPUInfo(sysInfo, cpuCacheUnit)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[34;1m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"CPUModel", "CPUCache", "CPUNumber", "CPUCores", "CPUThreads"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), cpuInfo[item])
				} else {
					fmt.Printf(textFormat, item, cpuInfo[item])
				}
			}
		}
		if gpuFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "GPU")
			gpuInfo := function.GetGPUInfo()
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[34;1m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"GPUAddress", "GPUDriver", "GPUProduct", "GPUVendor"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), gpuInfo[item])
				} else {
					fmt.Printf(textFormat, item, gpuInfo[item])
				}
			}
		}
		if memoryFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Memory")
			// 获取Memory配置项
			if confTree != nil {
				if confTree.Has("memory.data_unit") {
					memoryDataUnit = confTree.Get("memory.data_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1mConfig file is missing '%s' item, using default value\x1b[0m\n", "memory.data_unit")
				}
				if confTree.Has("memory.percent_unit") {
					memoryPercentUnit = confTree.Get("memory.percent_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1mConfig file is missing '%s' item, using default value\x1b[0m\n", "memory.percent_unit")
				}
			}
			memInfo := function.GetMemoryInfo(memoryDataUnit, memoryPercentUnit)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[34;1m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"MemoryUsedPercent", "MemoryTotal", "MemoryUsed", "MemoryAvail", "MemoryFree", "MemoryBuffCache", "MemoryShared"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), memInfo[item])
				} else {
					fmt.Printf(textFormat, item, memInfo[item])
				}
			}
		}
		if swapFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Swap")
			// 获取Memory配置项
			if confTree != nil {
				if confTree.Has("memory.data_unit") {
					memoryDataUnit = confTree.Get("memory.data_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1mConfig file is missing '%s' item, using default value\x1b[0m\n", "memory.data_unit")
				}
			}
			swapInfo := function.GetSwapInfo(memoryDataUnit)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[34;1m%v\x1b[0m\n"
			// 顺序输出
			if swapInfo["SwapDisabled"] == true {
				items = []string{"SwapDisabled"}
				for _, item := range items {
					if genealogyCfg.Has(item) {
						fmt.Printf("🚫%v\n", genealogyCfg.Get(item).(string))
					} else {
						fmt.Printf("🚫%v\n", item)
					}
				}
			} else {
				items = []string{"SwapTotal", "SwapFree"}
				for _, item := range items {
					if genealogyCfg.Has(item) {
						fmt.Printf(textFormat, genealogyCfg.Get(item).(string), swapInfo[item])
					} else {
						fmt.Printf(textFormat, item, swapInfo[item])
					}
				}
			}
		}
		if storageFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Storage")
			storageInfo := function.GetStorageInfo()
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
			for index := 1; index <= len(storageInfo); index++ {
				outputInfo := []string{"磁盘." + strconv.Itoa(index)}
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
			for _, data := range tableData { // 填充表格
				table.Append(data)
			}
			table.Render() // 渲染表格
		}
		if netFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Network")
			networkInfo := function.GetNetworkInfo()
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
			for index := 1; index <= len(networkInfo); index++ {
				outputInfo := []string{"网卡." + strconv.Itoa(index)}
				for _, item := range items {
					outputValue := networkInfo[strconv.Itoa(index)].(map[string]interface{})[item].(string)
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

			for _, data := range tableData { // 填充表格
				table.Append(data)
			}
			table.Render() // 渲染表格
		}
		if osFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "OS")
			osInfo := function.GetOSInfo(sysInfo)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[35m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"Arch", "Platform", "OS", "Kernel", "TimeZone", "Hostname"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), osInfo[item])
				} else {
					fmt.Printf(textFormat, item, osInfo[item])
				}
			}
		}
		if loadFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Load")
			loadInfo := function.GetLoadInfo()
			textFormat := "\x1b[30;1m%-6v:\x1b[0m \x1b[35m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"Load1", "Load5", "Load15"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), loadInfo[item])
				} else {
					fmt.Printf(textFormat, item, loadInfo[item])
				}
			}
		}
		if processFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Process")
			procsInfo := function.GetProcessInfo()
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[35m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"Process"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), procsInfo[item])
				} else {
					fmt.Printf(textFormat, item, procsInfo[item])
				}
			}
		}
		if timeFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Time")
			timeInfo, _ := function.GetTimeInfo()
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[36m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"StartTime", "Uptime", "BootTime"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), timeInfo[item])
				} else {
					fmt.Printf(textFormat, item, timeInfo[item])
				}
			}
		}
		if userFlag {
			fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "User")
			userInfo := function.GetUserInfo()
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[36m%v\x1b[0m\n"
			// 顺序输出
			items = []string{"UserName", "User", "UserUid", "UserGid", "UserHomeDir"}
			for _, item := range items {
				if genealogyCfg.Has(item) {
					fmt.Printf(textFormat, genealogyCfg.Get(item).(string), userInfo[item])
				} else {
					fmt.Printf(textFormat, item, userInfo[item])
				}
			}
		}
		if updateFlag {
			if onlyFlag {
				// 仅输出可更新包信息，专为系统更新检测插件服务
				updateInfo, err := function.GetUpdateInfo(updateRecordFile, 0)
				if err != nil {
					fmt.Printf("\x1b[36;1m%s\x1b[0m\n", err)
				} else {
					for num, info := range updateInfo {
						fmt.Printf("%v: %v\n", num+1, info)
					}
				}
			} else {
				fmt.Printf("\x1b[37m>>>>>>>>>>\x1b[0m %s\n", "Update")
				// 获取update配置项
				if confTree != nil {
					if confTree.Has("update.record_file") {
						updateRecordFile = confTree.Get("update.record_file").(string)
					} else {
						fmt.Printf("\x1b[34;1mConfig file is missing '%s' item, using default value\x1b[0m\n", "update.record_file")
					}
				}
				textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[32;1m%v\x1b[0m\n"
				listFormat := "%8v: \x1b[32m%v\x1b[0m\n"
				// 输出更新状态监测
				daemonInfo, _ := function.GetUpdateDaemonInfo()
				items = []string{"DaemonStatus"}
				for _, item := range items {
					if genealogyCfg.Has(item) {
						fmt.Printf(textFormat, genealogyCfg.Get(item).(string), daemonInfo[item])
					} else {
						fmt.Printf(textFormat, item, daemonInfo[item])
					}
				}
				// 输出具体更新信息
				updateInfo, err := function.GetUpdateInfo(updateRecordFile, 0)
				if err != nil {
					fmt.Printf("\x1b[36;1m%s\x1b[0m\n", err)
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
	getCmd.Flags().BoolP("process", "", false, "Get Process information")
	getCmd.Flags().BoolP("product", "", false, "Get Product information")
	getCmd.Flags().BoolP("storage", "", false, "Get Storage information")
	getCmd.Flags().BoolP("swap", "", false, "Get Swap information")
	getCmd.Flags().BoolP("net", "", false, "Get Network information")
	getCmd.Flags().BoolP("time", "", false, "Get Time information")
	getCmd.Flags().BoolP("user", "", false, "Get User information")
	getCmd.Flags().BoolP("update", "", false, "Get Update information")
	getCmd.Flags().BoolP("only", "", false, "Get update package information only")

	getCmd.Flags().BoolP("help", "h", false, "help for get command")
	rootCmd.AddCommand(getCmd)
}
