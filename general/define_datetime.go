/*
File: define_datetime.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-06-07 14:39:41

Description: 处理日期/时间
*/

package general

import (
	"path"
	"time"
)

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

// UnixTime2DayHourMinuteSecond Unix 时间戳转换为天、小时、分钟、秒
//
// 参数：
//   - totalSeconds: Unix 时间戳
//
// 返回：
//   - day: 天
//   - hour: 小时
//   - minute: 分钟
//   - second: 秒
func UnixTime2DayHourMinuteSecond(unixTime int64) (day, hour, minute, second int64) {
	day = unixTime / 86400
	hour = (unixTime - day*86400) / 3600
	minute = (unixTime - day*86400 - hour*3600) / 60
	second = unixTime - day*86400 - hour*3600 - minute*60
	return day, hour, minute, second
}

// UnixTime2TimeString Unix 时间戳转换为字符串格式
//
// 参数：
//   - timeStamp: Unix 时间戳
//
// 返回：
//   - 格式化的 Unix 时间戳字符串
func UnixTime2TimeString(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")
}
