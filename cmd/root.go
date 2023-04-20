/*
File: root.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 10:02:55

Description: 程序未带子命令或参数时执行
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yhyj/eniac/function"
)

// 在没有任何子命令的情况下调用时的基本命令
var rootCmd = &cobra.Command{
	Use:   "eniac",
	Short: "for system interaction",
	Long:  `Eniac is a system interactive command line tool`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// 由main.main调用
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var varHome = function.GetVariable("HOME")
var cfgFile = varHome + "/.config/eniac/config.toml"

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "help for Eniac")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "Config file")
}
