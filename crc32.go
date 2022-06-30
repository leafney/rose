package rose

import (
	"hash/crc32"
)

func crc32Int(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func Crc32Val(key string) string {
	val := crc32Int(key)
	return UInt32ToStr(val)
}

func Crc32(key string) string {
	val := crc32Int(key)
	nVal := val % 32
	return UInt32ToStr(nVal)
}

func Crc32Mod(key string, mod uint32) string {
	val := crc32Int(key)
	nVal := val % mod
	return UInt32ToStr(nVal)
}
