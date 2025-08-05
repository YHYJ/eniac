//go:build darwin

/*
File: define_toml_darwin.go
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
	Product ProductConfig `toml:"product"`
	Storage StorageConfig `toml:"storage"`
	Swap    SwapConfig    `toml:"swap"`
	Time    TimeConfig    `toml:"time"`
	User    UserConfig    `toml:"user"`
}

// 默认配置
var defaultConf = map[string]any{
	"main": map[string]any{
		"colorful": true,
		"cycle":    true,
	},
	"genealogy": map[string]any{
		"bios": map[string]any{
			"items": []string{
				"BIOSVendor",
				"BIOSVersion",
				"BIOSDate",
			},
		},
		"board": map[string]any{
			"items": []string{
				"BoardVendor",
				"BoardName",
				"BoardVersion",
			},
		},
		"cpu": map[string]any{
			"items": []string{
				"CPUModel",
				"CPUNumber",
				"CPUCores",
				"CPUThreads",
				"CPUCache",
			},
			"cache_unit": "KB",
		},
		"gpu": map[string]any{
			"items": []string{
				"GPUAddress",
				"GPUDriver",
				"GPUProduct",
				"GPUVendor",
			},
		},
		"load": map[string]any{
			"items": []string{
				"Load1",
				"Load5",
				"Load15",
				"Process",
			},
		},
		"memory": map[string]any{
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
		"nic": map[string]any{
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
		"os": map[string]any{
			"items": []string{
				"OS",
				"CurrentKernel",
				"Platform",
				"Arch",
				"TimeZone",
				"Hostname",
			},
		},
		"product": map[string]any{
			"items": []string{
				"ProductVendor",
				"ProductName",
			},
		},
		"storage": map[string]any{
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
		"swap": map[string]any{
			"items": map[string]any{
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
		"time": map[string]any{
			"items": []string{
				"StartTime",
				"Uptime",
				"BootTime",
			},
		},
		"user": map[string]any{
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
