/*
File: config.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-23 15:09:18

Description: 程序子命令'config'时执行
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yhyj/eniac/function"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Operate configuration file",
	Long:  `Operate configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取配置文件路径
		cfgFile, _ := cmd.Flags().GetString("config")
		// 解析参数
		createFlag, _ := cmd.Flags().GetBool("create")
		forceFlag, _ := cmd.Flags().GetBool("force")
		printFlag, _ := cmd.Flags().GetBool("print")

		// 检查配置文件是否存在
		cfgFileExist := function.CheckFileExist(cfgFile)

		// 执行配置文件操作
		if createFlag {
			if cfgFileExist {
				if forceFlag {
					function.DeleteFile(cfgFile)
					function.CreateFile(cfgFile)
					function.WriteTomlConfig(cfgFile)
					fmt.Println("配置文件已重置")
				} else {
					fmt.Println("配置文件已存在")
				}
			} else {
				function.CreateFile(cfgFile)
				function.WriteTomlConfig(cfgFile)
				fmt.Println("配置文件已创建")
			}
		}

		if printFlag {
			if cfgFileExist {
				configuration, err := function.GetTomlConfig(cfgFile)
				if err != nil {
					fmt.Printf("\x1b[36;1m%s\x1b[0m\n", err)
				} else {
					fmt.Println(configuration)
				}
			} else {
				fmt.Println("未指定配置文件，默认配置文件也不存在")
			}
		}
	},
}

func init() {
	configCmd.Flags().BoolP("create", "", false, "Create a default configuration file")
	configCmd.Flags().BoolP("force", "", false, "Overwrite existing configuration files")
	configCmd.Flags().BoolP("print", "", false, "Print configuration file content")

	configCmd.Flags().BoolP("help", "h", false, "help for config")
	rootCmd.AddCommand(configCmd)
}