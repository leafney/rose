/**
 * @Author:      leafney
 * @Date:        2022-06-30 15:29
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import (
	"fmt"
	"testing"
)

func TestCRC32(t *testing.T) {
	//t.Log(Crc32Val("hello"))
	//t.Log(Crc32("hello"))
	for i := 0; i < 100; i++ {
		t.Log(Crc32Mod(fmt.Sprintf("%v", i), 32))
	}

}
