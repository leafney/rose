package rose

import (
	"math"
	"strconv"
)

//获取float64保留n位小数
func Float64Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc(f*pow10N+0.5) / pow10N //0.5是为了四舍五入
}
