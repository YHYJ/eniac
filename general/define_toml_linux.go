//go:build linux

/*
File: define_toml_linux.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-11 14:58:58

Description: 操作 TOML 配置文件
*/

package general

type GenealogyConfig struct {
	Bios    BiosConfig    `toml:"bios"`
	Board   BoardConfig   `toml:"board"`
	CPU     CPUConfig     `toml:"cpu"`
	GPU     GPUConfig     `toml:"gpu"`
	Load    LoadConfig    `toml:"load"`
	Memory  MemoryConfig  `toml:"memory"`
	Nic     NicConfig     `toml:"nic"`
	OS      OSConfig      `toml:"os"`
	Package PackageConfig `toml:"package"`
	Product ProductConfig `toml:"product"`
	Storage StorageConfig `toml:"storage"`
	Swap    SwapConfig    `toml:"swap"`
	Time    TimeConfig    `toml:"time"`
	Update  UpdateConfig  `toml:"update"`
	User    UserConfig    `toml:"user"`
}
type PackageConfig struct {
	Items []string `toml:"items"`
}
type UpdateConfig struct {
	Items      []string `toml:"items"`
	RecordFile string   `toml:"record_file"`
}

// 可更新安装包记录文件
var UpdateRecordFile = "/tmp/system-checkupdates.log"

// 默认配置
var defaultConf = map[string]interface{}{
	"main": map[string]interface{}{
		"colorful": true,
		"cycle":    true,
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
				"CurrentKernel",
				"LatestKernel",
				"Platform",
				"Arch",
				"TimeZone",
				"Hostname",
			},
		},
		"package": map[string]interface{}{
			"items": []string{
				"PackageAsExplicitCount",
				"PackageAsDependencyCount",
				"PackageTotalCount",
				"PackageTotalSize",
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
				"LastCheckTime",
				"UpdatablePackageQuantity",
				"UpdatablePackageList",
			},
			"record_file": UpdateRecordFile,
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
