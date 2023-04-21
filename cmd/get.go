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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get system information",
	Long:  `Get system information`,
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置文件
		cfgFile, _ := cmd.Flags().GetString("config")
		confTree, err := function.GetTomlConfig(cfgFile)
		if err != nil {
			fmt.Printf("\x1b[36;1m%s\x1b[0m\n", err)
		} else {
			// 获取配置项
			memoryCfg := confTree.Get("memory").(*toml.Tree)
			dataUnit := memoryCfg.Get("data_unit").(string)
			percentUnit := memoryCfg.Get("percent_unit").(string)
			// 获取系统信息（集中获取一次后分配到不同的参数）
			var sysInfo sysinfo.SysInfo
			sysInfo.GetSysInfo()
			// 解析参数
			biosFlag, _ := cmd.Flags().GetBool("bios")
			boardFlag, _ := cmd.Flags().GetBool("board")
			cpuFlag, _ := cmd.Flags().GetBool("cpu")
			loadFlag, _ := cmd.Flags().GetBool("load")
			memFlag, _ := cmd.Flags().GetBool("mem")
			osFlag, _ := cmd.Flags().GetBool("os")
			procsFlag, _ := cmd.Flags().GetBool("procs")
			productFlag, _ := cmd.Flags().GetBool("product")
			storageFlag, _ := cmd.Flags().GetBool("storage")
			swapFlag, _ := cmd.Flags().GetBool("swap")
			timeFlag, _ := cmd.Flags().GetBool("time")
			userFlag, _ := cmd.Flags().GetBool("user")
			// 执行对应函数
			if biosFlag {
				biosInfo, _ := function.GetBIOSInfo(sysInfo)
				fmt.Println(biosInfo)
			}
			if boardFlag {
				boardInfo, _ := function.GetBoardInfo(sysInfo)
				fmt.Println(boardInfo)
			}
			if cpuFlag {
				cpuInfo, _ := function.GetCPUInfo(sysInfo)
				fmt.Println(cpuInfo)
			}
			if loadFlag {
				loadInfo, _ := function.GetLoadInfo()
				fmt.Println(loadInfo)
			}
			if memFlag {
				memInfo, _ := function.GetMemoryInfo(dataUnit, percentUnit)
				fmt.Println(memInfo)
			}
			if osFlag {
				osInfo, _ := function.GetOSInfo(sysInfo)
				fmt.Println(osInfo)
			}
			if procsFlag {
				procsInfo, _ := function.GetProcsInfo()
				fmt.Println(procsInfo)
			}
			if productFlag {
				productInfo, _ := function.GetProductInfo(sysInfo)
				fmt.Println(productInfo)
			}
			if storageFlag {
				storageInfo, _ := function.GetStorageInfo(sysInfo)
				fmt.Println(storageInfo)
			}
			if swapFlag {
				swapInfo, _ := function.GetSwapInfo(dataUnit)
				fmt.Println(swapInfo)
			}
			if timeFlag {
				timeInfo, _ := function.GetTimeInfo()
				fmt.Println(timeInfo)
			}
			if userFlag {
				userInfo, _ := function.GetUserInfo()
				fmt.Println(userInfo)
			}
		}
	},
}

func init() {
	getCmd.Flags().BoolP("bios", "", false, "Get BIOS information")
	getCmd.Flags().BoolP("board", "", false, "Get Board information")
	getCmd.Flags().BoolP("cpu", "", false, "Get CPU information")
	getCmd.Flags().BoolP("load", "", false, "Get Load information")
	getCmd.Flags().BoolP("mem", "", false, "Get Memory information")
	getCmd.Flags().BoolP("os", "", false, "Get OS information")
	getCmd.Flags().BoolP("procs", "", false, "Get Procs information")
	getCmd.Flags().BoolP("product", "", false, "Get Product information")
	getCmd.Flags().BoolP("storage", "", false, "Get Storage information")
	getCmd.Flags().BoolP("swap", "", false, "Get Swap information")
	getCmd.Flags().BoolP("time", "", false, "Get Time information")
	getCmd.Flags().BoolP("user", "", false, "Get User information")

	getCmd.Flags().BoolP("help", "h", false, "help for get")
	rootCmd.AddCommand(getCmd)
}
