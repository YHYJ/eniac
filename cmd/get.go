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
			fmt.Printf("\x1b[36;1m%s\x1b[0m\n", err)
		} else {
			// 设置配置项默认值
			exampleGenealogy := map[string]interface{}{
				"genealogy": map[string]string{},
			}
			exampleGenealogyObject, _ := toml.TreeFromMap(exampleGenealogy)
			var (
				cpuCacheUnit      string     = "KB"
				memoryDataUnit    string     = "GB"
				memoryPercentUnit string     = "%"
				genealogyCfg      *toml.Tree = exampleGenealogyObject
				updateRecordFile  string     = "/tmp/system-checkupdates.log"
			)
			// 获取genealogy配置项
			if confTree.Has("genealogy") {
				genealogyCfg = confTree.Get("genealogy").(*toml.Tree)
			} else {
				fmt.Printf("\x1b[34;1m%s\x1b[0m\n", "config file is missing 'genealogy' configuration item, using default value")
			}
			// 采集系统信息（集中采集一次后分配到不同的参数）
			var sysInfo sysinfo.SysInfo
			sysInfo.GetSysInfo()
			// 解析参数
			var biosFlag, boardFlag, cpuFlag, loadFlag, memoryFlag, osFlag, processFlag, productFlag, storageFlag, swapFlag, timeFlag, userFlag, updateFlag bool
			allFlag, _ := cmd.Flags().GetBool("all")
			if allFlag {
				biosFlag = true
				boardFlag = true
				cpuFlag = true
				loadFlag = true
				memoryFlag = true
				osFlag = true
				processFlag = true
				productFlag = true
				storageFlag = true
				swapFlag = true
				timeFlag = true
				userFlag = true
				updateFlag = true
			} else {
				biosFlag, _ = cmd.Flags().GetBool("bios")
				boardFlag, _ = cmd.Flags().GetBool("board")
				cpuFlag, _ = cmd.Flags().GetBool("cpu")
				loadFlag, _ = cmd.Flags().GetBool("load")
				memoryFlag, _ = cmd.Flags().GetBool("memory")
				osFlag, _ = cmd.Flags().GetBool("os")
				processFlag, _ = cmd.Flags().GetBool("process")
				productFlag, _ = cmd.Flags().GetBool("product")
				storageFlag, _ = cmd.Flags().GetBool("storage")
				swapFlag, _ = cmd.Flags().GetBool("swap")
				timeFlag, _ = cmd.Flags().GetBool("time")
				userFlag, _ = cmd.Flags().GetBool("user")
				updateFlag, _ = cmd.Flags().GetBool("update")
			}
			// 执行对应函数
			if productFlag {
				fmt.Println("----------Product Information----------")
				productInfo, _ := function.GetProductInfo(sysInfo)
				// 顺序输出
				var slice = []string{"ProductVendor", "ProductName"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), productInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, productInfo[key])
					}
				}
			}
			if boardFlag {
				fmt.Println("----------Board Information----------")
				boardInfo, _ := function.GetBoardInfo(sysInfo)
				// 顺序输出
				var slice = []string{"BoardVendor", "BoardName", "BoardVersion"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), boardInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, boardInfo[key])
					}
				}
			}
			if biosFlag {
				fmt.Println("----------BIOS Information----------")
				biosInfo, _ := function.GetBIOSInfo(sysInfo)
				// 顺序输出
				var slice = []string{"BIOSVendor", "BIOSVersion", "BIOSDate"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), biosInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, biosInfo[key])
					}
				}
			}
			if cpuFlag {
				fmt.Println("----------CPU Information----------")
				// 获取CPU配置项
				if confTree.Has("cpu.cache_unit") {
					cpuCacheUnit = confTree.Get("cpu.cache_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1m%s\x1b[0m\n", "config file is missing 'cpu.cache_unit' item, using default value")
				}
				cpuInfo, _ := function.GetCPUInfo(sysInfo, cpuCacheUnit)
				// 顺序输出
				var slice = []string{"CPUModel", "CPUCache", "CPUNumber", "CPUCores", "CPUThreads"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), cpuInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, cpuInfo[key])
					}
				}
			}
			if storageFlag {
				fmt.Println("----------Storage Information----------")
				storageInfo, _ := function.GetStorageInfo(sysInfo)
				for index, values := range storageInfo {
					// 顺序输出
					var slice = []string{"StorageName", "StorageSize", "StorageDriver", "StorageVendor", "StorageModel", "StorageSerial"}
					fmt.Println(index)
					for _, name := range slice {
						if genealogyCfg.Has(name) {
							fmt.Printf("%4v%v: %v\n", "", genealogyCfg.Get(name).(string), values.(map[string]interface{})[name])
						} else {
							fmt.Printf("%4v%v: %v\n", "", name, values.(map[string]interface{})[name])
						}
					}
				}
			}
			if memoryFlag {
				fmt.Println("----------Memory Information----------")
				// 获取Memory配置项
				if confTree.Has("memory.data_unit") {
					memoryDataUnit = confTree.Get("memory.data_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1m%s\x1b[0m\n", "config file is missing 'memory.data_unit' item, using default value")
				}
				if confTree.Has("memory.percent_unit") {
					memoryPercentUnit = confTree.Get("memory.percent_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1m%s\x1b[0m\n", "config file is missing 'memory.percent_unit' item, using default value")
				}
				memInfo, _ := function.GetMemoryInfo(memoryDataUnit, memoryPercentUnit)
				// 顺序输出
				var slice = []string{"MemoryTotal", "MemoryUsed", "MemoryUsedPercent", "MemoryFree", "MemoryShared", "MemoryBuffCache", "MemoryAvail"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), memInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, memInfo[key])
					}
				}
			}
			if swapFlag {
				fmt.Println("----------Swap Information----------")
				// 获取Memory配置项
				if confTree.Has("memory.data_unit") {
					memoryDataUnit = confTree.Get("memory.data_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1m%s\x1b[0m\n", "config file is missing 'memory.data_unit' item, using default value")
				}
				swapInfo, _ := function.GetSwapInfo(memoryDataUnit)
				// 顺序输出
				if swapInfo["SwapDisabled"] == true {
					var slice = []string{"SwapDisabled"}
					for _, key := range slice {
						if genealogyCfg.Has(key) {
							fmt.Printf("%v\n", genealogyCfg.Get(key).(string))
						} else {
							fmt.Printf("%v\n", key)
						}
					}
				} else {
					var slice = []string{"SwapTotal", "SwapFree"}
					for _, key := range slice {
						if genealogyCfg.Has(key) {
							fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), swapInfo[key])
						} else {
							fmt.Printf("%v: %v\n", key, swapInfo[key])
						}
					}
				}
			}
			if osFlag {
				fmt.Println("----------OS Information----------")
				osInfo, _ := function.GetOSInfo(sysInfo)
				// 顺序输出
				var slice = []string{"Platform", "OS", "Kernel", "Arch", "Hostname", "TimeZone"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), osInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, osInfo[key])
					}
				}
			}
			if loadFlag {
				fmt.Println("----------Load Information----------")
				loadInfo, _ := function.GetLoadInfo()
				// 顺序输出
				var slice = []string{"Load1", "Load5", "Load15"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), loadInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, loadInfo[key])
					}
				}
			}
			if processFlag {
				fmt.Println("----------Process Information----------")
				procsInfo, _ := function.GetProcessInfo()
				// 顺序输出
				var slice = []string{"Process"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), procsInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, procsInfo[key])
					}
				}
			}
			if timeFlag {
				fmt.Println("----------Time Information----------")
				timeInfo, _ := function.GetTimeInfo()
				// 顺序输出
				var slice = []string{"Uptime", "BootTime", "StartTime"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), timeInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, timeInfo[key])
					}
				}
			}
			if userFlag {
				fmt.Println("----------User Information----------")
				userInfo, _ := function.GetUserInfo()
				// 顺序输出
				var slice = []string{"User", "UserName", "UserUid", "UserGid", "UserHomeDir"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), userInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, userInfo[key])
					}
				}
			}
			if updateFlag {
				fmt.Println("----------Update Information----------")
				// 获取update配置项
				if confTree.Has("update.record_file") {
					updateRecordFile = confTree.Get("update.record_file").(string)
				} else {
					fmt.Printf("\x1b[34;1m%s\x1b[0m\n", "config file is missing 'update.record_file' item, using default value")
				}
				updateInfo, err := function.GetUpdateInfo(updateRecordFile, 0)
				if err != nil {
					fmt.Printf("\x1b[36;1m%s\x1b[0m\n", err)
				} else {
					for num, info := range updateInfo {
						fmt.Println(num+1, info)
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
	getCmd.Flags().BoolP("load", "", false, "Get Load information")
	getCmd.Flags().BoolP("memory", "", false, "Get Memory information")
	getCmd.Flags().BoolP("os", "", false, "Get OS information")
	getCmd.Flags().BoolP("process", "", false, "Get Process information")
	getCmd.Flags().BoolP("product", "", false, "Get Product information")
	getCmd.Flags().BoolP("storage", "", false, "Get Storage information")
	getCmd.Flags().BoolP("swap", "", false, "Get Swap information")
	getCmd.Flags().BoolP("time", "", false, "Get Time information")
	getCmd.Flags().BoolP("user", "", false, "Get User information")
	getCmd.Flags().BoolP("update", "", false, "Get Update information")

	getCmd.Flags().BoolP("help", "h", false, "help for get")
	rootCmd.AddCommand(getCmd)
}
