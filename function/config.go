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
		return nil, fmt.Errorf("Open %s: no such file or directory", filePath)
	}
	if !isTomlFile(filePath) {
		return nil, fmt.Errorf("Open %s: is not a toml file", filePath)
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
		"cpu": map[string]interface{}{
			"cache_unit": "KB",
		},
		"memory": map[string]interface{}{
			"data_unit":    "GB",
			"percent_unit": "%",
		},
		"nic": map[string]interface{}{
			// 设备的PCI ID，使用命令`lspci`查看，例如`lspci`结果是'01:00.0'，则实际值为'0000:01:00.0'
			"address": "0000:01:00.0",
		},
		"genealogy": map[string]interface{}{
			"BIOSVendor":        "BIOS厂商",
			"BIOSVersion":       "BIOS版本",
			"BIOSDate":          "BIOS发布日期",
			"BoardVendor":       "主板厂商",
			"BoardName":         "主板名称",
			"BoardVersion":      "主板版本",
			"CPUModel":          "CPU型号",
			"CPUNumber":         "CPU插槽数",
			"CPUCores":          "CPU核心数",
			"CPUThreads":        "CPU线程数",
			"CPUCache":          "CPU缓存",
			"GPUAddress":        "显卡地址",
			"GPUDriver":         "显卡驱动",
			"GPUProduct":        "显卡型号",
			"GPUVendor":         "显卡厂商",
			"OS":                "操作系统",
			"Arch":              "系统架构",
			"Kernel":            "内核版本",
			"Platform":          "系统类型",
			"Hostname":          "主机名称",
			"TimeZone":          "系统时区",
			"Load1":             "1分钟负载",
			"Load5":             "5分钟负载",
			"Load15":            "15分钟负载",
			"NicAddress":        "网卡地址",
			"NicDriver":         "网卡驱动",
			"NicProduct":        "网卡型号",
			"NicVendor":         "网卡厂商",
			"MemoryTotal":       "内存大小",
			"MemoryUsed":        "已用内存",
			"MemoryUsedPercent": "内存使用率",
			"MemoryFree":        "空闲内存",
			"MemoryShared":      "共享内存",
			"MemoryBuffCache":   "缓冲内存",
			"MemoryAvail":       "可用内存",
			"SwapDisabled":      "交换空间未启用",
			"SwapTotal":         "总交换空间",
			"SwapFree":          "空闲交换空间",
			"Process":           "进程数",
			"ProductVendor":     "产品厂商",
			"ProductName":       "产品名称",
			"StorageName":       "存储设备名称",
			"StorageType":       "存储设备类型",
			"StorageDriver":     "存储设备驱动",
			"StorageVendor":     "存储设备厂商",
			"StorageModel":      "存储设备型号",
			"StorageSerial":     "存储设备序列号",
			"StorageRemovable":  "存储设备可移除",
			"StorageSize":       "存储设备容量",
			"BootTime":          "系统启动时间",
			"Uptime":            "系统运行时间",
			"StartTime":         "系统启动用时",
			"User":              "用户名",
			"UserName":          "昵称",
			"UserUid":           "用户ID",
			"UserGid":           "属组ID",
			"UserHomeDir":       "用户主目录",
			"UpdateList":        "更新列表",
			"DaemonStatus":      "更新状态",
		},
		"update": map[string]interface{}{
			"record_file": "/tmp/system-checkupdates.log",
		},
	}
	if !FileExist(filePath) {
		return 0, fmt.Errorf("Open %s: no such file or directory", filePath)
	}
	if !isTomlFile(filePath) {
		return 0, fmt.Errorf("Open %s: is not a toml file", filePath)
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
