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
		userFlag, _ := cmd.Flags().GetBool("user")
		if userFlag {
			userInfo, _ := function.GetUserInfo()
			fmt.Println(userInfo)
		}
	},
}

func init() {
	getCmd.Flags().BoolP("help", "h", false, "Help for get")
	getCmd.Flags().BoolP("user", "u", false, "Get user info")
	rootCmd.AddCommand(getCmd)
}
