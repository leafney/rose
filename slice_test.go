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
