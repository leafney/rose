/**
 * @Author:      leafney
 * @Date:        2022-06-30 15:29
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "testing"

func TestCRC32(t *testing.T) {
	t.Log(Crc32Val("hello"))
	t.Log(Crc32("hello"))
	t.Log(Crc32Mod("hello", 10))
}
