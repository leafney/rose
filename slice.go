/**
 * @Author:      leafney
 * @Date:        2022-06-30 10:56
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// SliceStrRmvEmpty Remove empty items from slice 移除空项
func SliceStrRmvEmpty(sls []string) []string {
	newSls := make([]string, 0)
	for k, v := range sls {
		if len(v) > 0 {
			newSls = append(newSls, sls[k])
		}
	}
	return newSls
}

// SliceStrRmvOne 从 slice 中移除指定项
func SliceStrRmvOne(sls []string, value string) []string {
	for i := 0; i < len(sls); i++ {
		if sls[i] == value {
			sls = append(sls[:i], sls[i+1:]...)
			i--
		}
	}
	return sls
}

// SliceStrToArray Divide a slice into a specified number of slice groups
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
func SliceReverse(slices interface{}) {
	sort.SliceStable(slices, func(i, j int) bool {
		return true
	})
}

// SliceStrReverse 切片反转
func SliceStrReverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// SliceRmvStrDuplicates 字符串切片去重
func SliceRmvStrDuplicates(slices []string) []string {
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

// SliceRmvSubSlice2  从 slice1 中移除 slice2 中存在的子项(效率高)
func SliceRmvSubSlice2(slice1, slice2 []string) []string {
	// 创建一个 map 存储要移除的元素
	removeMap := make(map[string]bool, len(slice2))
	for _, s2 := range slice2 {
		removeMap[s2] = true
	}

	// 创建结果切片,预先分配足够的内存空间
	result := make([]string, 0, len(slice1))

	// 遍历源切片,检查每个元素是否在要移除的 map 中
	for _, s1 := range slice1 {
		if !removeMap[s1] {
			result = append(result, s1)
		}
	}

	return result
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

// SliceSortByLength 将slice中字符串按照长度排序
// asc 由小到大；desc 由大到小
func SliceSortByLength(slices []string, asc bool) []string {
	sort.Slice(slices, func(i, j int) bool {
		if asc {
			return len(slices[i]) < len(slices[j])
		} else {
			return len(slices[i]) > len(slices[j])
		}
	})
	return slices
}

// SliceShuffleInt 洗牌算法
func SliceShuffleInt(list []int) (newList []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range r.Perm(len(list)) {
		newList = append(newList, list[i])
	}
	return
}

// SliceShuffleInt64 洗牌算法
func SliceShuffleInt64(list []int64) (newList []int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range r.Perm(len(list)) {
		newList = append(newList, list[i])
	}
	return
}

// SliceShuffleStr 洗牌算法
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

// SliceSplitPage 对slice分页
// page: 设置当前页码；pageSize: 设置获取每页条数；defSize：默认每页条数；count: 总数据条数
func SliceSplitPage(page, pageSize, defSize int64, count int) (sliceStart, sliceEnd int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = defSize
	}
	if int(pageSize) > count {
		return 0, count
	}
	pageCount := int64(math.Ceil(float64(count) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}

	sliceStart = int((page - 1) * pageSize)
	sliceEnd = sliceStart + int(pageSize)

	if sliceEnd > count {
		sliceEnd = count
	}
	return sliceStart, sliceEnd
}

// SlicePageToOffsetInt 将集合Page分页方式转换为Offset分页方式
func SlicePageToOffsetInt(page, size int) (offset, count int) {
	// 计算偏移量 offset 和每页显示数量 count
	offset = (page - 1) * size
	count = size
	return
}

// SlicePageToOffsetInt64 将集合Page分页方式转换为Offset分页方式
func SlicePageToOffsetInt64(page, size int64) (offset, count int64) {
	// 计算偏移量 offset 和每页显示数量 count
	offset = (page - 1) * size
	count = size
	return
}

// SliceGroupStr 将字符串切片分成若干组，每个组中元素个数为num个；不足一组时最后一组中包含剩余的所有元素
func SliceGroupStr(arr []string, num int64) [][]string {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]string{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]string, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

// SliceGroupInt64 将字符串切片分成若干组，每个组中元素个数为num个；不足一组时最后一组中包含剩余的所有元素
func SliceGroupInt64(arr []int64, num int64) [][]int64 {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]int64{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]int64, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

func SliceIntToMap(arr []int, prefix string) map[string]int {
	res := make(map[string]int, len(arr))
	for k, v := range arr {
		key := fmt.Sprintf("%s%d", prefix, k+1)
		res[key] = v
	}
	return res
}

func SliceInt64ToMap(arr []int64, prefix string) map[string]int64 {
	res := make(map[string]int64, len(arr))
	for k, v := range arr {
		key := fmt.Sprintf("%s%d", prefix, k+1)
		res[key] = v
	}
	return res
}

func SliceStrToMap(arr []string, prefix string) map[string]string {
	res := make(map[string]string, len(arr))
	for k, v := range arr {
		key := fmt.Sprintf("%s%d", prefix, k+1)
		res[key] = v
	}
	return res
}

// SliceGetElementByIndex 获取指定索引元素，如果索引越界则返回错误
func SliceGetElementByIndex[T any](slice []T, index int) (T, error) {
	var zeroValue T
	if index < 0 || index >= len(slice) {
		return zeroValue, fmt.Errorf("索引 %d 越界，切片长度为 %d", index, len(slice))
	}
	return slice[index], nil
}

// SliceGetElementValByIndex 获取指定索引元素，如果索引越界则返回类型零值
func SliceGetElementValByIndex[T any](slice []T, index int) T {
	val, _ := SliceGetElementByIndex(slice, index)
	return val
}

//func SliceContainsOne[T any](slice []T, val T) bool {
//	for _, v := range slice {
//		if val == v {
//			return true
//		}
//	}
//	return false
//}

func SliceContainsStr(slice []string, val string) bool {
	return SliceExistStr(slice, val)
}

func SliceContainsInt(slice []int, val int) bool {
	return SliceExistInt(slice, val)
}

func SliceContainsInt64(slice []int64, val int64) bool {
	return SliceExistInt64(slice, val)
}

// SliceContainsSliceAnyOne 判断 slice A 中是否存在 slice B 中的任意元素
func SliceContainsSliceAnyOne(a []string, b []string) bool {
	// 将slice B中的元素存入map中
	bMap := make(map[string]bool)
	for _, item := range b {
		bMap[item] = true
	}

	// 检查slice A中的元素是否在map中
	for _, item := range a {
		if _, exists := bMap[item]; exists {
			return true
		}
	}

	return false
}

// SliceInclusionRelationship 判断两个切片之间的包含关系（获取共有的元素、仅在第一个数组中存在的元素和仅在第二个数组中存在的元素
func SliceInclusionRelationship[T comparable](A, B []T) (share, left, right []T) {
	// 创建一个map来存储A中的元素
	aMap := make(map[T]bool)
	for _, item := range A {
		aMap[item] = true
	}

	// 遍历B中的元素，如果元素在aMap中存在，则将其添加到share中，否则添加到right中
	for _, item := range B {
		if _, ok := aMap[item]; ok {
			share = append(share, item)
			delete(aMap, item) // 从aMap中删除已经匹配的元素
		} else {
			right = append(right, item)
		}
	}

	// 将aMap中剩余的元素添加到left中
	for item := range aMap {
		left = append(left, item)
	}

	return share, left, right
}

// SliceRmvDuplicates 泛型方法，移除切片中重复的元素
func SliceRmvDuplicates[T any, K comparable](items []T, getKey func(T) K) []T {
	seen := make(map[K]struct{}) // 用来存储已经出现的字段值
	var result []T

	for _, item := range items {
		key := getKey(item)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// Deprecated

// Deprecated: Use SliceContainsStr instead.
// SliceExistStr 判断slice中是否存在
func SliceExistStr(slices []string, val string) bool {
	for _, v := range slices {
		if val == v {
			return true
		}
	}
	return false
}

// Deprecated: Use SliceContainsInt64 instead.
// SliceExistInt64 判断slice中是否存在
func SliceExistInt64(slices []int64, val int64) bool {
	for _, v := range slices {
		if val == v {
			return true
		}
	}
	return false
}

// Deprecated: Use SliceContainsInt instead.
// SliceExistInt 判断slice中是否存在
func SliceExistInt(slices []int, val int) bool {
	for _, v := range slices {
		if val == v {
			return true
		}
	}
	return false
}
