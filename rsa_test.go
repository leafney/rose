/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-05-07 20:28
 * @Description:
 */

package rose

import "testing"

func TestRsaEncrypt(t *testing.T) {
	// RSA 生成密钥对
	pri, pub, err := RsaGenerateKey(2048)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("private [%v]", string(pri))
	t.Logf("public [%v]", string(pub))

	//	RSA加密
	str := "HellO World"
	encrypt, _ := RsaEncryptByte(pub, str)
	t.Logf("encrypt [%v]", string(encrypt))

	//	RSA解密
	decrypt, _ := RsaDecryptByte(pri, string(encrypt))
	t.Logf("decrypt [%v]", string(decrypt))
}
