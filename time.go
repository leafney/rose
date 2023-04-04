package rose

import (
	"fmt"
	"math"
	"time"
)

const (
	timeLayoutShortY   = "2006"
	timeLayoutShortYM  = "200601"
	timeLayoutShortYMD = "20060102"
	timeLayoutLongYM   = "2006-01"
	timeLayoutDateYMD  = "2006-01-02"
	timeLayoutDateYMDH = "2006-01-02 15:04"
	timeLayoutDateTime = "2006-01-02 15:04:05"
	timeLayoutTimeH    = "15:04"
	timeLayoutTimeHS   = "15:04:05"
)

type TFormat string

const (
	TFShortY   TFormat = "2006"
	TFShortYM  TFormat = "200601"
	TFShortYMD TFormat = "20060102"
	TFLongYM   TFormat = "2006-01"
	TFDateYMD  TFormat = "2006-01-02"
	TFDateYMDH TFormat = "2006-01-02 15:04"
	TFDateTime TFormat = "2006-01-02 15:04:05"
)

// ----------------------------

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

// TNowDateMD eg: 2023-03-23
func TNowDateMD() string {
	return time.Now().Format(timeLayoutDateYMD)
}

// TNowDateMDH eg: 2023-03-23 23:07
func TNowDateMDH() string {
	return time.Now().Format(timeLayoutDateYMDH)
}

// TNowDateTime eg: 2023-03-23 23:06:42
func TNowDateTime() string {
	return time.Now().Format(timeLayoutDateTime)
}

func TNowFormat(format TFormat) string {
	return time.Now().Format(string(format))
}

// ----------------------------

// TUnixSToTime Converts Unix Epoch from seconds to time.Time
func TUnixSToTime(s int64) time.Time {
	return time.Unix(s, 0)
}

func TUnixMsToTime(ms int64) time.Time {
	return time.UnixMilli(ms)
}

// TUnixSToDateMD 根据时间戳返回日期 eg: 2019-04-17
func TUnixSToDateMD(s int64) string {
	if s <= 0 {
		return ""
	}
	return time.Unix(s, 0).Format(timeLayoutDateYMD)
}

// TUnixSToDateMDH eg: 1595225361 => 2020-07-20 14:09
func TUnixSToDateMDH(s int64) string {
	if s <= 0 {
		return ""
	}
	return time.Unix(s, 0).Format(timeLayoutDateYMDH) //"2006-01-02 15:04"
}

// TUnixSToDateTime eg: 1595225361 => 2020-07-20 14:09:21
func TUnixSToDateTime(s int64) string {
	if s <= 0 {
		return ""
	}
	return time.Unix(s, 0).Format(timeLayoutDateTime)
}

// TUnixSToFormat 将时间戳转换为指定格式的时间字符串
func TUnixSToFormat(s int64, format TFormat) string {
	if s <= 0 {
		return ""
	}
	return time.Unix(s, 0).Format(string(format))
}

// ----------------------------

// TMonth 获取当前月份 eg: 202303
func TMonth() string {
	return time.Now().Format(timeLayoutShortYM)
}

// TNextMonth 获取下月月份 eg：202304
func TNextMonth() string {
	return time.Now().AddDate(0, 1, 0).Format(timeLayoutShortYM)
}

// TDate 获取当前月份日期 eg: 2019-01-09
func TDate() string {
	return time.Now().Format(timeLayoutDateYMD)
}

// TNextDate 获取下月月份日期 eg: 2019-02-09
func TNextDate() string {
	return time.Now().AddDate(0, 1, 0).Format(timeLayoutDateYMD)
}

// ----------------------------

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

// ----------------------------

func TStartTheMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

func TStartTheDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func TEndTheDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}

// TEndTheDayRemainSec 截止到今日的24点之前剩余的秒数
func TEndTheDayRemainSec() int64 {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return int64(endOfDay.Sub(now).Seconds())
}

// ----------------------------

func ParseDate(dt string) (time.Time, error) {
	return time.Parse(timeLayoutDateYMD, dt)
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

// ----------------------------

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

// TSecByYear one year for 365 days
func TSecByYear(year int) int64 {
	duration := time.Duration(year) * 365 * 24 * time.Hour
	return int64(duration.Seconds())
}

// TSecByMonth one month for 30 days
func TSecByMonth(month int) int64 {
	duration := time.Duration(month) * 30 * 24 * time.Hour
	return int64(duration.Seconds())
}

func TSecByDay(day int) int64 {
	duration := time.Duration(day) * 24 * time.Hour
	return int64(duration.Seconds())
}

func TSecByHour(hour int) int64 {
	duration := time.Duration(hour) * time.Hour
	return int64(duration.Seconds())
}

func TSecByMin(min int) int64 {
	duration := time.Duration(min) * time.Minute
	return int64(duration.Seconds())
}

// TSecNowToAdd 当前时间增加指定的天数、小时数、分钟数，得到未来时间点的时间戳
// 例如：增加一小时，最后的时间戳即为 Now().Unix() + 3600
func TSecNowToAdd(days int, hours int, minutes int) int64 {
	return TSecNowToAddY(0, 0, days, hours, minutes, 0)
}

func TSecNowToAddY(years, months, days int, hours, minutes, sec int) int64 {
	now := time.Now()
	future := now.AddDate(years, months, days).
		Add(time.Duration(hours) * time.Hour).
		Add(time.Duration(minutes) * time.Minute).
		Add(time.Duration(sec) * time.Second).
		Unix()
	return future
}

// TSecUntilNowToAdd 当前时间增加指定的天数、小时数、分钟数，获取当前时间戳截止到未来时间点之间的差值秒数
// 例如：增加一小时，差值即为 3600 秒
func TSecUntilNowToAdd(days int, hours int, minutes int) int64 {
	return TSecUntilNowToAddY(0, 0, days, hours, minutes, 0)
}

func TSecUntilNowToAddY(years, months, days int, hours, minutes, sec int) int64 {
	now := time.Now()
	future := now.AddDate(years, months, days).
		Add(time.Duration(hours) * time.Hour).
		Add(time.Duration(minutes) * time.Minute).
		Add(time.Duration(sec) * time.Second)
	duration := future.Sub(now).Seconds()
	return int64(duration)
}

// ----------------------------

// TODO 待后期完善

// ParseGMTTimeOfRFC1123 GMT
// eg: Mon, 20 Jul 2020 06:09:21 GMT => Time
// https://golang.org/pkg/time/#pkg-constants
func ParseGMTTimeOfRFC1123(gmt string) (time.Time, error) {
	return time.Parse(time.RFC1123, gmt)
}

// TFormatSec 将秒转换成时分秒形式
// 简写：00:40
// 完整：47:55:49
// isAll 是否显示完整格式
func TFormatSec(sec int64, isAll bool) string {

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
