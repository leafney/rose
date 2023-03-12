/**
 * @Author:      leafney
 * @Date:        2023-03-12 20:00
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "sort"

type KeyValue struct {
	Key   string
	Value string
}

type sortedMap struct {
	pairs []KeyValue
}

func (sm *sortedMap) Len() int {
	return len(sm.pairs)
}

func (sm *sortedMap) Swap(i, j int) {
	sm.pairs[i], sm.pairs[j] = sm.pairs[j], sm.pairs[i]
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.pairs[i].Key < sm.pairs[j].Key
}

func SortedMap(m map[string]string) *sortedMap {
	pairs := make([]KeyValue, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, KeyValue{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})
	return &sortedMap{pairs: pairs}
}

func SortedMapNew() *sortedMap {
	return &sortedMap{pairs: make([]KeyValue, 0)}
}

func (sm *sortedMap) Get(key string) (string, bool) {
	//for _, pair := range sm.pairs {
	//	if pair.Key == key {
	//		return pair.Value, true
	//	}
	//}
	//return "", false

	/*
		By ChatGPT:
		在 Get() 方法中，你使用了线性搜索算法，这样会使得大型有序 map 的查找性能受到影响。因此，你可以尝试采用二分查找算法，这样可以将时间复杂度从线性减少到对数级别。
	*/

	// 优化，采用二分查找法
	idx := sort.Search(sm.Len(), func(i int) bool {
		return sm.pairs[i].Key >= key
	})
	if idx < sm.Len() && sm.pairs[idx].Key == key {
		return sm.pairs[idx].Value, true
	}
	return "", false
}

func (sm *sortedMap) Push(key, value string) {
	//sm.pairs = append(sm.pairs, KeyValue{key, value})
	//sort.Slice(sm.pairs, func(i, j int) bool {
	//	return sm.pairs[i].Key < sm.pairs[j].Key
	//})

	/*
		By ChatGPT:
		你的 Push() 方法在添加新键值对时，使用了一个额外的排序操作。这样会使得添加新键值对的操作变慢。相反，你可以使用二分查找算法来寻找新键值对应该插入的位置，并插入到对应位置。这样可以将添加新键值对的时间复杂度从 O(n) 降低到 O(log n)。
	*/

	// 优化，采用二分查找法
	idx := sort.Search(sm.Len(), func(i int) bool {
		return sm.pairs[i].Key >= key
	})
	sm.pairs = append(sm.pairs, KeyValue{})
	copy(sm.pairs[idx+1:], sm.pairs[idx:])
	sm.pairs[idx] = KeyValue{Key: key, Value: value}
}

func (sm *sortedMap) Remove(key string) {
	//for i, pair := range sm.pairs {
	//	if pair.Key == key {
	//		sm.pairs = append(sm.pairs[:i], sm.pairs[i+1:]...)
	//		break
	//	}
	//}

	// 优化，采用二分查找法
	idx := sort.Search(sm.Len(), func(i int) bool {
		return sm.pairs[i].Key >= key
	})
	if idx < sm.Len() && sm.pairs[idx].Key == key {
		sm.pairs = append(sm.pairs[:idx], sm.pairs[idx+1:]...)
	}
}
