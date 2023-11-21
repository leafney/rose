package rose

import "regexp"

// IsPhone 正则匹配手机号格式是否正确(相对宽松的校验格式
func RegIsPhone(phone string) bool {
	pattern := `^1[3-9](\d{9})$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(phone)
}

// IsEmail 正则匹配邮箱格式是否正确
func RegIsEmail(email string) bool {
	pattern := `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// PhoneNumMask 手机号码中间4位掩码处理
func PhoneNumMask(phone string) string {
	if RegIsPhone(phone) {
		pattern := `^(\d{3})\d{4}(\d{4})$`
		re := regexp.MustCompile(pattern)
		return re.ReplaceAllString(phone, "$1****$2")
	}
	return phone
}

// PhoneNumTail 获取手机尾号后4位
func PhoneNumTail(phone string) string {
	if RegIsPhone(phone) {
		return phone[len(phone)-4:]
	}
	return phone
}

// RegIsNormalString 校验字符串是否只包含字母、数字和下划线
func RegIsNormalString(input string) bool {
	// 定义正则表达式，只允许字母、数字和下划线
	reg := regexp.MustCompile(`^[A-Za-z0-9_]+$`)

	// 校验字符串是否符合正则表达式
	if !reg.MatchString(input) {
		return false
	}

	// 校验字符串是否包含空格
	if RegIsContainsSpace(input) {
		return false
	}

	return true
}

// RegIsContainsSpace 判断字符串是否包含空格
func RegIsContainsSpace(input string) bool {
	// 定义正则表达式，匹配空格
	reg := regexp.MustCompile(`\s`)

	// 校验字符串是否包含空格
	return reg.MatchString(input)
}

// RegIsChinese 校验字符串是否只包含中文字符
func RegIsChinese(input string) bool {
	// 定义正则表达式，匹配中文字符
	reg := regexp.MustCompile(`^[\p{Han}]+$`)

	// 校验字符串是否只包含中文字符
	return reg.MatchString(input)
}
