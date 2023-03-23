## Rose

**Golang Toolkit**

[英文](README.md)

----

### 安装

```sh
go get -u github.com/leafney/rose
```

----

### 方法列表

#### snowflake算法

* `InitWorker` -- 初始化worker
* `GetNextId` -- 获取唯一id

---

#### Url

* `UrlJoin` -- 合并url请求链接
* `UrlJoinWithQuery` -- 合并带有请求参数的链接

---

#### Str

- `StrIsEmpty` -- 判断字符串是否为空
- `StrToInt` -- 将字符串转换为int
- `StrToIntWithErr`
- `StrToInt64`
- `StrToInt64WithErr`
- `StrToFloat64`
- `StrToFloat64WithErr`
- `StrToBool`
- `StrToBoolWithErr`
- `StrToChar`
- `StrToUnderscoreName` -- 将字符串的驼峰写法转成下划线写法
- `StrToCamelName` -- 将字符串下划线写法转成驼峰写法
- `StrSplitAny` -- 对字符串使用任意字符分隔，支持单个或多个
- `Substr`
- `JoinString`

----

#### Num

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

#### Time

- `TNow` -- 当前时间
- `TNowS` -- 当前时间戳，10位
- `TNowStr` -- 当前时间戳字符串格式
- `TNowMs` -- 当前时间戳，13位
- `TNowMStr` -- 当前时间戳字符串格式
- `TUnixSToTime` -- 将10位时间戳转换成 `time.Time` 类型
- `TUnixMsToTime` -- 将13位时间戳转换成 `time.Time` 类型
- `ToDayRemainSec` -- 截止到今日的24点之前剩余的秒数
- `GetDate`
- `GetDateFromUnix` -- 将10位时间戳转成 `2022-07-17` 格式
- `GetTimeFromUnix` -- 将10位时间戳转成 `2022-07-17 15:23:10` 格式
- `GetTFromUnix`
- `GetTimeFromUnixFormat`
- `GetTimeMHFromUnix`
- `GetTimeParse`
- `GetDateParse`
- `StrDateMH2Time`
- `MonthStart`
- `TodayStart`
- `TodayEnd`
- `NowDate`
- `NowDateTime`
- `ParseDate`
- `ParseDateTime`
- `ParseStringTime`
- `ParseGMTTimeOfRFC1123`
- `FormatSec` -- 将秒转换成分秒、时分秒格式


-----

#### reqx

http请求封装，链式调用，参考req。

```
 resp,err:= Get("http://jsonplaceholder.typicode.com/posts/1").
    SetDebug(true).
    SetTimeout(1 * time.Second).
    SetHeader("aaa", "bbb").
    Do()
```

-----
