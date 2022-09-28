/**
 * @Author:      leafney
 * @Date:        2022-06-30 10:56
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Remove empty items from slice
func SliceStrRmvEmpty(sls []string) []string {
	newSls := make([]string, 0)
	for k, v := range sls {
		if len(v) > 0 {
			newSls = append(newSls, sls[k])
		}
	}
	return newSls
}

// Divide a slice into a specified number of slice groups
func SliceStrToArray(arr []string, size int) (result [][]string) {
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

// SliceReverse 切片反转
func SliceReverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// SliceRmvDuplicates 切片去重
func SliceRmvDuplicates(slices []string) []string {
	result := make([]string, 0, len(slices))
	temp := map[string]struct{}{}
	for _, item := range slices {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// SliceRmvSubSlice slice1中去除slice2中有的item
func SliceRmvSubSlice(slice1, slice2 []string) (newSlice []string) {
	for _, s1 := range slice1 {
		flag := true
		for _, s2 := range slice2 {
			if s1 == s2 {
				flag = false
				break
			}
		}
		if flag {
			newSlice = append(newSlice, s1)
		}
	}
	return
}

// SliceExistStr 判断slice中是否存在
func SliceExistStr(slices []string, val string) bool {
	for _, v := range slices {
		if val == v {
			return true
		}
	}
	return false
}

func SliceExistInt64(slices []int64, val int64) bool {
	for _, v := range slices {
		if val == v {
			return true
		}
	}
	return false
}

// SliceRandomItemStr 随机选择Slice中的一项
func SliceRandomItemStr(slices []string) string {
	if len(slices) == 0 {
		return ""
	} else if len(slices) == 1 {
		return slices[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		// https://stackoverflow.com/questions/33994677/pick-a-random-value-from-a-go-slice
		return slices[rand.Intn(len(slices))]
	}
}

// SliceRandomItemInt64 随机选择slice中的一项
func SliceRandomItemInt64(slices []int64) int64 {
	if len(slices) == 0 {
		return 0
	} else if len(slices) == 1 {
		return slices[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		// https://stackoverflow.com/questions/33994677/pick-a-random-value-from-a-go-slice
		return slices[rand.Intn(len(slices))]
	}
}

// 洗牌算法
func SliceShuffleInt(list []int) (newList []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range r.Perm(len(list)) {
		newList = append(newList, list[i])
	}
	return
}

// 洗牌算法
func SliceShuffleInt64(list []int64) (newList []int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range r.Perm(len(list)) {
		newList = append(newList, list[i])
	}
	return
}

// 洗牌算法
func SliceShuffleStr(list []string) (newList []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range r.Perm(len(list)) {
		newList = append(newList, list[i])
	}
	return
}

// SliceIntToStr 将int类型slice拼接成指定字符分隔的字符串
func SliceIntToStr(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	res := make([]string, len(slice))
	for i, v := range slice {
		res[i] = strconv.Itoa(v)
	}
	return strings.Join(res, sep)
}

// SliceInt64ToStr 将int64类型slice拼接成指定字符分隔的字符串
func SliceInt64ToStr(slice []int64, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	res := make([]string, len(slice))
	for i, v := range slice {
		res[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(res, sep)
}

// SliceStrToStr 将slice拼接成指定字符分隔的字符串
func SliceStrToStr(slice []string, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	return strings.Join(slice, sep)
}

// SlicePage slice分页
func SlicePage(page, pageSize, defSize int64, nums int) (sliceStart, sliceEnd int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = defSize
	}
	if int(pageSize) > nums {
		return 0, nums
	}
	pageCount := int64(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}

	sliceStart = int((page - 1) * pageSize)
	sliceEnd = sliceStart + int(pageSize)

	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd
}
