/**
 * @Author:      leafney
 * @Date:        2021-09-07 11:12
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description: unix时间戳相关
 */

package xtime

import "time"

func Now() time.Time {
	return time.Now()
}

// 当前时间戳（秒 10位
func NowS() int64 {
	return Now().Unix()
}

// 当前时间戳（毫秒 13位
func NowMs() int64 {
	//这种计算毫秒时间戳的方法比较推荐，参考自：https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds
	// [time: add Time.UnixMilli and Time.UnixMicro (like Time.UnixNano) · Issue #44196 · golang/go](https://github.com/golang/go/issues/44196)
	// time.Now().UnixNano() / int64(time.Millisecond) 官方增加了新方法支持获取毫秒
	return Now().UnixMilli()
}

// Converts Unix Epoch from seconds to time.Time
func UnixSToTime(s int64) time.Time {
	return time.Unix(s, 0)
}

func UnixMsToTime(ms int64) time.Time {
	return time.UnixMilli(ms)
}

func AddDateUnix(year, month, day int) int64 {
	return Now().AddDate(year, month, day).Unix()
}

func GetMonthByUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return UnixSToTime(t).Format(layoutMonth)
}

// Return date based on timestamp
func GetDateByUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return UnixSToTime(t).Format(layoutDate)
}

// eg: 1595225361 => 2020-07-20 14:09:21
func GetTimeByUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return UnixSToTime(t).Format(layoutDateTime)
}

// 根据时间戳返回指定格式的时间信息
func GetTimeByUnixFormat(t int64, format string) string {
	if t <= 0 {
		return ""
	}
	return UnixSToTime(t).Format(format)
}

// eg: 1595225361 => 2020-07-20 14:09
func GetTimeMHByUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return UnixSToTime(t).Format(layoutDateMH) //"2006-01-02 15:04"
}
