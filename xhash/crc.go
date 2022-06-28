/**
 * @Author:      leafney
 * @Date:        2022-04-25 15:43
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package xhash

import (
	"github.com/leafney/rose/xconv"
	"hash/crc32"
)

func Crc32(key string) string {
	value := crc32.ChecksumIEEE([]byte(key))
	mVal := value % 32
	return xconv.UInt32ToStr(mVal)
}

func CrcMod(key string, mod uint32) string {
	value := crc32.ChecksumIEEE([]byte(key))
	mVal := value % mod
	return xconv.UInt32ToStr(mVal)
}
