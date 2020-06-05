package rose

import (
	"hash/crc32"
	"strconv"
)

func CRC32Int(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func CRC32Str(str string) string {
	res := crc32.ChecksumIEEE([]byte(str))
	return strconv.FormatUint(uint64(res), 10)
}
