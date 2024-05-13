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
	"os"
	"path/filepath"
)

// RSAGenerateKeySave 生成 RSA 密钥对并保存到指定目录
// savePath：保存到的目录路径
func RSAGenerateKeySave(bits int, keyName, savePath string) error {
	// 确保保存的目录存在
	if err := DEnsurePathExist(savePath); err != nil {
		return err
	}

	private, public, err := RsaGenerateKey(bits)
	if err != nil {
		return err
	}

	privateSavePath := filepath.Join(savePath, fmt.Sprintf("%s.pem", keyName))
	publicSavePath := filepath.Join(savePath, fmt.Sprintf("%s_pub.pem", keyName))

	if err := os.WriteFile(privateSavePath, private, 0644); err != nil {
		return err
	}

	if err := os.WriteFile(publicSavePath, public, 0644); err != nil {
		return err
	}

	return nil
}

// RsaGenerateKey 生成 RSA 密钥对
func RsaGenerateKey(bits int) (private, public []byte, err error) {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	// 将私钥转换为PEM格式字符串
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509PrivateKey,
	})

	// 将公钥转换为PEM格式字符串
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return privateKeyPEM, publicKeyPEM, nil
}

// RsaEncrypt RSA加密
// pemData：公钥内容；plainText：要加密的字符串数据；
func RsaEncrypt(pemData, plainText string) (string, error) {
	return RsaEncryptB64([]byte(pemData), plainText, false)
}

// RsaDecrypt RSA解密
// pemData：私钥内容；cipherText：需要解密的字符串数据；
func RsaDecrypt(pemData, cipherText string) (string, error) {
	return RsaDecryptB64([]byte(pemData), cipherText, false)
}

// RsaEncryptB64 RSA加密，公钥加密
// pemData：公钥内容；plainText：要加密的字符串数据；isBase64：返回结果是否需要经过 base64 编码
func RsaEncryptB64(pemData []byte, plainText string, isBase64 bool) (string, error) {
	cipherByte, err := RsaEncryptByte(pemData, plainText)
	if err != nil {
		return "", err
	}

	if isBase64 {
		return base64.StdEncoding.EncodeToString(cipherByte), nil
	} else {
		return string(cipherByte), nil
	}
}

// RsaDecryptB64 RSA解密，私钥解密
// pemData：私钥内容；cipherText：需要解密的字符串数据；isBase64：解密的数据是否需要 base64 解码
func RsaDecryptB64(pemData []byte, cipherText string, isBase64 bool) (string, error) {
	var cipherData string
	if isBase64 {
		cipherByte, err := base64.StdEncoding.DecodeString(cipherText)
		if err != nil {
			return "", err
		}
		cipherData = string(cipherByte)
	} else {
		cipherData = cipherText
	}

	plainByte, err := RsaDecryptByte(pemData, cipherData)
	return string(plainByte), err
}

// RsaEncryptByte RSA加密，公钥加密
func RsaEncryptByte(pemData []byte, plainText string) ([]byte, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to casting public key to RSA public key")
	}

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(plainText))
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

// RsaDecryptByte RSA解密，私钥解密
func RsaDecryptByte(pemData []byte, cipherText string) ([]byte, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(cipherText))
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
