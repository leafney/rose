/**
 * @Author:      leafney
 * @Date:        2021-10-09 18:14
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package conv

import "strconv"

// Convert a numeric string to a numeric type
func StrToInt(s string) int {
	if s == "" {
		return 0
	}
	if i, err := strconv.Atoi(s); err != nil {
		return 0
	} else {
		return i
	}
}

// Return the specified parameter if the conversion fails
func StrToIntDef(s string, v int) int {
	if s == "" {
		return v
	}
	if i, err := strconv.Atoi(s); err != nil {
		return v
	} else {
		return i
	}
}

func StrToIntErr(s string) (int, error) {
	return strconv.Atoi(s)
}

// 将字符串转换为int64(转换失败返回0
func StrToInt64(s string) int64 {
	if s == "" {
		return 0
	}
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		return 0
	} else {
		return i
	}
}

//
func StrToInt64Def(s string, v int64) int64 {
	if s == "" {
		return v
	}
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		return v
	} else {
		return i
	}
}

// 将字符串转换为int64
func StrToInt64Err(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// 将数字字符串转换成float64类型
func StrToFloat64(s string) float64 {
	if s == "" {
		return 0.0
	}
	if i, err := strconv.ParseFloat(s, 64); err != nil {
		return 0.0
	} else {
		return i
	}
}

//
func StrToFloat64Def(s string, v float64) float64 {
	if s == "" {
		return v
	}
	if i, err := strconv.ParseFloat(s, 64); err != nil {
		return 0.0
	} else {
		return i
	}
}

//
func StrToFloat64Err(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
