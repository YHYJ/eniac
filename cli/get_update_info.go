/*
File: get_update_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-25 14:30:33

Description: 获取更新信息
*/

package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yhyj/eniac/general"
)

// GetUpdateInfo 读取更新信息记录文件
//
//   - 参数 line=0 时读取全部行
//
// 参数：
//   - filePath: 更新信息记录文件路径
//   - line: 读取指定行
//
// 返回：
//   - 更新信息
//   - 错误信息
func GetUpdateInfo(filePath string, line int) ([]string, error) {
	if filePath != "" && general.FileExist(filePath) {
		var textSlice []string
		// 打开文件
		text, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer text.Close()

		// 创建一个扫描器对象按行遍历
		scanner := bufio.NewScanner(text)
		// 行计数
		count := 1
		// 逐行读取，输出指定行
		for scanner.Scan() {
			if line == count {
				textSlice = append(textSlice, scanner.Text())
				break
			}
			textSlice = append(textSlice, scanner.Text())
			count++
		}
		return textSlice, nil
	} else if filePath == "" {
		return nil, nil
	} else {
		return nil, fmt.Errorf("open %s: no such file", filePath)
	}
}

// GetUpdateDaemonInfo 获取更新检测服务的信息
//
// 返回：
//   - 更新检测服务的信息
//   - 错误信息
func GetUpdateDaemonInfo() (map[string]interface{}, error) {
	daemonInfo := make(map[string]interface{})
	daemonArgs := []string{"is-active", "system-checkupdates.timer"}
	daemonStatus, err := general.RunCommandGetResult("systemctl", daemonArgs)
	if err != nil {
		return nil, err
	}
	daemonInfo["DaemonStatus"] = daemonStatus

	return daemonInfo, nil
}
