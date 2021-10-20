/**
 * @Author:      leafney
 * @Date:        2021-09-07 11:12
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package time

import "time"

// 当前时间戳（秒 10位
func NowS() int64 {
	return time.Now().Unix()
}

// 当前时间戳（毫秒 13位
func NowMs() int64 {
	//这种计算毫秒时间戳的方法比较推荐，参考自：https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func AddDateUnix(year, month, day int) int64 {
	return time.Now().AddDate(year, month, day).Unix()
}

func GetMonthByUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(layoutMonth)
}

// Return date based on timestamp
func GetDateByUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(layoutDate)
}

// eg: 1595225361 => 2020-07-20 14:09:21
func GetTimeByUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(layoutDateTime)
}

// 根据时间戳返回指定格式的时间信息
func GetTimeByUnixFormat(t int64, format string) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(format)
}

// eg: 1595225361 => 2020-07-20 14:09
func GetTimeMHByUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(layoutDateMH) //"2006-01-02 15:04"
}
