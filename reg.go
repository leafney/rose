package rose

import "regexp"

// IsPhone 正则匹配手机号格式是否正确(相对宽松的校验格式
func IsPhone(phone string) bool {
	pattern := `^1[3-9](\d{9})$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(phone)
}

// IsEmail 正则匹配邮箱格式是否正确
func IsEmail(email string) bool {
	pattern := `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// PhoneNumStar 隐藏手机号码中间4位
func PhoneNumStar(phone string) string {
	if IsPhone(phone) {
		pattern := `^(\d{3})\d{4}(\d{4})$`
		re := regexp.MustCompile(pattern)
		return re.ReplaceAllString(phone, "$1****$2")
	}
	return phone
}

// PhoneTailNum 获取手机尾号后4位
func PhoneTailNum(phone string) string {
	if IsPhone(phone) {
		return phone[len(phone)-4:]
	}
	return phone
}
