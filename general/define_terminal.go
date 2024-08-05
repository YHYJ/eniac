/*
File: define_terminal.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-08-05 08:36:27

Description: 获取终端信息
*/

package general

import (
	"os"

	"golang.org/x/term"
)

// GetTerminalSize 获取终端尺寸
//
// 返回：
//   - 终端宽度（列）
//   - 终端高度（行）
//   - 错误信息
func GetTerminalSize() (width, height int, err error) {
	width, height, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}
