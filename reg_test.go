/**
 * @Author:      leafney
 * @Date:        2022-11-06 01:54
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

import "testing"

func TestPhoneNumStar(t *testing.T) {
	phone := "17092021576"
	t.Log(RegIsPhone(phone))
	t.Log(PhoneNumMask(phone))
	t.Log(PhoneNumTail(phone))
	t.Log(RegIsNormalString("hello@qq"))
	t.Log(RegIsChinese("你好"), RegIsChinese("he 嘿"))
}
