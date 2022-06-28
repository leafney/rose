package xtime

import (
	"fmt"
	"math"
	"time"
)

// 截止到今日的24点之前的秒数
func ToDayRemainSec() int64 {
	now := time.Now()
	t, _ := time.ParseInLocation(layoutDate, now.AddDate(0, 0, 1).Format(layoutDate), time.Local)
	return t.Unix() - now.Unix()
}

/*
获取当前时间戳截止到明天早晨1点之间的总秒数 (1,"01:00:00")
*/
func DelayTimeToTomorrow(addDays int, addHourStr string) int64 {
	t := time.Now()
	tm := t.AddDate(0, 0, addDays)

	newTimeStr := fmt.Sprintf("%s %s", tm.Format(layoutDate), addHourStr) //格式：2006-01-02 15:04:05
	nt, _ := time.ParseInLocation(layoutDateTime, newTimeStr, time.Local)

	dt := nt.Sub(t).Seconds()
	fdt := math.Floor(dt + 0.5) // 通过+0.5来实现
	return int64(fdt)
}

// 202110
func GetMonth() string {
	return time.Now().Format(layoutMonth)
}

// last month
func GetMonthLast() string {
	return time.Now().AddDate(0, -1, 0).Format(layoutMonth)
}

// next month
func GetMonthNext() string {
	return time.Now().AddDate(0, 1, 0).Format(layoutMonth)
}

// 返回今天日期 2019-01-09
func GetDate() string {
	return time.Now().Format(layoutDate)
}

// ************

func GetTimeParse(times string) int64 {
	if times == "" {
		return 0
	}
	parse, _ := time.ParseInLocation(layoutDateTime, times, time.Local)
	return parse.Unix()
}

func GetDateParse(dates string) int64 {
	if "" == dates {
		return 0
	}
	parse, _ := time.ParseInLocation(layoutDate, dates, time.Local)
	return parse.Unix()
}

//
func StrDateMH2Time(dates string) time.Time {
	return time.Unix(GetDateParse(dates), 0)
}

func MonthStart() time.Time {
	now := time.Now()
	y, m, _ := now.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
}

func MonthEnd() time.Time {
	now := time.Now()
	y, m, _ := now.Date()
	return time.Date(y, m+1, 1, 0, 0, 0, -1, now.Location())
}

func TodayStart() time.Time {
	now := time.Now()
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location())
}

func TodayEnd() time.Time {
	now := time.Now()
	y, m, d := now.Date()
	return time.Date(y, m, d, 23, 59, 59, 1e9-1, now.Location())
}

func GetDateTime() string {
	return time.Now().Format(layoutDateTime)
}

func ParseDate(dt string) (time.Time, error) {
	return time.Parse(layoutDate, dt)
}

func ParseDateTime(dt string) (time.Time, error) {
	return time.Parse(layoutDateTime, dt)
}

func ParseStringTime(tm, lc string) (time.Time, error) {
	loc, err := time.LoadLocation(lc)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(layoutDateTime, tm, loc)
}

// GMT
// eg: Mon, 20 Jul 2020 06:09:21 GMT =>
// https://golang.org/pkg/time/#pkg-constants
func ParseGMTTimeOfRFC1123(gmt string) (time.Time, error) {
	return time.Parse(time.RFC1123, gmt)
}
