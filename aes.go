package rose

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

/*
	key为16,24,32位字符串，分别对应AES-128，AES-192，AES-256 加密方法
	- AesCBCEncrypt
	- AesCBCDecrypt
	- AesCFBEncrypt
	- AesCFBDecrypt
	- AesECBEncrypt
	- AesECBDecrypt
*/

func AesEncrypt(key, data string) (string, error) {
	keyBt := aesKeyPadding([]byte(key))
	dataBt := []byte(data)

	block, err := aes.NewCipher(keyBt)
	if err != nil {
		return "", err
	}

	// 加密块大小
	blockSize := block.BlockSize()
	//	填充
	encryptBytes := pkcs7Padding(dataBt, blockSize)
	// 初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	// 使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, keyBt[:blockSize])
	// 执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func AesDecrypt(key, data string) (string, error) {
	dataBt, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	keyBt := aesKeyPadding([]byte(key))

	block, err := aes.NewCipher(keyBt)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBt[:blockSize])
	crypted := make([]byte, len(dataBt))
	blockMode.CryptBlocks(crypted, dataBt)
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return "", err
	}
	return string(crypted), nil
}

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
