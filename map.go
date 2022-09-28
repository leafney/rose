/**
 * @Author:      leafney
 * @Date:        2022-07-17 15:03
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "sort"

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
