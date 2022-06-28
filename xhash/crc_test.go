/**
 * @Author:      leafney
 * @Date:        2022-04-25 16:15
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package xhash

import (
	"fmt"
	"testing"
)

func TestCrc32(t *testing.T) {
	t.Log(Crc32("12345"))

	for i := 0; i < 10; i++ {
		t.Log(CrcMod(fmt.Sprintf("%v", i), 3))
	}
}
