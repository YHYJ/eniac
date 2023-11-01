/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 09:50:55

Description: 子命令`version`功能函数
*/

package general

import "fmt"

// 程序信息
const (
	Name    string = "Eniac"
	Version string = "v1.1.2"
	Project string = "github.com/yhyj/eniac"
)

// 编译信息
var (
	GitCommitHash string = "unknown"
	BuildTime     string = "unknown"
	BuildBy       string = "unknown"
)

func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", Version)
	if !only {
		programInfo = fmt.Sprintf("%s version: %s\nGit commit hash: %s\nBuilt on: %s\nBuilt by: %s\n", Name, Version, GitCommitHash, BuildTime, BuildBy)
	}
	return programInfo
}
