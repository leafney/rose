package rose

import (
	"fmt"
	"strings"
)

// MaskPhone 手机号 保留前3后4
func MaskPhone(phone string) string {
	return MaskStr(phone, 3, 4, "*", 4)
}

// MaskEmail 邮箱 仅用户名部分 保留前1后1
func MaskEmail(email string) string {
	if !strings.Contains(email, "@") {
		return email
	}
	sps := strings.Split(email, "@")
	if len(sps[0]) < 1 || len(sps[1]) < 1 {
		return email
	}
	return fmt.Sprintf("%s@%s", MaskStr(sps[0], 1, 1, "*", 4), sps[1])
}

// MaskPwd 密码
func MaskPwd(pwd string) string {
	return "********"
}

// MaskBankCard 银行卡号 保留前4后4
func MaskBankCard(card string) string {
	return MaskStr(card, 4, 4, "*", 4)
}

// MaskIDCard 身份证号 保留前3后4
func MaskIDCard(id string) string {
	return MaskStr(id, 3, 4, "*", 4)
}

// MaskStr 对字符串掩码处理，保留前front，后after，使用rep替换中间内容，并指定rep的数量
func MaskStr(s string, front, after int, rep string, count int) string {
	if s == "" {
		return ""
	}

	if len(s) < front+after {
		return s
	}

	var builder strings.Builder
	builder.WriteString(s[:front])
	builder.WriteString(strings.Repeat(rep, count))
	builder.WriteString(s[len(s)-after:])
	return builder.String()
}
