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
	t.Log(IsPhone(phone))
	t.Log(PhoneNumStar(phone))
	t.Log(PhoneNumEnd(phone))
}
