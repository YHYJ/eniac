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
			fmt.Printf("\x1b[36;1m%s, %s\x1b[0m\n", err, "use default configuration")
		}
		// 设置配置项默认值
		defaultGenealogyCfg, _ := toml.TreeFromMap(map[string]interface{}{"genealogy": map[string]string{}})
		var (
			cpuCacheUnit      string     = "KB"
			memoryDataUnit    string     = "GB"
			memoryPercentUnit string     = "%"
			storageAddress    string     = ""
			nicAddress        string     = ""
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
		var biosFlag, boardFlag, cpuFlag, gpuFlag, loadFlag, memoryFlag, osFlag, processFlag, productFlag, storageFlag, swapFlag, nicFlag, timeFlag, userFlag, updateFlag, onlyFlag bool
		allFlag, _ := cmd.Flags().GetBool("all")
		if allFlag {
			biosFlag = true
			boardFlag = true
			gpuFlag = true
			cpuFlag = true
			loadFlag = true
			memoryFlag = true
			osFlag = true
			processFlag = true
			productFlag = true
			storageFlag = true
			swapFlag = true
			nicFlag = true
			timeFlag = true
			userFlag = true
			updateFlag = true
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
			nicFlag, _ = cmd.Flags().GetBool("nic")
			timeFlag, _ = cmd.Flags().GetBool("time")
			userFlag, _ = cmd.Flags().GetBool("user")
			updateFlag, _ = cmd.Flags().GetBool("update")
			onlyFlag, _ = cmd.Flags().GetBool("only")
		}
		// 执行对应函数
		if productFlag {
			fmt.Println("----------Product Information----------")
			productInfo, _ := function.GetProductInfo(sysInfo)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[33;1m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"ProductVendor", "ProductName"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), productInfo[key])
				} else {
					fmt.Printf(textFormat, key, productInfo[key])
				}
			}
		}
		if boardFlag {
			fmt.Println("----------Board Information----------")
			boardInfo, _ := function.GetBoardInfo(sysInfo)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[33;1m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"BoardVendor", "BoardName", "BoardVersion"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), boardInfo[key])
				} else {
					fmt.Printf(textFormat, key, boardInfo[key])
				}
			}
		}
		if biosFlag {
			fmt.Println("----------BIOS Information----------")
			biosInfo, _ := function.GetBIOSInfo(sysInfo)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[33;1m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"BIOSVendor", "BIOSVersion", "BIOSDate"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), biosInfo[key])
				} else {
					fmt.Printf(textFormat, key, biosInfo[key])
				}
			}
		}
		if cpuFlag {
			fmt.Println("----------CPU Information----------")
			// 获取CPU配置项
			if confTree != nil {
				if confTree.Has("cpu.cache_unit") {
					cpuCacheUnit = confTree.Get("cpu.cache_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1mConfig file is missing '%s' item, using default value\x1b[0m\n", "cpu.cache_unit")
				}
			}
			cpuInfo, _ := function.GetCPUInfo(sysInfo, cpuCacheUnit)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[34m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"CPUModel", "CPUCache", "CPUNumber", "CPUCores", "CPUThreads"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), cpuInfo[key])
				} else {
					fmt.Printf(textFormat, key, cpuInfo[key])
				}
			}
		}
		if gpuFlag {
			fmt.Println("----------GPU Information----------")
			gpuInfo := function.GetGPUInfo()
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[34m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"GPUAddress", "GPUDriver", "GPUProduct", "GPUVendor"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), gpuInfo[key])
				} else {
					fmt.Printf(textFormat, key, gpuInfo[key])
				}
			}
		}
		if storageFlag {
			fmt.Println("----------Storage Information----------")
			// 获取Storage配置项
			if confTree != nil {
				if confTree.Has("storage.address") {
					storageAddress = confTree.Get("storage.address").(string)
				} else {
					fmt.Printf("\x1b[34;1mConfig file is missing '%s' item, using default value\x1b[0m\n", "storage.address")
				}
			}
			storageInfo := function.GetStorageInfo(storageAddress)
			titleFormat := "\x1b[30;1m%v\x1b[0m\n"
			textFormat := "\x1b[34;1m%4v%v:\x1b[0m \x1b[34m%v\x1b[0m\n"
			for index, values := range storageInfo {
				// 顺序输出
				var slice = []string{"StorageName", "StorageSize", "StorageType", "StorageDriver", "StorageVendor", "StorageModel", "StorageSerial", "StorageRemovable"}
				fmt.Printf(titleFormat, index)
				for _, name := range slice {
					if genealogyCfg.Has(name) {
						fmt.Printf(textFormat, "", genealogyCfg.Get(name).(string), values.(map[string]interface{})[name])
					} else {
						fmt.Printf(textFormat, "", name, values.(map[string]interface{})[name])
					}
				}
			}
		}
		if memoryFlag {
			fmt.Println("----------Memory Information----------")
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
			memInfo, _ := function.GetMemoryInfo(memoryDataUnit, memoryPercentUnit)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[34m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"MemoryUsedPercent", "MemoryTotal", "MemoryUsed", "MemoryAvail", "MemoryFree", "MemoryBuffCache", "MemoryShared"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), memInfo[key])
				} else {
					fmt.Printf(textFormat, key, memInfo[key])
				}
			}
		}
		if swapFlag {
			fmt.Println("----------Swap Information----------")
			// 获取Memory配置项
			if confTree != nil {
				if confTree.Has("memory.data_unit") {
					memoryDataUnit = confTree.Get("memory.data_unit").(string)
				} else {
					fmt.Printf("\x1b[34;1mConfig file is missing '%s' item, using default value\x1b[0m\n", "memory.data_unit")
				}
			}
			swapInfo, _ := function.GetSwapInfo(memoryDataUnit)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[34m%v\x1b[0m\n"
			// 顺序输出
			if swapInfo["SwapDisabled"] == true {
				var slice = []string{"SwapDisabled"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf("🚫%v\n", genealogyCfg.Get(key).(string))
					} else {
						fmt.Printf("🚫%v\n", key)
					}
				}
			} else {
				var slice = []string{"SwapTotal", "SwapFree"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf(textFormat, genealogyCfg.Get(key).(string), swapInfo[key])
					} else {
						fmt.Printf(textFormat, key, swapInfo[key])
					}
				}
			}
		}
		if nicFlag {
			fmt.Println("----------NIC Information----------")
			// 获取NIC配置项
			if confTree != nil {
				if confTree.Has("nic.address") {
					nicAddress = confTree.Get("nic.address").(string)
				} else {
					fmt.Printf("\x1b[34;1mConfig file is missing '%s' item, using default value\x1b[0m\n", "nic.address")
				}
			}
			nicInfo := function.GetNicInfo(nicAddress)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[34m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"NicAddress", "NicDriver", "NicVendor", "NicProduct"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), nicInfo[key])
				} else {
					fmt.Printf(textFormat, key, nicInfo[key])
				}
			}
		}
		if osFlag {
			fmt.Println("----------OS Information----------")
			osInfo, _ := function.GetOSInfo(sysInfo)
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[35m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"Arch", "Platform", "OS", "Kernel", "TimeZone", "Hostname"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), osInfo[key])
				} else {
					fmt.Printf(textFormat, key, osInfo[key])
				}
			}
		}
		if loadFlag {
			fmt.Println("----------Load Information----------")
			loadInfo, _ := function.GetLoadInfo()
			textFormat := "\x1b[30;1m%-6v:\x1b[0m \x1b[35m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"Load1", "Load5", "Load15"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), loadInfo[key])
				} else {
					fmt.Printf(textFormat, key, loadInfo[key])
				}
			}
		}
		if processFlag {
			fmt.Println("----------Process Information----------")
			procsInfo, _ := function.GetProcessInfo()
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[35m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"Process"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), procsInfo[key])
				} else {
					fmt.Printf(textFormat, key, procsInfo[key])
				}
			}
		}
		if timeFlag {
			fmt.Println("----------Time Information----------")
			timeInfo, _ := function.GetTimeInfo()
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[36m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"StartTime", "Uptime", "BootTime"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), timeInfo[key])
				} else {
					fmt.Printf(textFormat, key, timeInfo[key])
				}
			}
		}
		if userFlag {
			fmt.Println("----------User Information----------")
			userInfo, _ := function.GetUserInfo()
			textFormat := "\x1b[30;1m%v:\x1b[0m \x1b[36m%v\x1b[0m\n"
			// 顺序输出
			var slice = []string{"UserName", "User", "UserUid", "UserGid", "UserHomeDir"}
			for _, key := range slice {
				if genealogyCfg.Has(key) {
					fmt.Printf(textFormat, genealogyCfg.Get(key).(string), userInfo[key])
				} else {
					fmt.Printf(textFormat, key, userInfo[key])
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
				fmt.Println("----------Update Information----------")
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
				var slice = []string{"DaemonStatus"}
				for _, key := range slice {
					if genealogyCfg.Has(key) {
						fmt.Printf(textFormat, genealogyCfg.Get(key).(string), daemonInfo[key])
					} else {
						fmt.Printf(textFormat, key, daemonInfo[key])
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
	getCmd.Flags().BoolP("nic", "", false, "Get NIC information")
	getCmd.Flags().BoolP("time", "", false, "Get Time information")
	getCmd.Flags().BoolP("user", "", false, "Get User information")
	getCmd.Flags().BoolP("update", "", false, "Get Update information")
	getCmd.Flags().BoolP("only", "", false, "Get update package information only")

	getCmd.Flags().BoolP("help", "h", false, "help for get")
	rootCmd.AddCommand(getCmd)
}
