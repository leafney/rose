/**
 * @Author:      leafney
 * @Date:        2022-07-17 14:44
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

func Sha256(data string) string {
	sha256 := sha256.New()
	sha256.Write([]byte(data))
	return hex.EncodeToString(sha256.Sum([]byte("")))
}

func Sha512(data string) string {
	sha512 := sha512.New()
	sha512.Write([]byte(data))
	return hex.EncodeToString(sha512.Sum([]byte("")))
}
