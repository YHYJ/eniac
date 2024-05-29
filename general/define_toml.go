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
	Cpu    CpuConfig    `toml:"cpu"`
	Memory MemoryConfig `toml:"memory"`
	Update UpdateConfig `toml:"update"`
}
type MainConfig struct {
	Colorful bool `toml:"colorful"`
}
type CpuConfig struct {
	CacheUnit string `toml:"cache_unit"`
}
type MemoryConfig struct {
	DataUnit    string `toml:"data_unit"`
	PercentUnit string `toml:"percent_unit"`
}
type UpdateConfig struct {
	RecordFile string `toml:"record_file"`
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
			"colorful": true,
		},
		"genealogy": map[string]interface{}{
			"cpu": map[string]interface{}{
				"cache_unit": "KB",
			},
			"memory": map[string]interface{}{
				"data_unit":    "GB",
				"percent_unit": "%",
			},
			"update": map[string]interface{}{
				"record_file": genealogyUpdateRecordFile,
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
