package rose

import (
	"math"
	"strconv"
)

func IntToInt64(i int) int64 {
	return int64(i)
}

func Int64ToInt(i int64) int {
	return int(i)
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func IntIsBetween(i, a, b int) bool {
	return a < i && i < b
}

func IntIsBetweenL(i, a, b int) bool {
	return a <= i && i < b
}

func IntIsBetweenR(i, a, b int) bool {
	return a < i && i <= b
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int64IsBetween(i, a, b int64) bool {
	return a < i && i < b
}

func Int64IsBetweenL(i, a, b int64) bool {
	return a <= i && i < b
}

func Int64IsBetweenR(i, a, b int64) bool {
	return a < i && i <= b
}

// Float64ToStr 将 float64 转换为str
func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Float32ToStr(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

func UInt32ToStr(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

// UInt64ToStr
func UInt64ToStr(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// Float64Round 获取float64保留n位小数
func Float64Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc(f*pow10N+0.5) / pow10N //0.5是为了四舍五入
}

// Float64RoundInt64 将小数四舍五入得到整数
func Float64RoundInt64(v float64) int64 {
	return int64(math.Floor(v + 0.5))
}
