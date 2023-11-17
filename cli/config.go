/*
File: config.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-24 16:08:27

Description: 子命令`config`的实现
*/

package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/yhyj/eniac/general"
)

// isTomlFile 判断文件是不是 toml 文件
func isTomlFile(filePath string) bool {
	if strings.HasSuffix(filePath, ".toml") {
		return true
	}
	return false
}

// GetTomlConfig 读取 toml 配置文件
func GetTomlConfig(filePath string) (*toml.Tree, error) {
	if !general.FileExist(filePath) {
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

// WriteTomlConfig 写入 toml 配置文件
func WriteTomlConfig(filePath string) (int64, error) {
	// 定义一个map[string]interface{}类型的变量并赋值
	exampleConf := map[string]interface{}{
		"cpu": map[string]interface{}{
			"cache_unit": "KB",
		},
		"parts": map[string]interface{}{
			"Product": "设备",
			"Board":   "主板",
			"BIOS":    "BIOS",
			"CPU":     "处理器",
			"GPU":     "显卡",
			"Memory":  "内存",
			"Swap":    "交换分区",
			"Disk":    "磁盘",
			"NIC":     "网卡",
			"OS":      "系统",
			"Load":    "负载",
			"Time":    "时间",
			"User":    "用户",
		},
		"memory": map[string]interface{}{
			"data_unit":    "GB",
			"percent_unit": "%",
		},
		"genealogy": map[string]interface{}{
			"BIOSVendor":        "BIOS厂商",
			"BIOSVersion":       "BIOS版本",
			"BIOSDate":          "BIOS发布",
			"BoardVendor":       "主板厂商",
			"BoardName":         "主板名称",
			"BoardVersion":      "主板版本",
			"CPUModel":          "CPU型号",
			"CPUNumber":         "CPU插槽",
			"CPUCores":          "CPU核心",
			"CPUThreads":        "CPU线程",
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
			"NicName":           "网卡名称",
			"NicPCIAddress":     "PCI 地址",
			"NicMacAddress":     "MAC 地址",
			"NicSpeed":          "网卡速率",
			"NicDuplex":         "工作模式",
			"NicDriver":         "网卡驱动",
			"NicProduct":        "网卡型号",
			"NicVendor":         "网卡厂商",
			"MemoryTotal":       "内存大小",
			"MemoryUsed":        "已用内存",
			"MemoryUsedPercent": "内存占用",
			"MemoryFree":        "空闲内存",
			"MemoryShared":      "共享内存",
			"MemoryBuffCache":   "缓冲内存",
			"MemoryAvail":       "可用内存",
			"SwapDisabled":      "交换空间未启用",
			"SwapTotal":         "交换空间大小",
			"SwapFree":          "可用交换空间",
			"Process":           "进程数",
			"ProductVendor":     "设备厂商",
			"ProductName":       "设备名称",
			"StorageName":       "磁盘名称",
			"StorageType":       "磁盘类型",
			"StorageDriver":     "磁盘驱动",
			"StorageVendor":     "磁盘厂商",
			"StorageModel":      "磁盘型号",
			"StorageSerial":     "磁盘序列号",
			"StorageRemovable":  "磁盘可移除",
			"StorageSize":       "磁盘容量",
			"BootTime":          "系统启动时间",
			"Uptime":            "系统运行时长",
			"StartTime":         "系统启动用时",
			"User":              "用户名称",
			"UserName":          "用户昵称",
			"UserUid":           "用户标识",
			"UserGid":           "属组标识",
			"UserHomeDir":       "用户目录",
			"UpdateList":        "更新列表",
			"DaemonStatus":      "更新状态",
		},
		"update": map[string]interface{}{
			"record_file": "/tmp/system-checkupdates.log",
		},
	}
	// 检测配置文件是否存在
	if !general.FileExist(filePath) {
		return 0, fmt.Errorf("Open %s: no such file or directory", filePath)
	}
	// 检测配置文件是否是 toml 文件
	if !isTomlFile(filePath) {
		return 0, fmt.Errorf("Open %s: is not a toml file", filePath)
	}
	// 把 exampleConf 转换为 *toml.Tree 类型
	tree, err := toml.TreeFromMap(exampleConf)
	if err != nil {
		return 0, err
	}
	// 打开一个文件并获取 io.Writer 接口
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	return tree.WriteTo(file)
}
