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

	"github.com/spf13/cobra"
	"github.com/yhyj/eniac/function"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get system information",
	Long:  `Get system information`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取配置项
		cfgFile, _ := cmd.Flags().GetString("config")
		confTree, err := function.GetTomlConfig(cfgFile)
		dataUnit := ""
		percentUnit := ""
		if err != nil {
			fmt.Printf("\x1b[36;1m%s\x1b[0m\n", err)
		} else {
			// 获取配置项
			dataUnit = confTree.Get("memory.data_unit").(string)
			percentUnit = confTree.Get("memory.percent_unit").(string)
		}
		// 解析参数执行对应函数
		loadFlag, _ := cmd.Flags().GetBool("load")
		memFlag, _ := cmd.Flags().GetBool("mem")
		swapFlag, _ := cmd.Flags().GetBool("swap")
		userFlag, _ := cmd.Flags().GetBool("user")
		if loadFlag {
			loadInfo, _ := function.GetLoadInfo()
			fmt.Println(loadInfo)
		}
		if memFlag {
			memInfo, _ := function.GetMemoryInfo(dataUnit, percentUnit)
			fmt.Println(memInfo)
		}
		if swapFlag {
			swapInfo, _ := function.GetSwapInfo(dataUnit)
			fmt.Println(swapInfo)
		}
		if userFlag {
			userInfo, _ := function.GetUserInfo()
			fmt.Println(userInfo)
		}
	},
}

func init() {
	getCmd.Flags().BoolP("load", "l", false, "Get load information")
	getCmd.Flags().BoolP("mem", "m", false, "Get memory information")
	getCmd.Flags().BoolP("swap", "s", false, "Get swap information")
	getCmd.Flags().BoolP("user", "u", false, "Get user information")

	getCmd.Flags().BoolP("help", "h", false, "help for get")
	rootCmd.AddCommand(getCmd)
}
