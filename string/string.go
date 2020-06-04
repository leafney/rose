package string

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func Md5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}

func StringToInt64(s string) int64 {
	r, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		return 0
	}
	return r
}

func StringToFloat64(s string) float64 {
	r, e := strconv.ParseFloat(s, 64)
	if e != nil {
		return 0
	}
	return r
}
