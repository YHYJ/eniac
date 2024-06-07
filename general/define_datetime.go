/*
File: define_datetime.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-07 14:39:41

Description: 关于日期和时间的方法
*/

package general

import "path"

// GetTimeZoneOriginal 获取时区信息的原始方法（检测 /etc/localtime 实际指向的文件）
//
// 返回：
//   - 时区信息
func GetTimeZoneOriginal() string {
	var localtimeFile = "/etc/localtime"
	var timeZone string

	filePath, err := ReadFileLink(localtimeFile)
	if err != nil {
		timeZone = ""
	}

	dir, file := path.Split(filePath)
	if len(dir) == 0 || len(file) == 0 {
		timeZone = ""
	}

	_, fname := path.Split(dir[:len(dir)-1])
	if fname == "zoneinfo" {
		timeZone = file
	} else {
		timeZone = path.Join(fname, file)
	}

	return timeZone
}
