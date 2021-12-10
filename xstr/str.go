package xstr

import (
	"math"
	"strings"
	"unicode"
)

// 判断字符串是否为空
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func SliceRmvEmpty(sls []string) []string {
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

// 字符串拼接
func JoinString(args ...string) string {
	buffer := NewBuffer()
	for _, arg := range args {
		buffer.append(arg)
	}
	return buffer.String()
}
