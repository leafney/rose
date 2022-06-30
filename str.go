package rose

import (
	"strconv"
	"strings"
	"unicode"
)

// Check if a string is empty
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
func StrToInt64WithErr(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// 将数字字符串转换成float64类型(转换失败返回0.0
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

// 将数字字符串转换成float64类型
func StrToFloat64WithErr(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// 将字符串驼峰式写法转为下划线写法
func StrToUnderscoreName(name string) string {
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

// 将字符串下划线写法转为驼峰写法
func StrToCamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 对字符串使用任意字符分隔
func StrSplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

//start：正数 - 在字符串的指定位置开始,超出字符串长度强制把start变为字符串长度
//       负数 - 在从字符串结尾的指定位置开始
//       0 - 在字符串中的第一个字符处开始
//length:正数 - 从 start 参数所在的位置返回
//       负数 - 从字符串末端返回
func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])

}

/**
 * 字符串拼接
 */
func JoinString(args ...string) string {
	var buffer strings.Builder
	for _, arg := range args {
		buffer.WriteString(arg)
	}
	return buffer.String()
}
