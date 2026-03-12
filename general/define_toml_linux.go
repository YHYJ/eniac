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
	Basis          string   `toml:"basis"`
	ArchRecordFile string   `toml:"arch_record_file"`
	AurRecordFile  string   `toml:"aur_record_file"`
	ArchDividing   string   `toml:"arch_dividing"`
	AurDividing    string   `toml:"aur_dividing"`
	Items          []string `toml:"items"`
}

// 配置项
var (
	// 允许用户修改的配置项
	UpdateBasis          = "update-checker.timer"  // 更新检测服务状态判断依据
	ArchUpdateRecordFile = "/tmp/checker-arch.log" // Arch Linux 官方仓库可更新包记录文件
	AurUpdateRecordFile  = "/tmp/checker-aur.log"  // AUR 可更新包记录文件
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
		"LatestKernel",
		"Platform",
		"Arch",
		"TimeZone",
		"Hostname",
	}
	packageItems = []string{
		"PackageAsExplicitCount",
		"PackageAsDependencyCount",
		"PackageTotalCount",
		"PackageTotalSize",
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
	updateItems = []string{
		"UpdateCheckDaemonStatus",
		"LastCheckTime",
		"UpdatablePackageQuantity",
		"UpdatablePackageList",
	}
	updateArchDividing = "······Arch Official Repository······"
	updateAurDividing  = "········Arch User Repository········"
	userItems          = []string{
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
		Package: PackageConfig{
			Items: packageItems,
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
		Update: UpdateConfig{
			Basis:          UpdateBasis,
			ArchRecordFile: ArchUpdateRecordFile,
			AurRecordFile:  AurUpdateRecordFile,
			ArchDividing:   updateArchDividing,
			AurDividing:    updateAurDividing,
			Items:          updateItems,
		},
		User: UserConfig{
			Items: userItems,
		},
	},
}
