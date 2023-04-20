/*
File: get_loadinfo.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 11:20:33

Description: 获取负载信息
*/

package function

import "github.com/shirou/gopsutil/load"

// LoadInfoStruct 负载信息结构体
type LoadInfoStruct struct {
	Load1  float64 `json:"load1"`  // 1分钟平均负载
	Load5  float64 `json:"load5"`  // 5分钟平均负载
	Load15 float64 `json:"load15"` // 15分钟平均负载
}

// GetLoadInfo 获取负载信息
func GetLoadInfo() (loadInfo LoadInfoStruct, err error) {
	info, _ := load.Avg()
	loadInfo.Load1 = info.Load1
	loadInfo.Load5 = info.Load5
	loadInfo.Load15 = info.Load15

	return loadInfo, nil
}
