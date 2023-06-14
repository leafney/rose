/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-06-14 17:28
 * @Description:
 */

package rose

import "testing"

func TestMapUnionStr(t *testing.T) {
	m1 := map[string]string{
		"name": "hello",
	}
	m2 := map[string]string{
		"name":    "world",
		"address": "beijing",
	}

	m3 := MapUnionStr(m1, m2)
	t.Log(m3)
}
