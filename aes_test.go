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
	pwdKey := "ABCDEFGHIJKLMNOP"
	origin := "Hello world"
	enStr, err := AesEncrypt(pwdKey, origin)
	t.Logf("RESULT [%v] err [%v]", enStr, err)
	deStr, err := AesDecrypt(pwdKey, enStr)
	t.Logf("RESULT [%v] err [%v]", deStr, err)
}
