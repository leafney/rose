package rose

import (
	"encoding/hex"
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

// CRC32Hash CRC32 校验
func CRC32Hash(input string) string {
	// 使用 CRC32 算法计算校验和
	hash := crc32.NewIEEE()
	hash.Write([]byte(input))

	// 将计算出的 CRC32 校验和转为十六进制字符串
	return hex.EncodeToString(hash.Sum(nil))
}
