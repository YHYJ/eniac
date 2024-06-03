/*
File: define_toml.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-11 14:58:58

Description: 操作 TOML 配置文件
*/

package general

import (
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
)

// 用于转换 Toml 配置树的结构体

type Config struct {
	Genealogy GenealogyConfig `toml:"genealogy"`
	Main      MainConfig      `toml:"main"`
}

type GenealogyConfig struct {
	Bios    BiosConfig    `toml:"bios"`
	Board   BoardConfig   `toml:"board"`
	CPU     CPUConfig     `toml:"cpu"`
	GPU     GPUConfig     `toml:"gpu"`
	Load    LoadConfig    `toml:"load"`
	Memory  MemoryConfig  `toml:"memory"`
	Nic     NicConfig     `toml:"nic"`
	OS      OSConfig      `toml:"os"`
	Product ProductConfig `toml:"product"`
	Storage StorageConfig `toml:"storage"`
	Swap    SwapConfig    `toml:"swap"`
	Time    TimeConfig    `toml:"time"`
	Update  UpdateConfig  `toml:"update"`
	User    UserConfig    `toml:"user"`
}
type MainConfig struct {
	Colorful bool `toml:"colorful"`
	TabStyle bool `toml:"tab_style"`
}

type BiosConfig struct {
	Items []string `toml:"items"`
}
type BoardConfig struct {
	Items []string `toml:"items"`
}
type CPUConfig struct {
	Items     []string `toml:"items"`
	CacheUnit string   `toml:"cache_unit"`
}
type GPUConfig struct {
	Items []string `toml:"items"`
}
type LoadConfig struct {
	Items []string `toml:"items"`
}
type MemoryConfig struct {
	Items       []string `toml:"items"`
	DataUnit    string   `toml:"data_unit"`
	PercentUnit string   `toml:"percent_unit"`
}
type NicConfig struct {
	Items []string `toml:"items"`
}
type OSConfig struct {
	Items []string `toml:"items"`
}
type ProductConfig struct {
	Items []string `toml:"items"`
}
type StorageConfig struct {
	Items []string `toml:"items"`
}
type SwapConfig struct {
	Items       SwapItemsConfig `toml:"items"`
	DataUnit    string          `toml:"data_unit"`
	PercentUnit string          `toml:"percent_unit"`
}
type TimeConfig struct {
	Items []string `toml:"items"`
}
type UpdateConfig struct {
	Items      []string `toml:"items"`
	RecordFile string   `toml:"record_file"`
}
type UserConfig struct {
	Items []string `toml:"items"`
}

type SwapItemsConfig struct {
	Available   []string `toml:"available"`
	Unavailable []string `toml:"unavailable"`
}

// isTomlFile 检测文件是不是 toml 文件
//
// 参数：
//   - filePath: 待检测文件路径
//
// 返回：
//   - 是 toml 文件返回 true，否则返回 false
func isTomlFile(filePath string) bool {
	if strings.HasSuffix(filePath, ".toml") {
		return true
	}
	return false
}

// GetTomlConfig 读取 toml 配置文件
//
// 参数：
//   - filePath: toml 配置文件路径
//
// 返回：
//   - toml 配置树
//   - 错误信息
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

// LoadConfigToStruct 将 Toml 配置树加载到结构体
//
// 参数：
//   - configTree: 解析 toml 配置文件得到的配置树
//
// 返回：
//   - 结构体
//   - 错误信息
func LoadConfigToStruct(configTree *toml.Tree) (*Config, error) {
	var config Config
	if err := configTree.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// WriteTomlConfig 写入 toml 配置文件
//
// 参数：
//   - filePath: toml 配置文件路径
//
// 返回：
//   - 写入的字节数
//   - 错误信息
func WriteTomlConfig(filePath string) (int64, error) {
	// 根据系统不同决定某些参数
	var (
		genealogyUpdateRecordFile = "" // 定义不同平台下的可更新安装包记录文件
	)
	if Platform == "linux" {
		genealogyUpdateRecordFile = "/tmp/system-checkupdates.log"
	} else if Platform == "darwin" {
		genealogyUpdateRecordFile = "/tmp/system-checkupdates.log"
	} else if Platform == "windows" {
	}
	// 定义一个map[string]interface{}类型的变量并赋值
	exampleConf := map[string]interface{}{
		"main": map[string]interface{}{
			"colorful":  true,
			"tab_style": false,
		},
		"genealogy": map[string]interface{}{
			"bios": map[string]interface{}{
				"items": []string{
					"BIOSVendor",
					"BIOSVersion",
					"BIOSDate",
				},
			},
			"board": map[string]interface{}{
				"items": []string{
					"BoardVendor",
					"BoardName",
					"BoardVersion",
				},
			},
			"cpu": map[string]interface{}{
				"items": []string{
					"CPUModel",
					"CPUNumber",
					"CPUCores",
					"CPUThreads",
					"CPUCache",
				},
				"cache_unit": "KB",
			},
			"gpu": map[string]interface{}{
				"items": []string{
					"GPUAddress",
					"GPUDriver",
					"GPUProduct",
					"GPUVendor",
				},
			},
			"load": map[string]interface{}{
				"items": []string{
					"Load1",
					"Load5",
					"Load15",
					"Process",
				},
			},
			"memory": map[string]interface{}{
				"items": []string{
					"MemoryUsedPercent",
					"MemoryTotal",
					"MemoryUsed",
					"MemoryAvail",
					"MemoryFree",
					"MemoryBuffCache",
					"MemoryShared",
				},
				"data_unit":    "GB",
				"percent_unit": "%",
			},
			"nic": map[string]interface{}{
				"items": []string{
					"NicName",
					"NicMacAddress",
					"NicDriver",
					"NicVendor",
					"NicProduct",
					"NicPCIAddress",
					"NicSpeed",
					"NicDuplex",
				},
			},
			"os": map[string]interface{}{
				"items": []string{
					"OS",
					"Kernel",
					"Platform",
					"Arch",
					"TimeZone",
					"Hostname",
				},
			},
			"product": map[string]interface{}{
				"items": []string{
					"ProductVendor",
					"ProductName",
				},
			},
			"storage": map[string]interface{}{
				"items": []string{
					"StorageName",
					"StorageSize",
					"StorageType",
					"StorageDriver",
					"StorageVendor",
					"StorageModel",
					"StorageSerial",
					"StorageRemovable",
				},
			},
			"swap": map[string]interface{}{
				"items": map[string]interface{}{
					"available": []string{
						"SwapTotal",
						"SwapFree",
					},
					"unavailable": []string{
						"SwapStatus",
					},
				},
				"data_unit":    "GB",
				"percent_unit": "%",
			},
			"time": map[string]interface{}{
				"items": []string{
					"StartTime",
					"Uptime",
					"BootTime",
				},
			},
			"update": map[string]interface{}{
				"items": []string{
					"UpdateDaemonStatus",
					"PackageQuantity",
					"PackageList",
				},
				"record_file": genealogyUpdateRecordFile,
			},
			"user": map[string]interface{}{
				"items": []string{
					"UserName",
					"User",
					"UserUid",
					"UserGid",
					"UserHomeDir",
				},
			},
		},
	}
	// 检测配置文件是否存在
	if !FileExist(filePath) {
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
