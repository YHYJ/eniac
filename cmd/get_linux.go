//go:build linux

/*
File: get_linux.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 10:53:10

Description: 执行子命令 'get'
*/

package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/yhyj/eniac/cli"
	"github.com/yhyj/eniac/general"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get system information",
	Long:  `Get system information, no flag entry alternate mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取配置文件路径
		configFile, _ := cmd.Flags().GetString("config")

		// 读取配置文件
		confTree, err := general.GetTomlConfig(configFile)
		if err != nil {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
			return
		}

		if cmd.Flags().NFlag() == 0 {
			// 抓取系统信息
			cli.GrabInformationToTab(confTree)
		} else {
			// 解析参数
			allFlag, _ := cmd.Flags().GetBool("all")
			allFlags := make(map[string]bool)
			if allFlag {
				allFlags["biosFlag"] = true
				allFlags["boardFlag"] = true
				allFlags["gpuFlag"] = true
				allFlags["cpuFlag"] = true
				allFlags["loadFlag"] = true
				allFlags["memoryFlag"] = true
				allFlags["osFlag"] = true
				allFlags["productFlag"] = true
				allFlags["storageFlag"] = true
				allFlags["swapFlag"] = true
				allFlags["nicFlag"] = true
				allFlags["timeFlag"] = true
				allFlags["userFlag"] = true
				allFlags["updateFlag"] = true
				allFlags["onlyFlag"] = false
			} else {
				allFlags["biosFlag"], _ = cmd.Flags().GetBool("bios")
				allFlags["boardFlag"], _ = cmd.Flags().GetBool("board")
				allFlags["gpuFlag"], _ = cmd.Flags().GetBool("gpu")
				allFlags["cpuFlag"], _ = cmd.Flags().GetBool("cpu")
				allFlags["loadFlag"], _ = cmd.Flags().GetBool("load")
				allFlags["memoryFlag"], _ = cmd.Flags().GetBool("memory")
				allFlags["osFlag"], _ = cmd.Flags().GetBool("os")
				allFlags["productFlag"], _ = cmd.Flags().GetBool("product")
				allFlags["storageFlag"], _ = cmd.Flags().GetBool("storage")
				allFlags["swapFlag"], _ = cmd.Flags().GetBool("swap")
				allFlags["nicFlag"], _ = cmd.Flags().GetBool("nic")
				allFlags["timeFlag"], _ = cmd.Flags().GetBool("time")
				allFlags["userFlag"], _ = cmd.Flags().GetBool("user")
				allFlags["updateFlag"], _ = cmd.Flags().GetBool("update")
				allFlags["onlyFlag"], _ = cmd.Flags().GetBool("only")
			}

			// 抓取系统信息
			cli.GrabInformationToTable(confTree, allFlags)

			// 显示通知
			general.Notification()
		}
	},
}

func init() {
	getCmd.Flags().Bool("all", false, "Get all information")
	getCmd.Flags().Bool("bios", false, "Get BIOS information")
	getCmd.Flags().Bool("board", false, "Get Board information")
	getCmd.Flags().Bool("cpu", false, "Get CPU information")
	getCmd.Flags().Bool("gpu", false, "Get GPU information")
	getCmd.Flags().Bool("load", false, "Get Load information")
	getCmd.Flags().Bool("memory", false, "Get Memory information")
	getCmd.Flags().Bool("os", false, "Get OS information")
	getCmd.Flags().Bool("product", false, "Get Product information")
	getCmd.Flags().Bool("storage", false, "Get Storage information")
	getCmd.Flags().Bool("swap", false, "Get Swap information")
	getCmd.Flags().Bool("nic", false, "Get NIC information")
	getCmd.Flags().Bool("time", false, "Get Time information")
	getCmd.Flags().Bool("user", false, "Get User information")
	getCmd.Flags().Bool("update", false, "Get Update information")
	getCmd.Flags().Bool("only", false, "Get update package information only")

	getCmd.Flags().BoolP("help", "h", false, "help for get command")
	rootCmd.AddCommand(getCmd)
}
