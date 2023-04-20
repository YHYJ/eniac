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
			memInfo, _ := function.GetMemoryInfo()
			fmt.Println(memInfo)
		}
		if swapFlag {
			swapInfo, _ := function.GetSwapInfo()
			fmt.Println(swapInfo)
		}
		if userFlag {
			userInfo, _ := function.GetUserInfo()
			fmt.Println(userInfo)
		}
	},
}

func init() {
	getCmd.Flags().BoolP("load", "l", false, "Get load info")
	getCmd.Flags().BoolP("mem", "m", false, "Get memory info")
	getCmd.Flags().BoolP("swap", "s", false, "Get swap info")
	getCmd.Flags().BoolP("user", "u", false, "Get user info")

	getCmd.Flags().BoolP("help", "h", false, "Help for get")
	rootCmd.AddCommand(getCmd)
}
