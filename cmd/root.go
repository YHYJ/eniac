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
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yhyj/eniac/general"
)

var rootCmd = &cobra.Command{
	Use:   "eniac",
	Short: "For system interaction",
	Long:  `eniac is a system interactive command line tool.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile = filepath.Join(general.UserInfo.HomeDir, ".config", "eniac", "config.toml")

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "help for eniac")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "Specify configuration file")
}
