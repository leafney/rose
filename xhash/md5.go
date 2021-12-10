package xhash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5Str(s string) string {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(s))
	return hex.EncodeToString(hashMd5.Sum(nil))
}

func Md5Buf(buf []byte) string {
	hashMd5 := md5.New()
	hashMd5.Write(buf)
	return hex.EncodeToString(hashMd5.Sum(nil))
}

//
func Md5File(reader io.Reader) string {
	var buf = make([]byte, 4096)
	hashMd5 := md5.New()
	for {
		n, err := reader.Read(buf)
		if err == io.EOF && n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			break
		}

		hashMd5.Write(buf[:n])
	}
	return hex.EncodeToString(hashMd5.Sum(nil))
}
