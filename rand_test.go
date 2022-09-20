/**
 * @Author:      leafney
 * @Date:        2022-09-20 19:07
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "testing"

func TestRandIntRange(t *testing.T) {

	for i := 0; i < 10; i++ {
		t.Log(RandIntRange(0, 5))
	}
}

func TestRandSomeStr(t *testing.T) {
	t.Log(RandSomeStr("abcdefghi", 2))
}
