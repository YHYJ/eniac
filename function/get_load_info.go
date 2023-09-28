/*
File: get_load_info.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-20 11:20:33

Description: 获取负载信息
*/

package function

// GetLoadInfo 获取负载信息
func GetLoadInfo() (loadInfo map[string]interface{}, err error) {
	loadInfo = make(map[string]interface{})
	loadInfo["Load1"] = loadData.Load1   // 1分钟内的负载
	loadInfo["Load5"] = loadData.Load5   // 5分钟内的负载
	loadInfo["Load15"] = loadData.Load15 // 15分钟内的负载

	return loadInfo, err
}
