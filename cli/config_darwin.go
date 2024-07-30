//go:build darwin

/*
File: config.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-24 16:08:27

Description: 子命令 'config' 的实现
*/

package cli

import (
	"strings"

	"github.com/gookit/color"
	"github.com/yhyj/eniac/general"
)

// CreateConfigFile 创建配置文件
//
// 参数：
//   - configFile: 配置文件路径
func CreateConfigFile(configFile string) {
	// 检查配置文件是否存在
	fileExist := general.FileExist(configFile)

	// 检测并创建配置文件
	if fileExist {
		// 询问是否覆写已存在的配置文件
		question := color.Sprintf(general.OverWriteTips, "Configuration")
		overWrite, err := general.AskUser(general.QuestionText(question), []string{"y", "N"})
		if err != nil {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
			return
		}

		switch overWrite {
		case "y":
			if err := general.DeleteFile(configFile); err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				return
			}
			if err := general.CreateFile(configFile); err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				return
			}
			if _, err := general.WriteTomlConfig(configFile); err != nil {
				fileName, lineNo := general.GetCallerInfo()
				color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				return
			}
			color.Printf("Create %s: %s\n", general.PrimaryText(configFile), general.SuccessText("file overwritten"))
		case "n":
			return
		default:
			color.Printf("%s\n", strings.Repeat(general.Separator3st, len(question)))
			color.Warn.Tips("%s: %s", "Unexpected answer", overWrite)
			return
		}
	} else {
		if err := general.CreateFile(configFile); err != nil {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
			return
		}
		if _, err := general.WriteTomlConfig(configFile); err != nil {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s %s\n", general.DangerText(general.ErrorInfoFlag), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
			return
		}
		color.Printf("Create %s: %s\n", general.PrimaryText(configFile), general.SuccessText("file created"))
	}
}
