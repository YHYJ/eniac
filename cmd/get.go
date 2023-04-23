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
			// 获取CPU配置项
			cpuCfg := confTree.Get("cpu").(*toml.Tree)
			cpuCacheUnit := cpuCfg.Get("cache_unit").(string)
			// 获取Memory配置项
			memoryCfg := confTree.Get("memory").(*toml.Tree)
			memoryDataUnit := memoryCfg.Get("data_unit").(string)
			memoryPercentUnit := memoryCfg.Get("percent_unit").(string)
			// 获取genealogy配置项
			genealogyCfg := confTree.Get("genealogy").(*toml.Tree)
			// 采集系统信息（集中采集一次后分配到不同的参数）
			var sysInfo sysinfo.SysInfo
			sysInfo.GetSysInfo()
			// 解析参数
			var biosFlag, boardFlag, cpuFlag, loadFlag, memoryFlag, osFlag, processFlag, productFlag, storageFlag, swapFlag, timeFlag, userFlag bool
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
			}
			// 执行对应函数
			if biosFlag {
				biosInfo, _ := function.GetBIOSInfo(sysInfo)
				// 顺序输出
				fmt.Println("----------BIOS Information----------")
				var slice = []string{"BIOSVendor", "BIOSVersion", "BIOSDate"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), biosInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, biosInfo[key])
					}
				}
			}
			if boardFlag {
				boardInfo, _ := function.GetBoardInfo(sysInfo)
				// 顺序输出
				fmt.Println("----------Board Information----------")
				var slice = []string{"BoardVendor", "BoardName", "BoardVersion"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), boardInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, boardInfo[key])
					}
				}
			}
			if cpuFlag {
				cpuInfo, _ := function.GetCPUInfo(sysInfo, cpuCacheUnit)
				// 顺序输出
				fmt.Println("----------CPU Information----------")
				var slice = []string{"CPUModel", "CPUCache", "CPUNumber", "CPUCores", "CPUThreads"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), cpuInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, cpuInfo[key])
					}
				}
			}
			if loadFlag {
				loadInfo, _ := function.GetLoadInfo()
				// 顺序输出
				fmt.Println("----------Load Information----------")
				var slice = []string{"Load1", "Load5", "Load15"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), loadInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, loadInfo[key])
					}
				}
			}
			if memoryFlag {
				memInfo, _ := function.GetMemoryInfo(memoryDataUnit, memoryPercentUnit)
				// 顺序输出
				fmt.Println("----------Memory Information----------")
				var slice = []string{"MemTotal", "MemUsed", "MemUsedPercent", "MemFree", "MemShared", "MemBuffCache", "MemAvail"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), memInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, memInfo[key])
					}
				}
			}
			if osFlag {
				osInfo, _ := function.GetOSInfo(sysInfo)
				// 顺序输出
				fmt.Println("----------OS Information----------")
				var slice = []string{"Platform", "OS", "Kernel", "Arch", "Hostname", "TimeZone"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), osInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, osInfo[key])
					}
				}
			}
			if processFlag {
				procsInfo, _ := function.GetProcessInfo()
				// 顺序输出
				fmt.Println("----------Process Information----------")
				var slice = []string{"Process"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), procsInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, procsInfo[key])
					}
				}
			}
			if productFlag {
				productInfo, _ := function.GetProductInfo(sysInfo)
				// 顺序输出
				fmt.Println("----------Product Information----------")
				var slice = []string{"ProductVendor", "ProductName"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), productInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, productInfo[key])
					}
				}
			}
			if storageFlag {
				storageInfo, _ := function.GetStorageInfo(sysInfo)
				for _, value := range storageInfo {
					// 顺序输出
					fmt.Println("----------Storage Information----------")
					var slice = []string{"StorageName", "StorageDriver", "StorageVendor", "StorageModel", "StorageSerial", "StorageSize"}
					for _, key := range slice {
						if genealogyCfg.Has(key) {
							fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), value.(map[string]interface{})[key])
						} else {
							fmt.Printf("%v: %v\n", key, value.(map[string]interface{})[key])
						}
					}
				}
			}
			if swapFlag {
				swapInfo, _ := function.GetSwapInfo(memoryDataUnit)
				// 顺序输出
				fmt.Println("----------Swap Information----------")
				var slice = []string{"SwapTotal", "SwapFree"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), swapInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, swapInfo[key])
					}
				}
			}
			if timeFlag {
				timeInfo, _ := function.GetTimeInfo()
				// 顺序输出
				fmt.Println("----------Time Information----------")
				var slice = []string{"Uptime", "BootTime"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), timeInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, timeInfo[key])
					}
				}
			}
			if userFlag {
				userInfo, _ := function.GetUserInfo()
				// 顺序输出
				fmt.Println("----------User Information----------")
				var slice = []string{"User", "UserName", "UserUid", "UserGid", "UserHomeDir"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("%v: %v\n", genealogyCfg.Get(key).(string), userInfo[key])
					} else {
						fmt.Printf("%v: %v\n", key, userInfo[key])
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

	getCmd.Flags().BoolP("help", "h", false, "help for get")
	rootCmd.AddCommand(getCmd)
}
