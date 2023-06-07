package rose

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

/*
	key为16,24,32位字符串，分别对应AES-128，AES-192，AES-256 加密方法
	- AesCBCEncrypt (AesEncrypt)
	- AesCBCDecrypt (AesDecrypt)
	- AesCFBEncrypt
	- AesCFBDecrypt
	- AesECBEncrypt
	- AesECBDecrypt
*/

func aesKeyPadding(key []byte) []byte {
	k := len(key)
	count := 0
	switch true {
	case k <= 16:
		count = 16 - k
	case k <= 24:
		count = 24 - k
	case k <= 32:
		count = 32 - k
	default:
		return key[:32]
	}

	if count == 0 {
		return key
	}

	return append(key, bytes.Repeat([]byte{0}, count)...)
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return data, errors.New("UnPadding data error")
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// --------

func AesCBCEncrypt(key, data []byte) ([]byte, error) {
	theKey := aesKeyPadding(key)
	block, err := aes.NewCipher(theKey)
	if err != nil {
		return nil, err
	}

	// 加密块大小
	blockSize := block.BlockSize()
	//	填充
	encryptBytes := pkcs7Padding(data, blockSize)
	// 初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	// 使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, theKey[:blockSize])
	// 执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)

	return crypted, nil
}

func AesCBCDecrypt(key, data []byte) ([]byte, error) {
	theKey := aesKeyPadding(key)
	block, err := aes.NewCipher(theKey)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, theKey[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

// AesEncrypt alias AesEncryptBase64
func AesEncrypt(key, data string) string {
	res, _ := AesEncryptBase64(key, data)
	return res
}

// AesDecrypt alias AesDecryptBase64
func AesDecrypt(key, data string) string {
	res, _ := AesDecryptBase64(key, data)
	return res
}

func AesEncryptHex(key, data string) (string, error) {
	res, err := AesCBCEncrypt([]byte(key), []byte(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(res), nil
}

func AesDecryptHex(key, data string) (string, error) {
	theData, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}
	res, err := AesCBCDecrypt([]byte(key), theData)
	return string(res), err
}

func AesEncryptBase64(key, data string) (string, error) {
	res, err := AesCBCEncrypt([]byte(key), []byte(data))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

func AesDecryptBase64(key, data string) (string, error) {
	theData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	res, err := AesCBCDecrypt([]byte(key), theData)
	return string(res), err
}
