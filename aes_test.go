/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2023-06-07 13:47
 * @Description:
 */

package rose

import "testing"

func TestAesEncrypt(t *testing.T) {
	pwdKey := "ABCDEFGHIJKLMNOPXY"
	origin := "Hello world"

	//enStr, err := AesEncrypt([]byte(pwdKey), []byte(origin))
	//t.Logf("RESULT [%v] err [%v]", string(enStr), err)
	//deStr, err := AesDecrypt([]byte(pwdKey), enStr)
	//t.Logf("RESULT [%v] err [%v]", string(deStr), err)

	//enStr, err := AesEncryptHex(pwdKey, origin)
	//t.Logf("RESULT [%v] err [%v]", enStr, err)
	//deStr, err := AesDecryptHex(pwdKey, enStr)
	//t.Logf("RESULT [%v] err [%v]", deStr, err)

	enStr, err := AesEncryptBase64(pwdKey, origin)
	t.Logf("RESULT [%v] err [%v]", enStr, err)
	deStr, err := AesDecryptBase64(pwdKey, enStr)
	t.Logf("RESULT [%v] err [%v]", deStr, err)
}
