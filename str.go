package rose

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

// StrIsEmpty Check if a string is empty
func StrIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// StrTrim TrimSpace
func StrTrim(s string) string {
	return strings.TrimSpace(s)
}

// StrTrimL TrimPrefix
func StrTrimL(s, left string) string {
	return strings.TrimPrefix(s, left)
}

// StrTrimR TrimSuffix
func StrTrimR(s, right string) string {
	return strings.TrimSuffix(s, right)
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

// StrToIntDef 将数字字符串转换为数值类型，转换失败使用默认值
func StrToIntDef(s string, def int) int {
	if s == "" {
		return def
	}
	if i, err := strconv.Atoi(s); err != nil {
		return def
	} else {
		return i
	}
}

// StrToIntErr 将数字字符串转换为数值类型，转换失败抛出异常
func StrToIntErr(s string) (int, error) {
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

// StrToInt64Def 转换失败返回预设值
func StrToInt64Def(s string, def int64) int64 {
	if s == "" {
		return def
	}
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		return def
	} else {
		return i
	}
}

// StrToInt64Err 将字符串转换为int64
func StrToInt64Err(s string) (int64, error) {
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

// StrToFloat64Def 将数字字符串转换成float64类型
func StrToFloat64Def(s string, def float64) float64 {
	if s == "" {
		return def
	}
	if i, err := strconv.ParseFloat(s, 64); err != nil {
		return def
	} else {
		return i
	}
}

// StrToFloat64Err 将数字字符串转换成float64类型
func StrToFloat64Err(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func StrToBool(s string) bool {
	if b, err := strconv.ParseBool(s); err != nil {
		return false
	} else {
		return b
	}
}

// StrToBoolDef 转换失败返回预设值
func StrToBoolDef(s string, def bool) bool {
	if b, err := strconv.ParseBool(s); err != nil {
		return def
	} else {
		return b
	}
}

func StrToBoolErr(s string) (bool, error) {
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

// -----------------

// StrAnySplit 对字符串使用任意一个或多个字符分隔，**同时排除空字符**
func StrAnySplit(s string, seps ...string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}

	sep := SliceStrToStr(seps, "")
	splitter := func(r rune) bool {
		return strings.ContainsRune(sep, r)
	}
	return strings.FieldsFunc(s, splitter)
}

// StrAnySplitFirst 获取分隔后数组第一项
func StrAnySplitFirst(s string, seps ...string) string {
	strs := StrAnySplit(s, seps...)
	if len(strs) > 0 {
		return strs[0]
	}
	return ""
}

// StrAnySplitLast 获取分隔后数组最后一项
func StrAnySplitLast(s string, seps ...string) string {
	strs := StrAnySplit(s, seps...)
	if len(strs) > 0 {
		return strs[len(strs)-1]
	}
	return ""
}

// StrAnySplitIndex 获取分隔后数组的指定索引项
func StrAnySplitIndex(s string, index int, seps ...string) string {
	strs := StrAnySplit(s, seps...)
	for i, val := range strs {
		if i == index {
			return val
		}
	}
	return ""
}

// StrAnyTrim 移除字符串首部以及尾部的任意指定字符
func StrAnyTrim(s string, seps ...string) string {
	// 先对分隔字符串按照长度由大到小排序
	seps = SliceSortByLength(seps, false)

	for _, sep := range seps {
		s = strings.Trim(s, sep)
	}
	return s
}

// StrAnyRemove 移除字符串中包含的任意指定字符
func StrAnyRemove(s string, seps ...string) string {
	// 先对分隔字符串按照长度由大到小排序
	seps = SliceSortByLength(seps, false)

	for _, sep := range seps {
		if strings.Contains(s, sep) {
			s = strings.ReplaceAll(s, sep, "")
		}
	}
	return s
}

// StrAnyReplace 将字符串中包含的任意指定字符串替换为新的字符串
func StrAnyReplace(s string, new string, seps ...string) string {
	// 先对分隔字符串按照长度由大到小排序
	seps = SliceSortByLength(seps, false)
	for _, sep := range seps {
		if strings.Contains(s, sep) {
			s = strings.ReplaceAll(s, sep, new)
		}
	}
	return s
}

// StrAnyContains 判断字符串中是否包含指定的任意字符串
func StrAnyContains(s string, seps ...string) bool {
	for _, sep := range seps {
		if strings.Contains(s, sep) {
			return true
		}
	}
	return false
}

// StrAnyPrefix 是否以任何前缀字符串开头，区分大小写
func StrAnyPrefix(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// StrAnySuffix 是否以任何后缀字符串结尾，区分大小写
func StrAnySuffix(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// StrAnyPrefixI 是否以任何前缀字符串开头，不区分大小写
func StrAnyPrefixI(s string, prefixes ...string) bool {
	st := strings.ToLower(s)
	for _, prefix := range prefixes {
		if strings.HasPrefix(st, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}

// StrAnySuffixI 是否以任何后缀字符串结尾，不区分大小写
func StrAnySuffixI(s string, suffixes ...string) bool {
	st := strings.ToLower(s)
	for _, suffix := range suffixes {
		if strings.HasSuffix(st, strings.ToLower(suffix)) {
			return true
		}
	}
	return false
}

// StrSpecialTrim 清除特定字符（如换行符、回车符、制表符、空格等），无论在前、后、还是中间
func StrSpecialTrim(s string) string {
	return strings.Map(func(r rune) rune {
		if r == '\n' || r == '\r' || r == '\t' || r == ' ' {
			return -1 // 删除这些字符
		}
		return r
	}, s)
}

// -------------------

// StrJoin 字符串拼接
func StrJoin(args ...string) string {
	var buffer strings.Builder
	for _, arg := range args {
		buffer.WriteString(arg)
	}
	return buffer.String()
}

// StrJoinWith alias strings.Join()
func StrJoinWith(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// StrEqualFold 比较两个字符串内容是否相同，忽略大小写差异
func StrEqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

// StrEqualFull 比较两个字符串是否完全相同，大小写敏感
func StrEqualFull(s, t string) bool {
	return s == t
}

// StrToLower 将字符串转换为小写形式
func StrToLower(s string) string {
	return strings.ToLower(s)
}

// StrToUpper 将字符串转换为大写形式
func StrToUpper(s string) string {
	return strings.ToUpper(s)
}

// StrSizeSplit 按照指定长度大小拆分字符串
func StrSizeSplit(s string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

// -----------------

// TODO 待优化

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

// Substring 返回截取后的字符串
// `start >=0` 从左往右，索引从0开始 ；`start <0` 从右往左，且倒数第一个字符索引为-1
func Substring(str string, start, length int) string {
	// 参数检查
	if length <= 0 {
		return ""
	}

	// 将字符串转换为 rune 切片
	runes := []rune(str)
	strLen := len(runes)

	// 如果 start 为负值，则从右往左开始算
	if start < 0 {
		start = strLen + start
	}

	// 计算起始索引
	//start = max(0, min(start, strLen)) // go v1.21
	start = IntMax(0, IntMin(start, strLen))
	// 计算结束索引
	//end := min(start+length, strLen) // go v1.21
	end := IntMin(start+length, strLen)

	return string(runes[start:end])
}

// StrToUniqueString 根据输入字符串，生成一个唯一的字符串，可以指定生成字符串长度
func StrToUniqueString(input string, minLength, maxLength int) string {
	timestamp := TNowS()
	hashedString := Sha1(fmt.Sprintf("%s%d", input, timestamp))
	hashedLength := len(hashedString)
	strLength := minLength + rand.Intn(maxLength-minLength+1)
	if strLength > hashedLength {
		strLength = hashedLength
	}
	return hashedString[:strLength]
}

// --------- Deprecated -------------------

// StrToIntWithDef 转换失败返回预设值
// Deprecated: Use rose.StrToIntDef instead.
func StrToIntWithDef(s string, def int) int {
	return StrToIntDef(s, def)
}

// Deprecated: Use rose.StrTrimL instead.
func StrLTrim(s, left string) string {
	return StrTrimL(s, left)
}

// Deprecated: Use rose.StrTrimR instead.
func StrRTrim(s, right string) string {
	return StrTrimR(s, right)
}

// StrToIntWithErr 将数字字符串转换成数字类型
// Deprecated: Use rose.StrToIntErr instead.
func StrToIntWithErr(s string) (int, error) {
	return StrToIntErr(s)
}

// Deprecated: Use rose.StrToInt64Def instead.
func StrToInt64WithDef(s string, def int64) int64 {
	return StrToInt64Def(s, def)
}

// Deprecated: Use rose.StrToInt64Err instead.
func StrToInt64WithErr(s string) (int64, error) {
	return StrToInt64Err(s)
}

// Deprecated: Use rose.StrToFloat64Def instead.
func StrToFloat64WithDef(s string, def float64) float64 {
	return StrToFloat64Def(s, def)
}

// Deprecated: Use rose.StrToFloat64Err instead.
func StrToFloat64WithErr(s string) (float64, error) {
	return StrToFloat64Err(s)
}

// Deprecated: Use rose.StrToBoolDef instead.
func StrToBoolWithDef(s string, def bool) bool {
	return StrToBoolDef(s, def)
}

// Deprecated: Use rose.StrToBoolErr instead.
func StrToBoolWithErr(s string) (bool, error) {
	return StrToBoolErr(s)
}

// Deprecated: Use rose.StrAnySplit instead.
func StrSplitAny(s string, seps ...string) []string {
	return StrAnySplit(s, seps...)
}

// Deprecated: Use rose.StrAnyTrim instead.
func StrTrimAny(s string, seps ...string) string {
	return StrAnyTrim(s, seps...)
}

// Deprecated: Use rose.StrAnyRemove instead.
func StrRemoveAny(s string, seps ...string) string {
	return StrAnyRemove(s, seps...)
}

// Deprecated: Use rose.StrAnyContains instead.
func StrContainsAny(s string, seps ...string) bool {
	return StrAnyContains(s, seps...)
}

// Deprecated: Use rose.StrAnyPrefix instead.
func StrPrefixAny(s string, prefixes ...string) bool {
	return StrAnyPrefix(s, prefixes...)
}

// Deprecated: Use rose.StrAnySuffix instead.
func StrSuffixAny(s string, suffixes ...string) bool {
	return StrAnySuffix(s, suffixes...)
}

// Deprecated: Use rose.StrAnyPrefixI instead.
func StrPrefixAnyI(s string, prefixes ...string) bool {
	return StrAnyPrefixI(s, prefixes...)
}

// Deprecated: Use rose.StrAnySuffixI instead.
func StrSuffixAnyI(s string, suffixes ...string) bool {
	return StrAnySuffixI(s, suffixes...)
}
