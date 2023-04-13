package rose

import (
	"errors"
	"fmt"
	"math"
	"time"
)

const (
	TFLayoutShortY     = "2006"
	TFLayoutShortYM    = "200601"
	TFLayoutShortYMD   = "20060102"
	TFLayoutShortYMDH  = "200601021504"
	TFLayoutShortYMDHS = "20060102150405"
	TFLayoutLongYM     = "2006-01"
	TFLayoutLongYMD    = "2006-01-02"
	TFLayoutLongYMDH   = "2006-01-02 15:04"
	TFLayoutLongYMDHS  = "2006-01-02 15:04:05"
	TFLayoutTimeH      = "15:04"
	TFLayoutTimeHS     = "15:04:05"
)

const tScale = int(60)

type TFormat string

const (
	TFShortY     TFormat = "2006"
	TFShortYM    TFormat = "200601"
	TFShortYMD   TFormat = "20060102"
	TFShortYMDH  TFormat = "200601021504"
	TFShortYMDHS TFormat = "20060102150405"
	TFLongYM     TFormat = "2006-01"
	TFLongYMD    TFormat = "2006-01-02"
	TFLongYMDH   TFormat = "2006-01-02 15:04"
	TFLongYMDHS  TFormat = "2006-01-02 15:04:05"
	TFTimeH      TFormat = "15:04"
	TFTimeHS     TFormat = "15:04:05"
)

// ----------------------------

// TNow 当前时间
func TNow() time.Time {
	return time.Now()
}

// TNowS 当前时间戳（秒 10位
func TNowS() int64 {
	return time.Now().Unix()
}

// TNowStr 当前时间戳字符串（秒 10位
func TNowStr() string {
	return Int64ToStr(TNowS())
}

// TNowMs 当前时间戳（毫秒 13位
func TNowMs() int64 {
	//这种计算毫秒时间戳的方法比较推荐，参考自：https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds
	//return time.Now().UnixNano() / int64(time.Millisecond)

	return time.Now().UnixMilli()
}

// TNowMStr 当前时间戳字符串（毫秒 13位
func TNowMStr() string {
	return Int64ToStr(TNowMs())
}

// TNowDateSYMD 当前日期字符串，格式：20230323
func TNowDateSYMD() string {
	return time.Now().Format(TFLayoutShortYMD)
}

// TNowDateLYMD 当前日期字符串，格式：2023-03-23
func TNowDateLYMD() string {
	return time.Now().Format(TFLayoutLongYMD)
}

// TNowDateSYMDH 当前日期字符串，格式：202303232307
func TNowDateSYMDH() string {
	return time.Now().Format(TFLayoutShortYMDH)
}

// TNowDateLYMDH 当前日期字符串，格式：2023-03-23 23:07
func TNowDateLYMDH() string {
	return time.Now().Format(TFLayoutLongYMDH)
}

// TNowDateTime 当前日期字符串，格式：2023-03-23 23:06:42
func TNowDateTime() string {
	return time.Now().Format(TFLayoutLongYMDHS)
}

// TNowDateTimeS 当前日期字符串，格式：20230323230642
func TNowDateTimeS() string {
	return time.Now().Format(TFLayoutShortYMDHS)
}

// TNowFormat 当前日期字符串，内置格式
func TNowFormat(layout TFormat) string {
	return time.Now().Format(string(layout))
}

// TNowFormatStr 当前日期字符串，自定义格式
func TNowFormatStr(layout string) string {
	return time.Now().Format(layout)
}

// ----------------------------

// TUnixSToTime 将毫秒时间戳转换为 time.Time
func TUnixSToTime(s int64) time.Time {
	return time.Unix(s, 0)
}

// TUnixMsToTime 将毫秒时间戳转换为time.Time
func TUnixMsToTime(ms int64) time.Time {
	return time.UnixMilli(ms)
}

// TUnixSToDateYMD 将秒时间戳转换为`2006-01-02`格式字符串
// eg: 1595225361 => 2020-07-20
func TUnixSToDateYMD(s int64) string {
	return TUnixSFormat(s, TFLongYMD)
}

// TUnixSToDateYMDH 将秒时间戳转换为`2006-01-02 15:04`格式字符串
// eg: 1595225361 => 2020-07-20 14:09
func TUnixSToDateYMDH(s int64) string {
	return TUnixSFormat(s, TFLongYMDH)
}

// TUnixSToDateTime 将秒时间戳转换为`2006-01-02 15:04:05`格式字符串
// eg: 1595225361 => 2020-07-20 14:09:21
func TUnixSToDateTime(s int64) string {
	return TUnixSFormat(s, TFLongYMDHS)
}

// TUnixSFormat 将秒时间戳转换为指定格式的时间字符串
func TUnixSFormat(s int64, layout TFormat) string {
	return TUnixSFormatStr(s, string(layout))
}

// TUnixSFormatStr 将秒时间戳转换为指定格式的时间字符串
func TUnixSFormatStr(s int64, layout string) string {
	if s <= 0 {
		return ""
	}
	return time.Unix(s, 0).Format(layout)
}

// TUnixMsFormat 将毫秒时间戳转换为指定格式的时间字符串
func TUnixMsFormat(ms int64, layout TFormat) string {
	return TUnixMsFormatStr(ms, string(layout))
}

// TUnixMsFormatStr 将毫秒时间戳转换为指定格式的时间字符串
func TUnixMsFormatStr(ms int64, layout string) string {
	if ms <= 0 {
		return ""
	}
	return time.UnixMilli(ms).Format(layout)
}

// ----------------------------

// TMonth 获取当前月份 eg: 202303
func TMonth() string {
	return time.Now().Format(TFLayoutShortYM)
}

// TNextMonth 获取下月月份 eg：202304
func TNextMonth() string {
	return time.Now().AddDate(0, 1, 0).Format(TFLayoutShortYM)
}

// TDate 获取当前月份日期 eg: 2019-01-09
func TDate() string {
	return time.Now().Format(TFLayoutLongYMD)
}

// TNextDate 获取下月月份日期 eg: 2019-02-09
func TNextDate() string {
	return time.Now().AddDate(0, 1, 0).Format(TFLayoutLongYMD)
}

// ----------------------------

// TParseDateTimeToTime 将时间字符串`2006-01-02 15:04:05`转换为time.Time
func TParseDateTimeToTime(ds string) time.Time {
	return TParseFormatToTimeLoc(TFLongYMDHS, ds, time.Local)
}

// TParseDateYMDToTime 将时间字符串`2006-01-02`转换为time.Time
func TParseDateYMDToTime(ds string) time.Time {
	return TParseFormatToTime(TFLongYMD, ds)
}

// TParseDateYMDHToTime 将时间字符串`2006-01-02 15:04`转换为time.Time
func TParseDateYMDHToTime(ds string) time.Time {
	return TParseFormatToTime(TFLongYMDH, ds)
}

// TParseDateTimeToUnix 将时间字符串`2006-01-02 15:04:05`转换为Unix时间戳
func TParseDateTimeToUnix(ds string) int64 {
	return TParseDateTimeToTime(ds).Unix()
}

// TParseDateYMDToUnix 将时间字符串`2006-01-02`转换为Unix时间戳
func TParseDateYMDToUnix(ds string) int64 {
	return TParseDateYMDToTime(ds).Unix()
}

// TParseDateYMDHToUnix 将时间字符串`2006-01-02 15:04`转换为Unix时间戳
func TParseDateYMDHToUnix(ds string) int64 {
	return TParseDateYMDHToTime(ds).Unix()
}

// TParseFormatToTime 将内置格式时间字符串转换为time.Time
func TParseFormatToTime(layout TFormat, ds string) time.Time {
	return TParseFormatToTimeLoc(layout, ds, time.Local)
}

// TParseFormatToTimeLoc 将内置格式时间字符串转换为time.Time，支持指定时区
func TParseFormatToTimeLoc(layout TFormat, ds string, loc *time.Location) time.Time {
	return TParseFormatStrToTimeLoc(string(layout), ds, loc)
}

// TParseFormatStrToTime 将自定义格式时间字符串转换为time.Time
func TParseFormatStrToTime(layout string, ds string) time.Time {
	return TParseFormatStrToTimeLoc(layout, ds, time.Local)
}

// TParseFormatStrToTimeLoc 将自定义格式时间字符串转换为time.Time，支持指定时区
func TParseFormatStrToTimeLoc(layout string, ds string, loc *time.Location) time.Time {
	if StrIsEmpty(ds) {
		return TNow()
	}
	parse, _ := time.ParseInLocation(layout, ds, loc)
	return parse
}

// TParseFormatToTimeE 将内置格式时间字符串转换为time.Time，支持异常捕获
func TParseFormatToTimeE(layout TFormat, ds string) (time.Time, error) {
	return TParseEFormatStrToTimeLocE(string(layout), ds, time.Local)
}

// TParseEFormatStrToTimeE 将自定义格式时间字符串转换为time.Time，支持异常捕获
func TParseEFormatStrToTimeE(layout string, ds string) (time.Time, error) {
	return TParseEFormatStrToTimeLocE(layout, ds, time.Local)
}

// TParseEFormatStrToTimeLocE 将自定义格式时间字符串转换为time.Time，支持指定时区，支持异常捕获
func TParseEFormatStrToTimeLocE(layout string, ds string, loc *time.Location) (time.Time, error) {
	if StrIsEmpty(ds) {
		return time.Time{}, errors.New("ds value empty")
	}
	return time.ParseInLocation(layout, ds, loc)
}

// TParseGMTTimeOfRFC1123 GMT
// eg: Mon, 20 Jul 2020 06:09:21 GMT => Time
// https://golang.org/pkg/time/#pkg-constants
func TParseGMTTimeOfRFC1123(gmt string) (time.Time, error) {
	return time.Parse(time.RFC1123, gmt)
}

// ----------------------------

// TDayStartTime 今天的起始时间
func TDayStartTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// TDayStartUnixS 今天的起始时间戳秒
func TDayStartUnixS() int64 {
	return TDayStartTime().Unix()
}

// TDayEndTime 今天的结束时间
func TDayEndTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}

// TDayEndUnixS 今天的结束时间戳秒
func TDayEndUnixS() int64 {
	return TDayEndTime().Unix()
}

// TDayEndRemainSec 截止到今日的24点之前剩余的秒数
func TDayEndRemainSec() int64 {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return int64(endOfDay.Sub(now).Seconds())
}

// TMonthStartTime 当前月份的第一天日期
func TMonthStartTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// TMonthStartStr 当前月份的第一天日期字符串
func TMonthStartStr() string {
	return TTimeFormat(TMonthStartTime(), TFLongYMD)
}

// TMonthEndTime 当前月份的最后一天日期
func TMonthEndTime() time.Time {
	now := time.Now()
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	return nextMonth.AddDate(0, 0, -1)
}

// TMonthEndStr 当前月份的最后一天日期字符串
func TMonthEndStr() string {
	return TTimeFormat(TMonthEndTime(), TFLongYMD)
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

// TNowAddUnixS 当前时间增加指定的年、月、日、天、小时、分钟，秒 后，得到未来时间点的时间戳
func TNowAddUnixS(years, months, days, hours, minutes, secs int) int64 {
	return TNowAddDateTime(years, months, days, hours, minutes, secs).Unix()
}

// TNowAddUnixSDHM 当前时间增加指定的天数、小时数、分钟数，得到未来时间点的时间戳
func TNowAddUnixSDHM(days int, hours int, minutes int) int64 {
	return TNowAddUnixS(0, 0, days, hours, minutes, 0)
}

// TNowAddDateTime 当前时间增加指定的年、月、日、天、小时、分钟，秒 后，得到未来的时间点
func TNowAddDateTime(years, months, days, hours, minutes, secs int) time.Time {
	return TNow().AddDate(years, months, days).
		Add(time.Duration(hours) * time.Hour).
		Add(time.Duration(minutes) * time.Minute).
		Add(time.Duration(secs) * time.Second)
}

// TNowAddDateTimeDHM 当前时间增加指定的天数、小时数、分钟数，得到未来的时间点
func TNowAddDateTimeDHM(days int, hours int, minutes int) time.Time {
	return TNowAddDateTime(0, 0, days, hours, minutes, 0)
}

// TNowAddInterval 当前时间增加指定的年、月、日、天、小时、分钟，秒 后，得到未来的时间点和当前时间点之间的差值
func TNowAddInterval(years, months, days, hours, minutes, secs int) time.Duration {
	now := time.Now()
	future := now.AddDate(years, months, days).
		Add(time.Duration(hours) * time.Hour).
		Add(time.Duration(minutes) * time.Minute).
		Add(time.Duration(secs) * time.Second)
	duration := future.Sub(now)
	return duration
}

// TNowAddIntervalDHM 当前时间增加指定的天数、小时数、分钟数，得到未来的时间点和当前时间点之间的差值
func TNowAddIntervalDHM(days int, hours int, minutes int) time.Duration {
	return TNowAddInterval(0, 0, days, hours, minutes, 0)
}

// TNowAddIntervalSec 当前时间增加指定的年、月、日、天、小时、分钟，秒 后，得到未来的时间点和当前时间点之间的差值秒数
func TNowAddIntervalSec(years, months, days, hours, minutes, secs int) int64 {
	duration := TNowAddInterval(years, months, days, hours, minutes, secs)
	return int64(math.Abs(duration.Seconds()))
}

// TNowAddIntervalSecDHM 当前时间增加指定的天数、小时数、分钟数，得到未来的时间点和当前时间点之间的差值秒数
func TNowAddIntervalSecDHM(days int, hours int, minutes int) int64 {
	return TNowAddIntervalSec(0, 0, days, hours, minutes, 0)
}

// ----------------------------

func TTimeS(t time.Time) int64 {
	return t.Unix()
}

func TTimeMs(t time.Time) int64 {
	return t.UnixMilli()
}

func TTimeFormat(t time.Time, layout TFormat) string {
	return t.Format(string(layout))
}

func TTimeFormatStr(t time.Time, layout string) string {
	return t.Format(layout)
}

// ----------------------------

// TSecByYear one year for 365 days
func TSecByYear(years int) int64 {
	duration := time.Duration(years) * 365 * 24 * time.Hour
	return int64(duration.Seconds())
}

// TSecByMonth one month for 30 days
func TSecByMonth(months int) int64 {
	duration := time.Duration(months) * 30 * 24 * time.Hour
	return int64(duration.Seconds())
}

func TSecByDay(days int) int64 {
	duration := time.Duration(days) * 24 * time.Hour
	return int64(duration.Seconds())
}

func TSecByHour(hours int) int64 {
	duration := time.Duration(hours) * time.Hour
	return int64(duration.Seconds())
}

func TSecByMin(minutes int) int64 {
	duration := time.Duration(minutes) * time.Minute
	return int64(duration.Seconds())
}

// TSecFormatMS 将秒转换为 `时:分:秒` 形式
// 当总秒数小于1小时时，为简写形式：00:40；当总秒数超过1小时时，为完整形式：47:55:49
func TSecFormatMS(secs int64) string {
	duration := time.Duration(secs) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := secs % 60

	if secs < 3600 {
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	} else {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}
}

// TSecFormatHMS 将秒转换为 `时:分:秒` 形式
// 完整形式：00:55:49
func TSecFormatHMS(secs int64) string {
	duration := time.Duration(secs) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := secs % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// ----------------------------

// --------- Deprecated -------------------

// Deprecated: Use rose.TNowAddIntervalSecDHM instead.
func TSecUntilNowToAdd(days int, hours int, minutes int) int64 {
	return TSecUntilNowToAddY(0, 0, days, hours, minutes, 0)
}

// Deprecated: Use rose.TNowAddIntervalSec instead.
func TSecUntilNowToAddY(years, months, days, hours, minutes, sec int) int64 {
	now := time.Now()
	future := now.AddDate(years, months, days).
		Add(time.Duration(hours) * time.Hour).
		Add(time.Duration(minutes) * time.Minute).
		Add(time.Duration(sec) * time.Second)
	duration := future.Sub(now).Seconds()
	return int64(duration)
}

// Deprecated: Use rose.TSecFormatHMS instead.
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
