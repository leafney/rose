/**
 * @Author:      leafney
 * @Date:        2021-10-22 17:44
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package xslice

import (
	"strconv"
	"strings"
)

func RmvStrItem(slice []string, items ...string) []string {
	isInEle := make(map[string]bool)
	for _, e1 := range items {
		isInEle[e1] = true
	}
	w := 0
	for _, e2 := range slice {
		if !isInEle[e2] {
			slice[w] = e2
			w++
		}
	}
	return slice[:w]
}

func RmvIntItem(slice []int, items ...int) []int {
	isInEle := make(map[int]bool)
	for _, e1 := range items {
		isInEle[e1] = true
	}
	w := 0
	for _, e2 := range slice {
		if !isInEle[e2] {
			slice[w] = e2
			w++
		}
	}
	return slice[:w]
}

func RmvInt64Item(slice []int64, items ...int64) []int64 {
	isInEle := make(map[int64]bool)
	for _, e1 := range items {
		isInEle[e1] = true
	}
	w := 0
	for _, e2 := range slice {
		if !isInEle[e2] {
			slice[w] = e2
			w++
		}
	}
	return slice[:w]
}

//
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

//
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

func SliceStrToStr(slice []string, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	return strings.Join(slice, sep)
}
