package rose

import (
	"fmt"
	"math"
	"time"
)

const (
	timeLayoutShortMonth = "200601"
	timeLayoutDate       = "2006-01-02"
	timeLayoutDateMH     = "2006-01-02 15:04"
	timeLayoutDateTime   = "2006-01-02 15:04:05"
)

func TNow() time.Time {
	return time.Now()
}

// TNowS 当前时间戳（秒 10位
func TNowS() int64 {
	return time.Now().Unix()
}

// TNowStr 当前时间戳（秒 10位 字符串形式
func TNowStr() string {
	return Int64ToStr(TNowS())
}

// TNowMs 当前时间戳（毫秒 13位
func TNowMs() int64 {
	//这种计算毫秒时间戳的方法比较推荐，参考自：https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds
	//return time.Now().UnixNano() / int64(time.Millisecond)

	return time.Now().UnixMilli()
}

// TNowMStr 当前时间戳（毫秒 13位 字符串形式
func TNowMStr() string {
	return Int64ToStr(TNowMs())
}

// TUnixSToTime Converts Unix Epoch from seconds to time.Time
func TUnixSToTime(s int64) time.Time {
	return time.Unix(s, 0)
}

func TUnixMsToTime(ms int64) time.Time {
	return time.UnixMilli(ms)
}

// TODO 这两个可以结合一下

// ToDayRemainSec 截止到今日的24点之前剩余的秒数
func ToDayRemainSec() int64 {
	now := time.Now()
	t, _ := time.ParseInLocation(timeLayoutDate, now.AddDate(0, 0, 1).Format(timeLayoutDate), time.Local)
	return t.Unix() - now.Unix()
}

// DelayTimeToTomorrow 获取当前时间戳截止到明天早晨1点之间的总秒数 (1,"01:00:00")
func DelayTimeToTomorrow(addDays int, addHourStr string) int64 {
	t := time.Now()
	tm := t.AddDate(0, 0, addDays)

	newTimeStr := fmt.Sprintf("%s %s", tm.Format("2006-01-02"), addHourStr) //格式：2006-01-02 15:04:05
	nt, _ := time.ParseInLocation("2006-01-02 15:04:05", newTimeStr, time.Local)

	dt := nt.Sub(t).Seconds()
	fdt := math.Floor(dt + 0.5) // 通过+0.5来实现
	return int64(fdt)
}

func GetMonth() string {
	return time.Now().Format(timeLayoutShortMonth)
}

func GetNextMonth() string {
	return time.Now().AddDate(0, 1, 0).Format(timeLayoutShortMonth)
}

// GetDate 返回今天日期 2019-01-09
func GetDate() string {
	return time.Now().Format(timeLayoutDate)
}

// 根据时间戳返回日期 2019-04-17
func GetDateFromUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(timeLayoutDate)
}

// eg: 1595225361 => 2020-07-20 14:09:21
func GetTimeFromUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(timeLayoutDateTime)
}

func GetTFromUnix(t int64) time.Time {
	return time.Unix(t, 0)
}

// 根据时间戳返回指定格式的时间信息
func GetTimeFromUnixFormat(t int64, format string) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(format)
}

// eg: 1595225361 => 2020-07-20 14:09
func GetTimeMHFromUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(timeLayoutDateMH) //"2006-01-02 15:04"
}

// ************

func GetTimeParse(times string) int64 {
	if "" == times {
		return 0
	}
	parse, _ := time.ParseInLocation("2006-01-02 15:04", times, time.Local)
	return parse.Unix()
}

func GetDateParse(dates string) int64 {
	if "" == dates {
		return 0
	}
	parse, _ := time.ParseInLocation("2006-01-02", dates, time.Local)
	return parse.Unix()
}

//
func StrDateMH2Time(dates string) time.Time {
	return time.Unix(GetDateParse(dates), 0)
}

func MonthStart() time.Time {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

func TodayStart() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func TodayEnd() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 23, 59, 59, 1e9-1, time.Local)
}

func NowDate() string {
	return time.Now().Format(timeLayoutDate)
}

func NowDateTime() string {
	return time.Now().Format(timeLayoutDateTime)
}

func ParseDate(dt string) (time.Time, error) {
	return time.Parse(timeLayoutDate, dt)
}

func ParseDateTime(dt string) (time.Time, error) {
	return time.Parse(timeLayoutDateTime, dt)
}

func ParseStringTime(tm, lc string) (time.Time, error) {
	loc, err := time.LoadLocation(lc)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(timeLayoutDateTime, tm, loc)
}

// ParseGMTTimeOfRFC1123 GMT
// eg: Mon, 20 Jul 2020 06:09:21 GMT => Time
// https://golang.org/pkg/time/#pkg-constants
func ParseGMTTimeOfRFC1123(gmt string) (time.Time, error) {
	return time.Parse(time.RFC1123, gmt)
}

// FormatSec 将秒转换成时分秒形式
// 简写：00:40
// 完整：47:55:49
// isAll 是否显示完整格式
func FormatSec(sec int64, isAll bool) string {

	rHour := math.Floor(float64(sec / 3600.0))
	tmpMin := math.Floor(float64(sec % 3600.0))
	rMin := math.Floor(tmpMin / 60.0)
	rSec := float64(sec % 60.0)

	if isAll {
		return fmt.Sprintf("%02d:%02d:%02d", int(rHour), int(rMin), int(rSec))
	} else {
		if sec >= 0.0 && sec < 3600.0 {
			return fmt.Sprintf("%02d:%02d", int(rMin), int(rSec))
		} else {
			return fmt.Sprintf("%02d:%02d:%02d", int(rHour), int(rMin), int(rSec))
		}
	}
}

// ----------

func TNowAdd(d time.Duration) time.Time {
	return TNow().Add(d)
}

func TNowAddDate(years int, months int, days int) time.Time {
	return TNow().AddDate(years, months, days)
}

func TNowAddYears(years int) time.Time {
	return TNowAddDate(years, 0, 0)
}

func TNowAddMonths(months int) time.Time {
	return TNowAddDate(0, months, 0)
}

func TNowAddDays(days int) time.Time {
	return TNowAddDate(0, 0, days)
}

func TTimes(t time.Time) int64 {
	return t.Unix()
}

func TTimeMs(t time.Time) int64 {
	return t.UnixMilli()
}
