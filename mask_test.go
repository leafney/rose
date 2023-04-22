package rose

import "testing"

func TestMaskPhone(t *testing.T) {
	t.Log(MaskPhone("13012345678"))
	t.Log(MaskPwd("12345678"))
	t.Log(MaskBankCard("1234567887654321"))
	t.Log(MaskIDCard("123456789012345678"))
	t.Log(MaskEmail("12@qq.com"))
	t.Log(MaskStr("hello world", 0, 4, "@", 5))
}
