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

// 配置项
var (
	// 允许用户修改的配置项
	// 使用默认值的配置项
	colorful  = true
	cycle     = true
	biosItems = []string{
		"BIOSVendor",
		"BIOSVersion",
		"BIOSDate",
	}
	boardItems = []string{
		"BoardVendor",
		"BoardName",
		"BoardVersion",
	}
	cpuItems = []string{
		"CPUModel",
		"CPUNumber",
		"CPUCores",
		"CPUThreads",
		"CPUCache",
	}
	cpuCacheUnit = "KB"
	gpuItems     = []string{
		"GPUAddress",
		"GPUDriver",
		"GPUProduct",
		"GPUVendor",
	}
	loadItems = []string{
		"Load1",
		"Load5",
		"Load15",
		"Process",
	}
	memoryItems = []string{
		"MemoryUsedPercent",
		"MemoryTotal",
		"MemoryUsed",
		"MemoryAvail",
		"MemoryFree",
		"MemoryBuffCache",
		"MemoryShared",
	}
	memoryDataUnit    = "GB"
	memoryPercentUnit = "%"
	nicItems          = []string{
		"NicName",
		"NicMacAddress",
		"NicDriver",
		"NicVendor",
		"NicProduct",
		"NicPCIAddress",
		"NicSpeed",
		"NicDuplex",
	}
	osItems = []string{
		"OS",
		"CurrentKernel",
		"Platform",
		"Arch",
		"TimeZone",
		"Hostname",
	}
	productItems = []string{
		"ProductVendor",
		"ProductName",
	}
	storageItems = []string{
		"StorageName",
		"StorageSize",
		"StorageType",
		"StorageDriver",
		"StorageVendor",
		"StorageModel",
		"StorageSerial",
		"StorageRemovable",
	}
	swapItemsAvailable = []string{
		"SwapTotal",
		"SwapFree",
	}
	swapItemsUnavailable = []string{
		"SwapStatus",
	}
	swapDataUnit    = "GB"
	swapPercentUnit = "%"
	timeItems       = []string{
		"StartTime",
		"Uptime",
		"BootTime",
	}
	userItems = []string{
		"UserName",
		"User",
		"UserUid",
		"UserGid",
		"UserHomeDir",
	}
)

// 配置
var appConfig = Config{
	Main: MainConfig{
		Colorful: colorful,
		Cycle:    cycle,
	},
	Genealogy: GenealogyConfig{
		Bios: BiosConfig{
			Items: biosItems,
		},
		Board: BoardConfig{
			Items: boardItems,
		},
		CPU: CPUConfig{
			CacheUnit: cpuCacheUnit,
			Items:     cpuItems,
		},
		GPU: GPUConfig{
			Items: gpuItems,
		},
		Load: LoadConfig{
			Items: loadItems,
		},
		Memory: MemoryConfig{
			DataUnit:    memoryDataUnit,
			PercentUnit: memoryPercentUnit,
			Items:       memoryItems,
		},
		Nic: NicConfig{
			Items: nicItems,
		},
		OS: OSConfig{
			Items: osItems,
		},
		Product: ProductConfig{
			Items: productItems,
		},
		Storage: StorageConfig{
			Items: storageItems,
		},
		Swap: SwapConfig{
			DataUnit:    swapDataUnit,
			PercentUnit: swapPercentUnit,
			Items: SwapItemsConfig{
				Available:   swapItemsAvailable,
				Unavailable: swapItemsUnavailable,
			},
		},
		Time: TimeConfig{
			Items: timeItems,
		},
		User: UserConfig{
			Items: userItems,
		},
	},
}
