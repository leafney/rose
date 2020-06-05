package rose

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

// 判断字符串是否为空
func StrIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// 将数字字符串转换成数字类型
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

func StrToIntWithErr(s string) (int, error) {
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

// 将字符串转换为int64
func StrToInt64WithErr(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func StrSliceRemoveEmpty(sls []string) []string {
	newSls := make([]string, 0)
	for k, v := range sls {
		if len(v) > 0 {
			newSls = append(newSls, sls[k])
		}
	}
	return newSls
}

// 将一个array分成多个指定数量的分组
func StrSliceArray(arr []string, size int) (result [][]string) {
	ln := len(arr)
	cs := int(math.Ceil(float64(ln) / float64(size)))

	for x := 0; x < cs; x++ {
		start := x * size
		end := start + size
		if end > ln {
			end = ln
		}
		result = append(result, arr[start:end])
	}
	return
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

// *************

// 驼峰式写法转为下划线写法
func UnderscoreName(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}

	return buffer.String()
}

// 下划线写法转为驼峰写法
func CamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
