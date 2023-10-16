/*
File: get_update_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-25 14:30:33

Description: 获取更新信息
*/

package function

import (
	"bufio"
	"fmt"
	"os"
)

// 读取更新信息文件，参数line为0时读取全部行
func GetUpdateInfo(filePath string, line int) ([]string, error) {
	var textSlice []string
	if !FileExist(filePath) {
		return nil, fmt.Errorf("open %s: no such file", filePath)
	}
	// 打开文件
	text, err := os.Open(filePath)
	// 相当于Python的with语句
	defer text.Close()
	// 处理错误
	if err != nil {
		return nil, err
	}
	// 行计数
	count := 1
	// 创建一个扫描器对象按行遍历
	scanner := bufio.NewScanner(text)
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
}

// 获取更新检测服务信息
func GetUpdateDaemonInfo() (map[string]interface{}, error) {
	daemonInfo := make(map[string]interface{})
	daemonArgs := []string{"is-active", "system-checkupdates.timer"}
	daemonStatus, err := RunCommandGetResult("systemctl", daemonArgs)
	if err != nil {
		return nil, err
	}
	daemonInfo["DaemonStatus"] = daemonStatus

	return daemonInfo, nil
}
