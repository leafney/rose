package rose

import (
	"math"
	"strconv"
)

func IntToInt64(i int) int64 {
	return int64(i)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

//获取float64保留n位小数
func Float64Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc(f*pow10N+0.5) / pow10N //0.5是为了四舍五入
}

/*
将 float64 转换为str
*/
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}
