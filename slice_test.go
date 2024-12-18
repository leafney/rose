/**
 * @Author:      leafney
 * @Date:        2022-06-30 11:02
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "testing"

func TestSliceStrRmvEmpty(t *testing.T) {

}

func TestSliceRmvSubStr(t *testing.T) {
	slc1 := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "mmm", "nnn"}
	slc2 := []string{"ddd", "fff", "nnn"}
	for i := 0; i < 100; i++ {
		t.Log(SliceRmvSubSlice(slc1, slc2))
	}
}

func BenchmarkSliceRmvSubStr(b *testing.B) {
	slc1 := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "mmm", "nnn"}
	slc2 := []string{"ddd", "fff", "nnn"}

	for i := 0; i < b.N; i++ {
		b.Log(SliceRmvSubSlice(slc1, slc2))
	}
}

func TestSliceRmvSubSlice2(t *testing.T) {
	slc1 := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "mmm", "nnn"}
	slc2 := []string{"ddd", "fff", "nnn"}

	r1 := SliceRmvSubSlice2(slc1, slc2)
	t.Log(r1, len(r1))
}

type Item struct {
	ID   string
	Name string
}

func TestSliceRmvDuplicates(t *testing.T) {
	items := []Item{
		{ID: "1", Name: "Item1"},
		{ID: "2", Name: "Item2"},
		{ID: "1", Name: "Item1 Duplicate"},
		{ID: "3", Name: "Item3"},
		{ID: "2", Name: "Item2 Duplicate"},
	}

	uniqueItems := SliceRmvDuplicates(items, func(item Item) string {
		return item.ID // 获取 ID 字段来判断重复
	})

	for _, item := range uniqueItems {
		t.Logf("id %s name %s", item.ID, item.Name)
	}
}
