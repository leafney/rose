package rose

import (
	"strconv"
	"strings"
	"unicode"
)

// StrIsEmpty Check if a string is empty
func StrIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// StrToInt 将数字字符串转换成数字类型
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

// StrToIntWithDef 转换失败返回预设值
func StrToIntWithDef(s string, def int) int {
	if s == "" {
		return def
	}
	if i, err := strconv.Atoi(s); err != nil {
		return def
	} else {
		return i
	}
}

// StrToIntWithErr 将数字字符串转换成数字类型
func StrToIntWithErr(s string) (int, error) {
	return strconv.Atoi(s)
}

// StrToInt64 将字符串转换为int64(转换失败返回0
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

// StrToInt64WithDef 转换失败返回预设值
func StrToInt64WithDef(s string, def int64) int64 {
	if s == "" {
		return def
	}
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		return def
	} else {
		return i
	}
}

// StrToInt64WithErr 将字符串转换为int64
func StrToInt64WithErr(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// StrToFloat64 将数字字符串转换成float64类型(转换失败返回0.0
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

// StrToFloat64WithDef 转换失败返回预设值
func StrToFloat64WithDef(s string, def float64) float64 {
	if s == "" {
		return def
	}
	if i, err := strconv.ParseFloat(s, 64); err != nil {
		return def
	} else {
		return i
	}
}

// 将数字字符串转换成float64类型
func StrToFloat64WithErr(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func StrToBool(s string) bool {
	if b, err := strconv.ParseBool(s); err != nil {
		return false
	} else {
		return b
	}
}

// StrToBoolWithDef 转换失败返回预设值
func StrToBoolWithDef(s string, def bool) bool {
	if b, err := strconv.ParseBool(s); err != nil {
		return def
	} else {
		return b
	}
}

func StrToBoolWithErr(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// StrToChar convert string to char slice
func StrToChar(s string) []string {
	c := make([]string, 0)
	if len(s) == 0 {
		c = append(c, "")
	}
	for _, v := range s {
		c = append(c, string(v))
	}
	return c
}

// StrToUnderScoreName 将字符串驼峰式写法转为下划线写法
func StrToUnderScoreName(name string) string {
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

// StrToCamelName 将字符串下划线写法转为驼峰写法
func StrToCamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// StrSplitAny 对字符串使用任意一个或多个字符分隔，同时排除空字符
func StrSplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

// StrRemoveAny 移除字符串中一个或多个字符
func StrRemoveAny(s string, seps ...string) string {
	for _, sep := range seps {
		if strings.Contains(s, sep) {
			s = strings.ReplaceAll(s, sep, "")
		}
	}
	return s
}

// StrContainsAny 判断是否包含其中的某个字符串
func StrContainsAny(s string, seps ...string) bool {
	for _, sep := range seps {
		if strings.Contains(s, sep) {
			return true
		}
	}
	return false
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

// StrJoin 字符串拼接
func StrJoin(args ...string) string {
	var buffer strings.Builder
	for _, arg := range args {
		buffer.WriteString(arg)
	}
	return buffer.String()
}

// StrEqualFold 比较两个字符串是否相同，不区分大小写
func StrEqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

// StrEqualFull 比较两个字符串是否完全相等，区分大小写
func StrEqualFull(s, t string) bool {
	return s == t
}

func StrPrefixAny(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func StrSuffixAny(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func StrPrefixAnyI(s string, prefixes ...string) bool {
	st := strings.ToLower(s)
	for _, prefix := range prefixes {
		if strings.HasPrefix(st, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}

// StrSuffixAnyI is case insensitive HasSuffix
func StrSuffixAnyI(s string, suffixes ...string) bool {
	st := strings.ToLower(s)
	for _, suffix := range suffixes {
		if strings.HasSuffix(st, strings.ToLower(suffix)) {
			return true
		}
	}
	return false
}
