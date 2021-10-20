/**
 * @Author:      leafney
 * @Date:        2021-10-20 12:23
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package conv

import "strconv"

func IntToInt64(i int) int64 {
	return int64(i)
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Float32ToStr(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}
