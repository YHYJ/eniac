/*
File: define_variable.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-04-18 13:46:00

Description: 操作变量
*/

package general

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

// ---------- 代码变量

var (
	FgBlackText        = color.FgBlack.Render        // 前景色 - 黑色
	FgWhiteText        = color.FgWhite.Render        // 前景色 - 白色
	FgLightWhiteText   = color.FgLightWhite.Render   // 前景色 - 亮白色
	FgGrayText         = color.FgGray.Render         // 前景色 - 灰色
	FgRedText          = color.FgRed.Render          // 前景色 - 红色
	FgLightRedText     = color.FgLightRed.Render     // 前景色 - 亮红色
	FgGreenText        = color.FgGreen.Render        // 前景色 - 绿色
	FgLightGreenText   = color.FgLightGreen.Render   // 前景色 - 亮绿色
	FgYellowText       = color.FgYellow.Render       // 前景色 - 黄色
	FgLightYellowText  = color.FgLightYellow.Render  // 前景色 - 亮黄色
	FgBlueText         = color.FgBlue.Render         // 前景色 - 蓝色
	FgLightBlueText    = color.FgLightBlue.Render    // 前景色 - 亮蓝色
	FgMagentaText      = color.FgMagenta.Render      // 前景色 - 品红
	FgLightMagentaText = color.FgLightMagenta.Render // 前景色 - 亮品红
	FgCyanText         = color.FgCyan.Render         // 前景色 - 青色
	FgLightCyanText    = color.FgLightCyan.Render    // 前景色 - 亮青色

	BgBlackText        = color.BgBlack.Render        // 背景色 - 黑色
	BgWhiteText        = color.BgWhite.Render        // 背景色 - 白色
	BgLightWhiteText   = color.BgLightWhite.Render   // 背景色 - 亮白色
	BgGrayText         = color.BgGray.Render         // 背景色 - 灰色
	BgRedText          = color.BgRed.Render          // 背景色 - 红色
	BgLightRedText     = color.BgLightRed.Render     // 背景色 - 亮红色
	BgGreenText        = color.BgGreen.Render        // 背景色 - 绿色
	BgLightGreenText   = color.BgLightGreen.Render   // 背景色 - 亮绿色
	BgYellowText       = color.BgYellow.Render       // 背景色 - 黄色
	BgLightYellowText  = color.BgLightYellow.Render  // 背景色 - 亮黄色
	BgBlueText         = color.BgBlue.Render         // 背景色 - 蓝色
	BgLightBlueText    = color.BgLightBlue.Render    // 背景色 - 亮蓝色
	BgMagentaText      = color.BgMagenta.Render      // 背景色 - 品红
	BgLightMagentaText = color.BgLightMagenta.Render // 背景色 - 亮品红
	BgCyanText         = color.BgCyan.Render         // 背景色 - 青色
	BgLightCyanText    = color.BgLightCyan.Render    // 背景色 - 亮青色

	InfoText      = color.Info.Render      // Info 文本
	NoteText      = color.Note.Render      // Note 文本
	LightText     = color.Light.Render     // Light 文本
	ErrorText     = color.Error.Render     // Error 文本
	DangerText    = color.Danger.Render    // Danger 文本
	NoticeText    = color.Notice.Render    // Notice 文本
	SuccessText   = color.Success.Render   // Success 文本
	CommentText   = color.Comment.Render   // Comment 文本
	PrimaryText   = color.Primary.Render   // Primary 文本
	WarnText      = color.Warn.Render      // Warn 文本
	QuestionText  = color.Question.Render  // Question 文本
	SecondaryText = color.Secondary.Render // Secondary 文本
)

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
	"Load":    {"zh": "负载", "en": "Load"},
	"Time":    {"zh": "时间", "en": "Time"},
	"User":    {"zh": "用户", "en": "User"},
	"Update":  {"zh": "更新", "en": "Update"},
}

// 各部分.条目的名称
var GenealogyName = map[string]map[string]string{
	"BIOSVendor":         {"zh": "BIOS 厂商", "en": "Vendor"},
	"BIOSVersion":        {"zh": "BIOS 版本", "en": "Version"},
	"BIOSDate":           {"zh": "BIOS 发布", "en": "Date"},
	"BoardVendor":        {"zh": "主板厂商", "en": "Vendor"},
	"BoardName":          {"zh": "主板名称", "en": "Name"},
	"BoardVersion":       {"zh": "主板版本", "en": "Version"},
	"CPUModel":           {"zh": "处理器型号", "en": "Model"},
	"CPUNumber":          {"zh": "处理器数量", "en": "Number"},
	"CPUCores":           {"zh": "处理器核心", "en": "Cores"},
	"CPUThreads":         {"zh": "处理器线程", "en": "Threads"},
	"CPUCache":           {"zh": "处理器缓存", "en": "Cache"},
	"GPUAddress":         {"zh": "显卡地址", "en": "Address"},
	"GPUDriver":          {"zh": "显卡驱动", "en": "Driver"},
	"GPUProduct":         {"zh": "显卡型号", "en": "Product"},
	"GPUVendor":          {"zh": "显卡厂商", "en": "Vendor"},
	"OS":                 {"zh": "操作系统", "en": "OS"},
	"Arch":               {"zh": "系统架构", "en": "Arch"},
	"Kernel":             {"zh": "内核版本", "en": "Kernel"},
	"Platform":           {"zh": "系统类型", "en": "Platform"},
	"Hostname":           {"zh": "主机名称", "en": "Hostname"},
	"TimeZone":           {"zh": "时区", "en": "Time zone"},
	"Load1":              {"zh": "1分钟平均负载", "en": "Load average (1 min)"},
	"Load5":              {"zh": "5分钟平均负载", "en": "Load average (5 min)"},
	"Load15":             {"zh": "15分钟平均负载", "en": "Load average (15 min)"},
	"NicName":            {"zh": "网卡名称", "en": "Name"},
	"NicPCIAddress":      {"zh": "PCI 地址", "en": "PCI Address"},
	"NicMacAddress":      {"zh": "MAC 地址", "en": "MAC Address"},
	"NicSpeed":           {"zh": "网卡速率", "en": "Speed"},
	"NicDuplex":          {"zh": "工作模式", "en": "Duplex"},
	"NicDriver":          {"zh": "网卡驱动", "en": "Driver"},
	"NicProduct":         {"zh": "网卡型号", "en": "Product"},
	"NicVendor":          {"zh": "网卡厂商", "en": "Vendor"},
	"MemoryTotal":        {"zh": "内存大小", "en": "Total"},
	"MemoryUsed":         {"zh": "已用内存", "en": "Used"},
	"MemoryUsedPercent":  {"zh": "内存占用", "en": "Used Percent"},
	"MemoryFree":         {"zh": "空闲内存", "en": "Free"},
	"MemoryShared":       {"zh": "共享内存", "en": "Shared"},
	"MemoryBuffCache":    {"zh": "缓冲内存", "en": "Buff Cache"},
	"MemoryAvail":        {"zh": "可用内存", "en": "Avail"},
	"SwapStatus":         {"zh": "交换空间状态", "en": "Swap Status"},
	"SwapTotal":          {"zh": "交换空间大小", "en": "Total"},
	"SwapFree":           {"zh": "空闲交换空间", "en": "Free"},
	"Process":            {"zh": "进程数", "en": "Process"},
	"ProductVendor":      {"zh": "设备厂商", "en": "Vendor"},
	"ProductName":        {"zh": "设备名称", "en": "Name"},
	"StorageName":        {"zh": "磁盘名称", "en": "Name"},
	"StorageType":        {"zh": "磁盘类型", "en": "Type"},
	"StorageDriver":      {"zh": "磁盘驱动", "en": "Driver"},
	"StorageVendor":      {"zh": "磁盘厂商", "en": "Vendor"},
	"StorageModel":       {"zh": "磁盘型号", "en": "Model"},
	"StorageSerial":      {"zh": "磁盘序列号", "en": "Serial"},
	"StorageRemovable":   {"zh": "磁盘可移除", "en": "Removable"},
	"StorageSize":        {"zh": "磁盘容量", "en": "Size"},
	"BootTime":           {"zh": "系统启动时间", "en": "Boot Time"},
	"Uptime":             {"zh": "系统运行时长", "en": "Uptime"},
	"StartTime":          {"zh": "系统启动用时", "en": "Startup Time"},
	"User":               {"zh": "用户名称", "en": "User"},
	"UserName":           {"zh": "用户昵称", "en": "Username"},
	"UserUid":            {"zh": "用户标识", "en": "UID"},
	"UserGid":            {"zh": "用户组标识", "en": "GID"},
	"UserHomeDir":        {"zh": "用户主目录", "en": "Home Dir"},
	"UpdateList":         {"zh": "更新列表", "en": "Update List"},
	"UpdateDaemonStatus": {"zh": "更新服务", "en": "Update Daemon"},
}

// ---------- 环境变量

// 用来处理不同系统之间的变量名差异
var platformChart = map[string]map[string]string{
	"windows": {
		"HOME":     "USERPROFILE",  // 用户主目录路径
		"USER":     "USERNAME",     // 当前登录用户名
		"SHELL":    "ComSpec",      // 默认 shell 或命令提示符路径
		"PWD":      "CD",           // 当前工作目录路径
		"HOSTNAME": "COMPUTERNAME", // 计算机主机名
	},
}

// 用户名，当程序提权运行时，使用 SUDO_USER 变量获取提权前的用户名
var UserName = func() string {
	if GetVariable("SUDO_USER") != "" {
		return GetVariable("SUDO_USER")
	}
	return GetVariable("USER")
}()

var Platform = runtime.GOOS                   // 操作系统
var Arch = runtime.GOARCH                     // 系统架构
var Sep = string(filepath.Separator)          // 路径分隔符
var UserInfo, _ = GetUserInfoByName(UserName) // 用户信息
var Language = GetLanguage()                  // 系统语言

var (
	programDir = strings.ToLower(Name)                      // 程序目录
	configDir  = filepath.Join(UserInfo.HomeDir, ".config") // 配置目录
	configFile = "config.toml"                              // 配置文件

	ConfigFile = filepath.Join(configDir, programDir, configFile) // 配置文件路径
)

// ---------- 函数

// GetVariable 获取环境变量
//
// 参数：
//   - key: 变量名
//
// 返回：
//   - 变量值
func GetVariable(key string) string {
	if innerMap, exists := platformChart[Platform]; exists {
		if _, variableExists := innerMap[key]; variableExists {
			key = platformChart[Platform][key]
		}
	}
	variable := os.Getenv(key)

	return variable
}

// GetLanguage 获取系统语言
//
// 返回:
//   - 系统语言，目前仅支持 zh 或 en
func GetLanguage() string {
	language := GetVariable("LANGUAGE")
	if strings.Contains(language, "zh") {
		return "zh"
	} else {
		return "en"

	}
}

// GetHostname 获取系统 HOSTNAME
//
// 返回：
//   - HOSTNAME 或空字符串
func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	return hostname
}

// SetVariable 设置环境变量
//
// 参数：
//   - key: 变量名
//   - value: 变量值
//
// 返回：
//   - 错误信息
func SetVariable(key, value string) error {
	return os.Setenv(key, value)
}

// GetUserInfoByName 根据用户名获取用户信息
//
// 参数：
//   - userName: 用户名
//
// 返回：
//   - 用户信息
//   - 错误信息
func GetUserInfoByName(userName string) (*user.User, error) {
	userInfo, err := user.Lookup(userName)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

// GetUserInfoById 根据 ID 获取用户信息
//
// 参数：
//   - userId: 用户 ID
//
// 返回：
//   - 用户信息
//   - 错误信息
func GetUserInfoById(userId int) (*user.User, error) {
	userInfo, err := user.LookupId(strconv.Itoa(userId))
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

// GetCurrentUserInfo 获取当前用户信息
//
// 返回：
//   - 用户信息
//   - 错误信息
func GetCurrentUserInfo() (*user.User, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}
	return currentUser, nil
}
