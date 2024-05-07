/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-05-07 20:28
 * @Description:
 */

package rose

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// RsaEncrypt RSA加密
// pemData：公钥内容；plainText：要加密的字符串数据；isBase64：返回结果是否需要经过 base64 编码
func RsaEncrypt(pemData, plainText string, isBase64 bool) (string, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return "", errors.New("failed to decode PEM block containing public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("failed to casting public key to RSA public key")
	}

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(plainText))
	if err != nil {
		return "", err
	}

	if isBase64 {
		return base64.StdEncoding.EncodeToString(cipherText), nil
	} else {
		return string(cipherText), nil
	}
}

// RsaDecrypt RSA解密
// pemData：私钥内容；cipherText：需要解密的字符串数据；isBase64：解密的数据是否需要 base64 解码
func RsaDecrypt(pemData, cipherText string, isBase64 bool) (string, error) {
	var err error
	var cipherData []byte
	if isBase64 {
		cipherData, err = base64.StdEncoding.DecodeString(cipherText)
		if err != nil {
			return "", err
		}
	} else {
		cipherData = []byte(cipherText)
	}

	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return "", errors.New("failed to decode PEM block containing public key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherData)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
