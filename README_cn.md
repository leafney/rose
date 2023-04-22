# Rose

**Golang 工具类**

[英文](README.md)

----

## 安装

```sh
go get -u github.com/leafney/rose
```

----

## 方法列表

### snowflake算法

* `InitWorker` -- 初始化worker
* `GetNextId` -- 获取唯一id

---

### Url

* `UrlJoin` -- 合并url请求链接
* `UrlJoinWithQuery` -- 合并带有请求参数的链接

---

### Str

- `StrIsEmpty(s string) bool` -- 判断字符串是否为空
- `StrTrim(s string) string` --
- `StrTrimL(s, left string) string` --
- `StrTrimR(s, right string) string` --
- `StrToInt(s string) int` --
- `StrToIntDef(s string, def int) int` --
- `StrToIntErr(s string) (int, error)` --
- `StrToInt64(s string) int64` --
- `StrToInt64Def(s string, def int64) int64` --
- `StrToInt64Err(s string) (int64, error)` --
- `StrToFloat64(s string) float64` --
- `StrToFloat64Def(s string, def float64) float64` --
- `StrToFloat64Err(s string) (float64, error)` --
- `StrToBool(s string) bool` --
- `StrToBoolDef(s string, def bool) bool` --
- `StrToBoolErr(s string) (bool, error)` --
- `StrToChar(s string) []string` --
- `StrAnySplit(s string, seps ...string) []string` -- 对字符串使用任意一个或多个字符分隔
- `StrAnyTrim(s string, seps ...string) string` -- 移除字符串首部以及尾部的任意指定字符
- `StrAnyRemove(s string, seps ...string) string` -- 移除字符串中包含的任意指定字符
- `StrAnyReplace(s string, new string, seps ...string) string` -- 将字符串中包含的任意指定字符串替换为新的字符串
- `StrAnyContains(s string, seps ...string) bool` -- 判断字符串中是否包含指定的任意字符串
- `StrAnyPrefix(s string, prefixes ...string) bool` -- 是否以任何前缀字符串开头，区分大小写
- `StrAnySuffix(s string, suffixes ...string) bool {` -- 是否以任何后缀字符串结尾，区分大小写
- `StrAnyPrefixI(s string, prefixes ...string) bool` -- 是否以任何前缀字符串开头，不区分大小写
- `StrAnySuffixI(s string, suffixes ...string) bool` -- 是否以任何后缀字符串结尾，不区分大小写
- `StrJoin(args ...string) string` -- 字符串拼接
- `StrEqualFold(s, t string) bool` -- 比较两个字符串是否相同，不区分大小写
- `StrEqualFull(s, t string) bool` -- 比较两个字符串是否完全相等，区分大小写
- `StrToLower(s string) string` -- 将字符串转换为小写形式
- `StrToUpper(s string) string` -- 将字符串转换为大写形式
- `StrToUnderscoreName` -- 将字符串的驼峰写法转成下划线写法
- `StrToCamelName` -- 将字符串下划线写法转成驼峰写法


----

### Num

- `IntToInt64`
- `Int64ToInt`
- `IntToStr`
- `Int64ToStr`
- `Float64ToStr`
- `Float32ToStr`
- `UInt32ToStr`
- `UInt64ToStr`
- `Float64Round` -- 将float64保留N位小数
- `Float64RoundInt64` -- 将小数四舍五入得到整数

----

### Time

#### now 相关

- `TNow() time.Time` -- 当前时间
- `TNowS() int64` -- 当前时间戳（秒 10 位
- `TNowStr() string` -- 当前时间戳字符串（秒 10 位
- `TNowMs() int64` -- 当前时间戳（毫秒 13 位
- `TNowMStr() string` -- 当前时间戳字符串（毫秒 13 位
- `TNowFormat(layout TFormat) string` -- 当前日期字符串，内置格式
- `TNowFormatStr(layout string) string` -- 当前日期字符串，自定义格式
- `TNowDateSYMD() string` -- 当前日期字符串，格式：20230323
- `TNowDateLYMD() string` -- 当前日期字符串，格式：2023-03-23  
  `TNowDateSYMDH() string` -- 当前日期字符串，格式：202303232307  
  `TNowDateLYMDH() string` -- 当前日期字符串，格式：2023-03-23 23:07
- `TNowDateTimeS() string` -- 当前日期字符串，格式：20230323230642
- `TNowDateTime() string` -- 当前日期字符串，格式：2023-03-23 23:06:42


#### NowAdd 相关

- `TNowAdd(d time.Duration) time.Time` --
- `TNowAddDate(years int, months int, days int) time.Time` --
- `TNowAddYears(years int) time.Time` --
- `TNowAddMonths(months int) time.Time` --
- `TNowAddDays(days int) time.Time` --
- `TNowAddUnixS(years, months, days int, hours, minutes, secs int) int64` -- 当前时间增加指定的年、月、日、天、小时、分钟，秒 后，得到未来时间点的时间戳
- `TNowAddUnixSDHM(days int, hours int, minutes int) int64` -- 当前时间增加指定的天数、小时数、分钟数，得到未来时间点的时间戳
- `TNowAddDateTime(years, months, days int, hours, minutes, secs int) time.Time` -- 当前时间增加指定的年、月、日、天、小时、分钟，秒 后，得到未来的时间点
- `TNowAddDateTimeDHM(days int, hours int, minutes int) time.Time` -- 当前时间增加指定的天数、小时数、分钟数，得到未来的时间点
- `TNowAddInterval(years, months, days, hours, minutes, secs int) time.Duration` -- 当前时间增加指定的年、月、日、天、小时、分钟，秒 后，得到未来的时间点和当前时间点之间的差值
- `TNowAddIntervalDHM(days int, hours int, minutes int) time.Duration` -- 当前时间增加指定的天数、小时数、分钟数，得到未来的时间点和当前时间点之间的差值
- `TNowAddIntervalSec(years, months, days, hours, minutes, secs int) int64` -- 当前时间增加指定的年、月、日、天、小时、分钟，秒 后，得到未来的时间点和当前时间点之间的差值秒数
- `TNowAddIntervalSecDHM(days int, hours int, minutes int) int64` -- 当前时间增加指定的天数、小时数、分钟数，得到未来的时间点和当前时间点之间的差值秒数


#### unix 相关

- `TUnixSToTime(s int64) time.Time` -- 将秒时间戳转换为对应的时间 time.Time
- `TUnixMsToTime(ms int64) time.Time` -- 将毫秒时间戳转换为 time.Time
- `TUnixSToDateYMD(s int64) string` -- 将秒时间戳转换为 `2006-01-02` 格式字符串
- `TUnixSToDateYMDH(s int64) string` -- 将秒时间戳转换为 `2006-01-02 15:04` 格式字符串
- `TUnixSToDateTime(s int64) string` -- 将秒时间戳转换为 `2006-01-02 15:04:05` 格式字符串
- `TUnixSFormat(s int64, layout TFormat) string` -- 将时间戳转换为指定格式的时间字符串
- `TUnixSFormatStr(s int64, layout string) string` --
- `TUnixMsFormat(ms int64, layout TFormat) string` --
- `TUnixMsFormatStr(ms int64, layout string) string` --


#### parse 相关

- `TParseDateTimeToTime(ds string) time.Time` -- 将时间字符串 `2006-01-02 15:04:05` 转换为 time.Time
- `TParseDateYMDToTime` -- 将时间字符串 `2006-01-02` 转换为 time.Time
- `TParseDateYMDHToTime` -- 将时间字符串 `2006-01-02 15:04` 转换为 time.Time
- `TParseDateTimeToUnix` -- 将时间字符串 `2006-01-02 15:04:05` 转换为 Unix 时间戳
- `TParseDateYMDToUnix` --
- `TParseDateYMDHToUnix` --
- `TParseFormatToTime(layout TFormat, ds string) time.Time` --
- `TParseFormatToTimeLoc(layout TFormat, ds string, loc *time.Location) time.Time` --
- `TParseFormatStrToTime(layout string, ds string) time.Time` --
- `TParseFormatStrToTimeLoc(layout string, ds string, loc *time.Location) time.Time` --
- `TParseFormatToTimeE(layout TFormat, ds string) (time.Time, error)` --
- `TParseEFormatStrToTimeE(layout string, ds string) (time.Time, error)` --
- `TParseEFormatStrToTimeLocE(layout string, ds string, loc *time.Location) (time.Time, error)` --



#### time 相关

- `TTimeS(t time.Time) int64` --
- `TTimeMs(t time.Time) int64` --
- `TTimeFormat(t time.Time, format TFormat) string` --
- `TTimeFormatStr(t time.Time, layout string) string` --


#### day 相关

- `TDayStartTime() time.Time` -- 今天的起始时间
- `TDayStartUnixS() int64` -- 今天的起始时间戳秒
- `TDayEndTime() time.Time` -- 今天的结束时间
- `TDayEndUnixS() int64` -- 今天的结束时间戳秒
- `TDayEndRemainSec() int64` -- 截止到今日的 24 点之前剩余的秒数


#### month 相关

- `TMonthS() string` -- 获取当前月份 eg: 202303
- `TMonthL() string` -- 获取当前月份 eg: 2023-03
- `TMonthNextS() string` -- 获取下月月份 eg：202304
- `TMonthNextL() string` -- 获取下月月份 eg：2023-04
- `TMonthLastS() string` -- 获取上月月份 eg: 202302
- `TMonthLastL() string` -- 获取上月月份 eg: 2023-02
- `TMonthStartTime() time.Time` -- 当前月份的第一天日期
- `TMonthStartStr() string` -- 当前月份的第一天日期字符串
- `TMonthEndTime() time.Time` -- 当前月份的最后一天日期
- `TMonthEndStr() string` -- 当前月份的最后一天日期字符串


#### date 相关

- `TDate() string` -- 获取当前日期 eg: 2019-01-09
- `TDateNext() string` -- 获取下月日期 eg: 2019-02-09
- `TDateLast() string` -- 获取上月日期 eg: 2018-12-09

#### sec 相关

- `TSecByYear(years int) int64` --
- `TSecByMonth(months int) int64` --
- `TSecByDay(days int) int64` --
- `TSecByHour(hours int) int64` --
- `TSecByMin(minutes int) int64` --
- `TSecFormatMS(secs int64) string` -- 将秒转换为 `时:分:秒` 形式，当总秒数小于 1 小时时，为简写形式：00:40；当总秒数超过 1 小时时，为完整形式：47:55:49
- `TSecFormatHMS(secs int64) string` -- 将秒转换为 `时:分:秒` 形式，完整形式：00:55:49


-----

### reqx

http请求封装，链式调用，参考req。

```
 resp,err:= Get("http://jsonplaceholder.typicode.com/posts/1").
    SetDebug(true).
    SetTimeout(1 * time.Second).
    SetHeader("aaa", "bbb").
    Do()
```

-----

### Mask 掩码

- `MaskPhone(phone string) string` -- 手机号 前3后4
- `MaskEmail(email string) string` -- 邮箱 仅用户名部分 前1后1
- `MaskPwd(pwd string) string` -- 密码
- `MaskBankCard(card string) string` -- 银行卡号 前4后4
- `MaskIDCard(id string) string` -- 身份证号 前3后4
- `MaskStr(s string, front, after int, rep string, count int) string` -- 对字符串掩码处理，保留前front，后after，使用rep替换中间内容，并指定rep的数量