//go:build darwin

/*
File: define_system_darwin.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-04 11:03:33

Description: 系统相关方法
*/

package general

// 版本范围
type versionRange struct {
	Lower string // 最初版本
	Upper string // 最后版本
}

// 系统版本范围和系统代号的对应关系，右包含
// 数据来源：https://support.apple.com/zh-cn/109033
var versionCode = map[versionRange]string{
	{Lower: "0.0.0", Upper: "10.0.4"}:    "Mac OS X Cheetah",
	{Lower: "10.0.4", Upper: "10.1.5"}:   "Mac OS X Puma",
	{Lower: "10.1.5", Upper: "10.2.8"}:   "Mac OS X Jaguar",
	{Lower: "10.2.8", Upper: "10.3.9"}:   "Mac OS X Panther",
	{Lower: "10.3.9", Upper: "10.4.11"}:  "Mac OS X Tiger",
	{Lower: "10.4.11", Upper: "10.5.8"}:  "Mac OS X Leopard",
	{Lower: "10.5.8", Upper: "10.6.8"}:   "Mac OS X Snow Leopard",
	{Lower: "10.6.8", Upper: "10.7.5"}:   "OS X Lion",
	{Lower: "10.7.5", Upper: "10.8.5"}:   "OS X Mountain Lion",
	{Lower: "10.8.5", Upper: "10.9.5"}:   "OS X Mavericks",
	{Lower: "10.9.5", Upper: "10.10.5"}:  "OS X Yosemite",
	{Lower: "10.10.5", Upper: "10.11.6"}: "OS X El Capitan",
	{Lower: "10.11.6", Upper: "10.12.6"}: "macOS Sierra",
	{Lower: "10.12.6", Upper: "10.13.6"}: "macOS High Sierra",
	{Lower: "10.13.6", Upper: "10.14.6"}: "macOS Mojave",
	{Lower: "10.14.6", Upper: "10.15.7"}: "macOS Catalina",
	{Lower: "10.15.7", Upper: "11.7.10"}: "macOS Big Sur",
	{Lower: "11.7.10", Upper: "12.7.5"}:  "macOS Monterey",
	{Lower: "12.7.5", Upper: "13.6.7"}:   "macOS Ventura",
	{Lower: "13.6.7", Upper: "14.5"}:     "macOS Sonoma",
}

// inVersionRange 检查指定的版本是否在指定的版本范围内
func inVersionRange(version string, vRange versionRange) bool {
	return version > vRange.Lower && version <= vRange.Upper
}

// FindSystemCode 根据系统版本查找其对应的系统代号
func FindSystemCode(version string) string {
	for vRange, code := range versionCode {
		if inVersionRange(version, vRange) {
			return code
		}
	}
	return ""
}
