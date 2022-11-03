package rose

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// 对字符串进行Base64编码
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// 对字符串进行Base64解码
func Base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func Base64URLEncode(str string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(str))
}

func Base64URLDecode(str string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func Base64HmacSha256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
