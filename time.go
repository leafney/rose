package rose

import (
	"errors"
	"fmt"
	"math"
	"time"
)

const (
	TimeLayoutShortY     = "2006"
	TimeLayoutShortYM    = "200601"
	TimeLayoutShortYMD   = "20060102"
	TimeLayoutShortYMDH  = "200601021504"
	TimeLayoutShortYMDHS = "20060102150405"
	TimeLayoutLongYM     = "2006-01"
	TimeLayoutLongYMD    = "2006-01-02"
	TimeLayoutLongYMDH   = "2006-01-02 15:04"
	TimeLayoutLongYMDHS  = "2006-01-02 15:04:05"
	TimeLayoutTimeH      = "15:04"
	TimeLayoutTimeHS     = "15:04:05"
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

const (
	TLocChina = "Asia/Shanghai"
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
	return time.Now().Format(TimeLayoutShortYMD)
}

// TNowDateLYMD 当前日期字符串，格式：2023-03-23
func TNowDateLYMD() string {
	return time.Now().Format(TimeLayoutLongYMD)
}

// TNowDateSYMDH 当前日期字符串，格式：202303232307
func TNowDateSYMDH() string {
	return time.Now().Format(TimeLayoutShortYMDH)
}

// TNowDateLYMDH 当前日期字符串，格式：2023-03-23 23:07
func TNowDateLYMDH() string {
	return time.Now().Format(TimeLayoutLongYMDH)
}

// TNowDateTime 当前日期字符串，格式：2023-03-23 23:06:42
func TNowDateTime() string {
	return time.Now().Format(TimeLayoutLongYMDHS)
}

// TNowDateTimeS 当前日期字符串，格式：20230323230642
func TNowDateTimeS() string {
	return time.Now().Format(TimeLayoutShortYMDHS)
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

// TParseDateTimeToTime 将时间字符串 `2006-01-02 15:04:05` 转换为 time.Time
func TParseDateTimeToTime(ds string) time.Time {
	return TParseFormatSTTInLocByTF(TFLongYMDHS, ds, time.Local)
}

// TParseDateYMDToTime 将时间字符串`2006-01-02`转换为time.Time
func TParseDateYMDToTime(ds string) time.Time {
	return TParseFormatSTTByTF(TFLongYMD, ds)
}

// TParseDateYMDHToTime 将时间字符串`2006-01-02 15:04`转换为time.Time
func TParseDateYMDHToTime(ds string) time.Time {
	return TParseFormatSTTByTF(TFLongYMDH, ds)
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

// ----------------------------

// TParseLocation 转换时区字符串 eg: "Asia/Shanghai"
func TParseLocation(location string) *time.Location {
	loc, _ := time.LoadLocation(location)
	return loc
}

// TParseFormatSTT 将自定义格式时间字符串转换为 time.Time
func TParseFormatSTT(layout string, ds string) time.Time {
	return TParseFormatSTTInLoc(layout, ds, time.Local)
}

// TParseFormatSTTE 将自定义格式时间字符串转换为time.Time，支持异常捕获
func TParseFormatSTTE(layout string, ds string) (time.Time, error) {
	return TParseFormatSTTInLocE(layout, ds, time.Local)
}

// TParseFormatSTTByTF 将内置格式时间字符串转换为 time.Time
func TParseFormatSTTByTF(layout TFormat, ds string) time.Time {
	return TParseFormatSTT(string(layout), ds)
}

// TParseFormatSTTByTFE 将内置格式时间字符串转换为 time.Time，支持异常捕获
func TParseFormatSTTByTFE(layout TFormat, ds string) (time.Time, error) {
	return TParseFormatSTTInLocE(string(layout), ds, time.Local)
}

// TParseFormatSTTInLoc 将自定义格式时间字符串转换为time.Time，支持指定时区
func TParseFormatSTTInLoc(layout string, ds string, loc *time.Location) time.Time {
	parse, _ := TParseFormatSTTInLocE(layout, ds, loc)
	return parse
}

// TParseFormatSTTInLocByTF 将内置格式时间字符串转换为time.Time，支持指定时区
func TParseFormatSTTInLocByTF(layout TFormat, ds string, loc *time.Location) time.Time {
	return TParseFormatSTTInLoc(string(layout), ds, loc)
}

// TParseFormatSTTInLocE StringToTime 将自定义格式时间字符串转换为time.Time，支持指定时区，支持异常捕获
func TParseFormatSTTInLocE(layout string, ds string, loc *time.Location) (time.Time, error) {
	if StrIsEmpty(ds) {
		return time.Time{}, errors.New("the time string is empty")
	}
	return time.ParseInLocation(layout, ds, loc)
}

func TParseFormatSTTInLocStr(layout string, ds string, location string) time.Time {
	parse, _ := TParseFormatSTTInLocStrE(layout, ds, location)
	return parse
}

// TParseFormatSTTInLocStrE 将 自定义格式时间字符串 转换为 time.Time，支持指定时区，支持异常捕获
func TParseFormatSTTInLocStrE(layout string, ds string, location string) (time.Time, error) {
	if StrIsEmpty(ds) {
		return time.Time{}, errors.New("the time string is empty")
	}
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return TParseFormatSTTInLocE(layout, ds, loc)
}

// TParseGMTSTTE 将 GMT time.RFC1123 时间格式字符串转换为 time.Time
//
// eg: Mon, 20 Jul 2020 06:09:21 GMT => Time
// https://golang.org/pkg/time/#pkg-constants
func TParseGMTSTTE(gmt string) (time.Time, error) {
	return time.Parse(time.RFC1123, gmt)
}

// TParseGMTSTTInLocE 将 GMT time.RFC1123 时间格式字符串转换为 time.Time ，指定时区
func TParseGMTSTTInLocE(gmt string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(time.RFC1123, gmt, loc)
}

func TParseGMTSTTInLocStrE(gmt string, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return TParseGMTSTTInLocE(gmt, loc)
}

// TParseRFC9990700STTE 将 时间字符串 2006-01-02T15:04:05.999-0700 格式化为 time.Time ，本地时区
//
// eg：2023-11-22T12:49:38.758+0000 => Time
func TParseRFC9990700STTE(ds string) (time.Time, error) {
	return TParseRFC9990700STTInLocE(ds, time.Local)
}

// TParseRFC9990700STTInLocE 将 时间字符串 2006-01-02T15:04:05.999-0700 格式化为 time.Time ，指定时区
func TParseRFC9990700STTInLocE(ds string, loc *time.Location) (time.Time, error) {
	layout := "2006-01-02T15:04:05.999-0700"
	return time.ParseInLocation(layout, ds, loc)
}

func TParseRFC9990700STTInLocStrE(ds string, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return TParseRFC9990700STTInLocE(ds, loc)
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

// TTotalSecondsInt 将天数、小时数、分钟数转化为秒
func TTotalSecondsInt(days, hours, minutes, seconds int) int {
	totalSeconds := (days * 24 * 60 * 60) + (hours * 60 * 60) + (minutes * 60) + seconds
	return totalSeconds
}

// TTotalSecondsInt64 将天数、小时数、分钟数转化为秒
func TTotalSecondsInt64(days, hours, minutes, seconds int64) int64 {
	totalSeconds := (days * 24 * 60 * 60) + (hours * 60 * 60) + (minutes * 60) + seconds
	return totalSeconds
}

// ----------------------------

// TMonthS 获取当前月份 eg: 202303
func TMonthS() string {
	return time.Now().Format(TimeLayoutShortYM)
}

// TMonthL 获取当前月份 eg: 2023-03
func TMonthL() string {
	return time.Now().Format(TimeLayoutLongYM)
}

// TMonthNextS 获取下月月份 eg：202304
func TMonthNextS() string {
	return time.Now().AddDate(0, 1, 0).Format(TimeLayoutShortYM)
}

// TMonthNextL 获取下月月份 eg：2023-04
func TMonthNextL() string {
	return time.Now().AddDate(0, 1, 0).Format(TimeLayoutLongYM)
}

// TMonthLastS 获取上月月份 eg: 202302
func TMonthLastS() string {
	return time.Now().AddDate(0, -1, 0).Format(TimeLayoutShortYM)
}

// TMonthLastL 获取上月月份 eg: 2023-02
func TMonthLastL() string {
	return time.Now().AddDate(0, -1, 0).Format(TimeLayoutLongYM)
}

// ----------------------------

// TMonthStartDay 当前月份的第一天的日期
func TMonthStartDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// TMonthStartTime 当前月份的第一天的起始时间
func TMonthStartTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// TMonthStartDayStr 当前月份的第一天日期字符串
func TMonthStartDayStr() string {
	return TTimeFormat(TMonthStartDay(), TFLongYMD)
}

// TMonthEndDay 当前月份的最后一天的日期
func TMonthEndDay() time.Time {
	now := time.Now()
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	return nextMonth.AddDate(0, 0, -1)
}

// TMonthEndTime 当前月份的最后一天的截止时间
func TMonthEndTime() time.Time {
	now := time.Now()
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 23, 59, 59, 0, now.Location())
	return nextMonth.AddDate(0, 0, -1)
}

// TMonthEndDayStr 当前月份的最后一天日期字符串
func TMonthEndDayStr() string {
	return TTimeFormat(TMonthEndDay(), TFLongYMD)
}

// TMonthRemainDays 获取当前日期到本月末的剩余天数（是否包含今天
func TMonthRemainDays(includeToday bool) int {
	now := time.Now()
	lastDayOfMonth := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, now.Location())
	remainDays := lastDayOfMonth.Day() - now.Day()
	if includeToday {
		remainDays += 1
	}

	return remainDays
}

// ----------------------------

// TDate 获取当前日期 eg: 2019-01-09
func TDate() string {
	return time.Now().Format(TimeLayoutLongYMD)
}

// TDateNext 获取下月日期 eg: 2019-02-09
func TDateNext() string {
	return time.Now().AddDate(0, 1, 0).Format(TimeLayoutLongYMD)
}

// TDateLast 获取上月日期 eg: 2018-12-09
func TDateLast() string {
	return time.Now().AddDate(0, -1, 0).Format(TimeLayoutLongYMD)
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

// TTimeFormat Time.Format()
func TTimeFormat(t time.Time, layout TFormat) string {
	return t.Format(string(layout))
}

func TTimeFormatStr(t time.Time, layout string) string {
	return t.Format(layout)
}

func TTimeFormatInLoc(t time.Time, layout string, loc *time.Location) string {
	return t.In(loc).Format(layout)
}

// TTimeFormatInLocStr 带有时区的时间格式化
func TTimeFormatInLocStr(t time.Time, layout, location string) string {
	loc := TParseLocation(location)
	return t.In(loc).Format(layout)
}

func TTimeFormatInLocByTF(t time.Time, layout TFormat, loc *time.Location) string {
	return TTimeFormatInLoc(t, string(layout), loc)
}

// TTimeFormatInLocStrByTF 带有时区的时间格式化，使用内置格式
//
// 可以使用时区格式如：`Asia/Shanghai` 或者 `time.UTC` `time.Local` 等
func TTimeFormatInLocStrByTF(t time.Time, layout TFormat, location string) string {
	return TTimeFormatInLocStr(t, string(layout), location)
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

// TSecFormatMS 将总秒数转换为 `分:秒` 或 `时:分:秒` 字符串形式
// 当总秒数小于1小时内，为简写形式：00:40；当总秒数超过1小时，为完整形式：47:55:49
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

// TSecFormatHMS 将总秒数转换为 `时:分:秒` 字符串形式
// 完整形式：00:55:49，如果总秒数小于 1 小时内，小时部分为 `00`
func TSecFormatHMS(secs int64) string {
	duration := time.Duration(secs) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := secs % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// TSecToDHMS 获取总秒数对应的天数、小时数、分钟数及秒数
func TSecToDHMS(secs int64) (days, hours, minutes, seconds int64) {
	days = secs / (60 * 60 * 24)
	hours = (secs % (60 * 60 * 24)) / (60 * 60)
	minutes = (secs % (60 * 60)) / 60
	seconds = secs % 60
	return
}

// ----------------------------

// TSleepS sleep seconds
func TSleepS(secs int64) {
	time.Sleep(time.Duration(secs) * time.Second)
}

func TSleepSRand(min, max int64) {
	time.Sleep(time.Duration(RandInt64Range(min, max)) * time.Second)
}

// TSleepMs sleep milliseconds
func TSleepMs(msecs int64) {
	time.Sleep(time.Duration(msecs) * time.Millisecond)
}

func TSleepMsRand(min, max int64) {
	time.Sleep(time.Duration(RandInt64Range(min, max)) * time.Millisecond)
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
