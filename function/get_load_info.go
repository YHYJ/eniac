/*
File: get_load_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 11:20:33

Description: 获取负载信息
*/

package function

import "github.com/shirou/gopsutil/load"

// GetLoadInfo 获取负载信息
func GetLoadInfo() (loadInfo map[string]interface{}, err error) {
	info, _ := load.Avg()
	loadInfo = make(map[string]interface{})
	loadInfo["Load1"] = info.Load1   // 1分钟内的负载
	loadInfo["Load5"] = info.Load5   // 5分钟内的负载
	loadInfo["Load15"] = info.Load15 // 15分钟内的负载

	return loadInfo, err
}
