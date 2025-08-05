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
	Items          []string `toml:"items"`
	Basis          string   `toml:"basis"`
	ArchRecordFile string   `toml:"arch_record_file"`
	AurRecordFile  string   `toml:"aur_record_file"`
	ArchDividing   string   `toml:"arch_dividing"`
	AurDividing    string   `toml:"aur_dividing"`
}

var (
	UpdateBasis          = "update-checker.timer"  // 更新检测服务状态判断依据
	ArchUpdateRecordFile = "/tmp/checker-arch.log" // Arch Linux 官方仓库可更新包记录文件
	AurUpdateRecordFile  = "/tmp/checker-aur.log"  // AUR 可更新包记录文件
)

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
				"LatestKernel",
				"Platform",
				"Arch",
				"TimeZone",
				"Hostname",
			},
		},
		"package": map[string]any{
			"items": []string{
				"PackageAsExplicitCount",
				"PackageAsDependencyCount",
				"PackageTotalCount",
				"PackageTotalSize",
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
		"update": map[string]any{
			"items": []string{
				"UpdateCheckDaemonStatus",
				"LastCheckTime",
				"UpdatablePackageQuantity",
				"UpdatablePackageList",
			},
			"basis":            UpdateBasis,
			"arch_record_file": ArchUpdateRecordFile,
			"aur_record_file":  AurUpdateRecordFile,
			"arch_dividing":    "······Arch Official Repository······",
			"aur_dividing":     "········Arch User Repository········",
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
