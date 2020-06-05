package rose

import "regexp"

// 正则匹配手机号格式是否正确(相对宽松的校验格式
func CheckPhoneFormatOK(phone string) bool {
	r, _ := regexp.Compile(`^1[3-9](\d{9})$`)
	return r.MatchString(phone)
}
