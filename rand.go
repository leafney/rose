package rose

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// String returns a random string ['a', 'z'] in the specified length
func RandStr(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

/*
生成指定长度的随机数字字符串
*/
func RandNumStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	// [用GO生成指定长度的随机字符串 - impressionw的博客 - CSDN博客](https://blog.csdn.net/impressionw/article/details/72765756)
	s := make([]int, n)
	//[Go生成随机数 - Cynhard的专栏 - CSDN博客](https://blog.csdn.net/u011304970/article/details/72721747)
	for i := range s {
		s[i] = rand.Intn(10)
	}

	// [go - One-liner to transform []int into string - Stack Overflow](https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string/37533144)
	a := strings.Trim(strings.Replace(fmt.Sprint(s), " ", "", -1), "[]")

	return a
}

// Int returns a random integer in range [min, max].
func RandIntRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func RandInt64(max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max)
}

func RandInt64Range(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	//if min >= max || min == 0 || max == 0 {
	//	return max
	//}
	return rand.Int63n(max-min) + min
}

/*
随机选择Slice中的一项
*/
func RandSliceItem(items []string) (string, error) {
	if len(items) == 0 {
		return "", errors.New("不能为空数组")
	} else if len(items) == 1 {
		return items[0], nil
	} else {
		rand.Seed(time.Now().UnixNano())
		// https://stackoverflow.com/questions/33994677/pick-a-random-value-from-a-go-slice
		return items[rand.Intn(len(items))], nil
	}
}

/*
洗牌算法
*/
func ShuffleForSliceInt(list []int) (newList []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range r.Perm(len(list)) {
		newList = append(newList, list[i])
	}
	return
}

/*
洗牌算法
*/
func ShuffleForSliceString(list []string) (newList []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range r.Perm(len(list)) {
		newList = append(newList, list[i])
	}
	return
}
