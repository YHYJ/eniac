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
)

var rootCmd = &cobra.Command{
	Use:   "eniac",
	Short: "For system interaction",
	Long:  `Eniac is a system interactive command line tool.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var cfgFile = "/etc/eniac/config.toml"

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "help for Eniac")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "Config file")
}
