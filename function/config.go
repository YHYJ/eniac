/*
File: config.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-24 16:08:27

Description: 子命令`config`的实现
*/

package function

import (
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
)

// 判断文件是不是toml文件
func isTomlFile(filePath string) bool {
	if strings.HasSuffix(filePath, ".toml") {
		return true
	}
	return false
}

// 读取toml配置文件
func GetTomlConfig(filePath string) (*toml.Tree, error) {
	if !FileExist(filePath) {
		return nil, fmt.Errorf("open %s: no such file or directory", filePath)
	}
	if !isTomlFile(filePath) {
		return nil, fmt.Errorf("open %s: is not a toml file", filePath)
	}
	tree, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}
	return tree, nil
}

// 写入toml配置文件
func WriteTomlConfig(filePath string) (int64, error) {
	// 定义一个map[string]interface{}类型的变量并赋值
	exampleConf := map[string]interface{}{
		"cpu": map[string]string{
			"cache_unit": "KB",
		},
		"memory": map[string]string{
			"data_unit":    "GB",
			"percent_unit": "%",
		},
		"genealogy": map[string]string{
			"BIOSVendor":        "BIOS厂商",
			"BIOSVersion":       "BIOS版本",
			"BIOSDate":          "BIOS发布日期",
			"BoardVendor":       "主板厂商",
			"BoardName":         "主板名称",
			"BoardVersion":      "主板版本",
			"CPUModel":          "CPU型号",
			"CPUNumber":         "CPU数量",
			"CPUCores":          "CPU核心数",
			"CPUThreads":        "CPU线程数",
			"CPUCache":          "CPU缓存",
			"OS":                "操作系统",
			"Arch":              "操作系统架构",
			"Kernel":            "内核版本",
			"Platform":          "平台",
			"Hostname":          "主机名",
			"TimeZone":          "时区",
			"Load1":             "1分钟负载",
			"Load5":             "5分钟负载",
			"Load15":            "15分钟负载",
			"MemoryTotal":       "总物理内存",
			"MemoryUsed":        "已用物理内存",
			"MemoryUsedPercent": "物理内存使用率",
			"MemoryFree":        "空闲物理内存",
			"MemoryShared":      "共享物理内存",
			"MemoryBuffCache":   "缓冲物理内存",
			"MemoryAvail":       "可用物理内存",
			"SwapTotal":         "总交换空间",
			"SwapFree":          "空闲交换空间",
			"Process":           "进程数",
			"ProductVendor":     "产品厂商",
			"ProductName":       "产品名称",
			"StorageName":       "存储设备名称",
			"StorageDriver":     "存储设备驱动",
			"StorageVendor":     "存储设备厂商",
			"StorageModel":      "存储设备型号",
			"StorageSerial":     "存储设备序列号",
			"StorageSize":       "存储设备容量",
			"BootTime":          "系统启动时间",
			"Uptime":            "系统运行时间",
			"User":              "用户名",
			"UserName":          "用户昵称",
			"UserUid":           "用户ID",
			"UserGid":           "用户组ID",
			"UserHomeDir":       "用户主目录",
		},
	}
	if !FileExist(filePath) {
		return 0, fmt.Errorf("open %s: no such file or directory", filePath)
	}
	if !isTomlFile(filePath) {
		return 0, fmt.Errorf("open %s: is not a toml file", filePath)
	}
	// 把exampleConf转换为*toml.Tree
	tree, err := toml.TreeFromMap(exampleConf)
	if err != nil {
		return 0, err
	}
	// 打开一个文件并获取io.Writer接口
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	return tree.WriteTo(file)
}
