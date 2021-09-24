package time

import (
	"fmt"
	"math"
	"time"
)

var (
	TimeLayoutDate     = "2006-01-02"
	TimeLayoutDateMH   = "2006-01-02 15:04"
	TimeLayoutDateTime = "2006-01-02 15:04:05"
)

// 当前时间戳
func TimeNow10() int64 {
	return time.Now().Unix()
}

// 当前时间戳
func TimeNow13() int64 {
	//这种计算毫秒时间戳的方法比较推荐，参考自：https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// TODO 这两个可以结合一下

// 截止到今日的24点之前的秒数
func TheDayExpireSeconds() int64 {
	now := time.Now()
	t, _ := time.ParseInLocation("2006-01-02", now.AddDate(0, 0, 1).Format("2006-01-02"), time.Local)
	return t.Unix() - now.Unix()
}

/*
获取当前时间戳截止到明天早晨1点之间的总秒数 (1,"01:00:00")
*/
func DelayTimeToTomorrow(addDays int, addHourStr string) int64 {
	t := time.Now()
	tm := t.AddDate(0, 0, addDays)

	newTimeStr := fmt.Sprintf("%s %s", tm.Format("2006-01-02"), addHourStr) //格式：2006-01-02 15:04:05
	nt, _ := time.ParseInLocation("2006-01-02 15:04:05", newTimeStr, time.Local)

	dt := nt.Sub(t).Seconds()
	fdt := math.Floor(dt + 0.5) // 通过+0.5来实现
	return int64(fdt)
}

// 返回今天日期 2019-01-09
func GetDate() string {
	return time.Now().Format(TimeLayoutDate)
}

// 根据时间戳返回日期 2019-04-17
func GetDateFromUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(TimeLayoutDate)
}

// eg: 1595225361 => 2020-07-20 14:09:21
func GetTimeFromUnix(t int64) string {
	if t <= 0 {
		return ""
	}
	return time.Unix(t, 0).Format(TimeLayoutDateTime)
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
	return time.Unix(t, 0).Format(TimeLayoutDateMH) //"2006-01-02 15:04"
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
	return time.Now().Format(TimeLayoutDate)
}

func NowDateTime() string {
	return time.Now().Format(TimeLayoutDateTime)
}

func ParseDate(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDate, dt)
}

func ParseDateTime(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDateTime, dt)
}

func ParseStringTime(tm, lc string) (time.Time, error) {
	loc, err := time.LoadLocation(lc)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(TimeLayoutDateTime, tm, loc)
}

// GMT
// eg: Mon, 20 Jul 2020 06:09:21 GMT =>
// https://golang.org/pkg/time/#pkg-constants
func ParseGMTTimeOfRFC1123(gmt string) (time.Time, error) {
	return time.Parse(time.RFC1123, gmt)
}
