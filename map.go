/**
 * @Author:      leafney
 * @Date:        2022-07-17 15:03
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import (
	"sort"
	"strconv"
)

// MapInt64SortByValue 将map中的key按照value大小排序后返回
func MapInt64SortByValue(m map[string]int64, asc bool) (nk []string) {
	type PairKV struct {
		Key   string
		Value int64
	}

	var res []PairKV
	for k, v := range m {
		res = append(res, PairKV{k, v})
	}

	sort.Slice(res, func(i, j int) bool {
		if asc {
			return res[i].Value < res[i].Value // 升序
		} else {
			return res[i].Value > res[i].Value // 降序
		}
	})

	for _, re := range res {
		nk = append(nk, re.Key)
	}
	return nk
}

// MapFloat64SortByValue 将map中的key按照value大小排序后返回
func MapFloat64SortByValue(m map[string]float64, asc bool) (nk []string) {
	type PairKV struct {
		Key   string
		Value float64
	}

	var res []PairKV
	for k, v := range m {
		res = append(res, PairKV{k, v})
	}

	sort.Slice(res, func(i, j int) bool {
		if asc {
			return res[i].Value < res[i].Value // 升序
		} else {
			return res[i].Value > res[i].Value // 降序
		}
	})

	for _, re := range res {
		nk = append(nk, re.Key)
	}
	return nk
}

// MapGetValueDef 获取map中key的值或返回指定的默认值
func MapGetValueDef(m map[string]string, key string, def string) string {
	if val, ok := m[key]; ok {
		return val
	}
	return def
}

// MapInterfaceToStr
func MapInterfaceToStr(m map[string]interface{}) map[string]string {
	res := make(map[string]string)
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			res[k] = vv
		case float64:
			res[k] = strconv.FormatFloat(vv, 'f', -1, 64)
		case int64:
			res[k] = strconv.FormatInt(vv, 10)
		case int:
			res[k] = strconv.Itoa(vv)
		case bool:
			res[k] = strconv.FormatBool(vv)
		default:
			//
		}
	}
	return res
}
