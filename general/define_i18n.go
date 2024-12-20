/*
File: define_i18n.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-05-29 16:21:23

Description: 国际化
*/

package general

// 各部分的名称
var PartName = map[string]map[string]string{
	"Product": {"zh": "设备", "en": "Product"},
	"Board":   {"zh": "主板", "en": "Board"},
	"BIOS":    {"zh": "BIOS", "en": "BIOS"},
	"CPU":     {"zh": "处理器", "en": "CPU"},
	"GPU":     {"zh": "显卡", "en": "GPU"},
	"Memory":  {"zh": "内存", "en": "Memory"},
	"Swap":    {"zh": "交换空间", "en": "Swap"},
	"Disk":    {"zh": "磁盘", "en": "Disk"},
	"NIC":     {"zh": "网卡", "en": "NIC"},
	"OS":      {"zh": "系统", "en": "OS"},
	"Package": {"zh": "安装包", "en": "Package"},
	"Load":    {"zh": "负载", "en": "Load"},
	"Time":    {"zh": "时间", "en": "Time"},
	"User":    {"zh": "用户", "en": "User"},
	"Update":  {"zh": "更新", "en": "Update"},
}

// 各部分.条目的名称
var GenealogyName = map[string]map[string]string{
	"BIOSVendor":               {"zh": "BIOS 厂商", "en": "Vendor"},
	"BIOSVersion":              {"zh": "BIOS 版本", "en": "Version"},
	"BIOSDate":                 {"zh": "BIOS 发布日期", "en": "Release Date"},
	"BoardVendor":              {"zh": "主板厂商", "en": "Vendor"},
	"BoardName":                {"zh": "主板名称", "en": "Name"},
	"BoardVersion":             {"zh": "主板版本", "en": "Version"},
	"CPUModel":                 {"zh": "处理器型号", "en": "Model"},
	"CPUNumber":                {"zh": "处理器数量", "en": "Number"},
	"CPUCores":                 {"zh": "处理器核心", "en": "Cores"},
	"CPUThreads":               {"zh": "处理器线程", "en": "Threads"},
	"CPUCache":                 {"zh": "处理器缓存", "en": "Cache"},
	"GPUAddress":               {"zh": "显卡地址", "en": "Address"},
	"GPUDriver":                {"zh": "显卡驱动", "en": "Driver"},
	"GPUProduct":               {"zh": "显卡型号", "en": "Product"},
	"GPUVendor":                {"zh": "显卡厂商", "en": "Vendor"},
	"OS":                       {"zh": "操作系统", "en": "OS"},
	"Arch":                     {"zh": "系统架构", "en": "Arch"},
	"Kernel":                   {"zh": "内核版本", "en": "Kernel"},
	"CurrentKernel":            {"zh": "当前内核版本", "en": "Current Kernel"},
	"LatestKernel":             {"zh": "最新内核版本", "en": "Latest Kernel"},
	"Platform":                 {"zh": "系统类型", "en": "Platform"},
	"Hostname":                 {"zh": "主机名称", "en": "Hostname"},
	"TimeZone":                 {"zh": "时区", "en": "Time zone"},
	"Load1":                    {"zh": "1分钟平均负载", "en": "Load average (1 min)"},
	"Load5":                    {"zh": "5分钟平均负载", "en": "Load average (5 min)"},
	"Load15":                   {"zh": "15分钟平均负载", "en": "Load average (15 min)"},
	"NicName":                  {"zh": "网卡名称", "en": "Name"},
	"NicPCIAddress":            {"zh": "PCI 地址", "en": "PCI Address"},
	"NicMacAddress":            {"zh": "MAC 地址", "en": "MAC Address"},
	"NicSpeed":                 {"zh": "网卡速率", "en": "Speed"},
	"NicDuplex":                {"zh": "工作模式", "en": "Duplex"},
	"NicDriver":                {"zh": "网卡驱动", "en": "Driver"},
	"NicProduct":               {"zh": "网卡型号", "en": "Product"},
	"NicVendor":                {"zh": "网卡厂商", "en": "Vendor"},
	"MemoryTotal":              {"zh": "内存大小", "en": "Total"},
	"MemoryUsed":               {"zh": "已用内存", "en": "Used"},
	"MemoryUsedPercent":        {"zh": "内存占用", "en": "Used Percent"},
	"MemoryFree":               {"zh": "空闲内存", "en": "Free"},
	"MemoryShared":             {"zh": "共享内存", "en": "Shared"},
	"MemoryBuffCache":          {"zh": "缓冲内存", "en": "Buff Cache"},
	"MemoryAvail":              {"zh": "可用内存", "en": "Avail"},
	"SwapStatus":               {"zh": "交换空间状态", "en": "Swap Status"},
	"SwapTotal":                {"zh": "交换空间大小", "en": "Total"},
	"SwapFree":                 {"zh": "空闲交换空间", "en": "Free"},
	"Process":                  {"zh": "进程数", "en": "Process"},
	"PackageTotalCount":        {"zh": "已安装包总数", "en": "Installed Package Total Count"},
	"PackageTotalSize":         {"zh": "已安装包总大小", "en": "Installed Package Total Size"},
	"PackageAsExplicitCount":   {"zh": "单独指定安装包数量", "en": "As Explicit Package Count"},
	"PackageAsDependencyCount": {"zh": "作为依赖安装包数量", "en": "As Dependency Package Count"},
	"ProductVendor":            {"zh": "设备厂商", "en": "Vendor"},
	"ProductName":              {"zh": "设备名称", "en": "Name"},
	"StorageName":              {"zh": "磁盘名称", "en": "Name"},
	"StorageType":              {"zh": "磁盘类型", "en": "Type"},
	"StorageDriver":            {"zh": "磁盘驱动", "en": "Driver"},
	"StorageVendor":            {"zh": "磁盘厂商", "en": "Vendor"},
	"StorageModel":             {"zh": "磁盘型号", "en": "Model"},
	"StorageSerial":            {"zh": "磁盘序列号", "en": "Serial"},
	"StorageRemovable":         {"zh": "磁盘可移除", "en": "Removable"},
	"StorageSize":              {"zh": "磁盘容量", "en": "Size"},
	"BootTime":                 {"zh": "系统启动时间", "en": "Boot Time"},
	"Uptime":                   {"zh": "系统运行时长", "en": "Uptime"},
	"StartTime":                {"zh": "系统启动用时", "en": "Startup Time"},
	"User":                     {"zh": "用户名称", "en": "User"},
	"UserName":                 {"zh": "用户昵称", "en": "Username"},
	"UserUid":                  {"zh": "用户标识", "en": "UID"},
	"UserGid":                  {"zh": "用户组标识", "en": "GID"},
	"UserHomeDir":              {"zh": "用户主目录", "en": "Home Dir"},
	"UpdateCheckDaemonStatus":  {"zh": "更新检测服务", "en": "Update Check Daemon"},
	"LastCheckTime":            {"zh": "最后检查时间", "en": "Last Check Time"},
	"UpdatablePackageList":     {"zh": "可更新包列表", "en": "Updatable Package List"},
	"UpdatablePackageQuantity": {"zh": "可更新包数量", "en": "Updatable Package Quantity"},
}
