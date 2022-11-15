package rose

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	// initialize seed once (https://stackoverflow.com/a/12321192/8155097)
	rand.Seed(time.Now().UnixNano())
}

// RandInt 随机生成 [0,max) 内int类型数字
func RandInt(max int) int {
	return rand.Intn(max)
}

// RandStr String returns a random string [a-z,A-Z,0-9] in the specified length
func RandStr(length int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return RandSomeStr(str, length)
}

// RandSomeStr 随机获取指定长度的指定字符串
func RandSomeStr(s string, length int) string {
	if StrIsEmpty(s) {
		return ""
	}

	letter := []rune(s)
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

// RandNumStr 随机生成指定长度的数字字符串
func RandNumStr(length int) string {
	// [用GO生成指定长度的随机字符串 - impressionw的博客 - CSDN博客](https://blog.csdn.net/impressionw/article/details/72765756)
	s := make([]int, length)
	//[Go生成随机数 - Cynhard的专栏 - CSDN博客](https://blog.csdn.net/u011304970/article/details/72721747)
	for i := range s {
		s[i] = rand.Intn(10)
	}

	// [go - One-liner to transform []int into string - Stack Overflow](https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string/37533144)
	a := strings.Trim(strings.Replace(fmt.Sprint(s), " ", "", -1), "[]")

	return a
}

// RandIntRange Int returns a random integer in range [min, max).
func RandIntRange(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandInt64 随机生成 [0,max) 内的Int64随机数
func RandInt64(max int64) int64 {
	return rand.Int63n(max)
}

// RandInt64Range 随机生成指定范围 [0,max) 内的随机数
func RandInt64Range(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}
