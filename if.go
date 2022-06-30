/**
 * @Author:      leafney
 * @Date:        2022-06-30 15:20
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

// If 三元表达式实现
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// If 三元表达式实现
func IfStr(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

// If 三元表达式实现
func IfInt64(condition bool, trueVal, falseVal int64) int64 {
	if condition {
		return trueVal
	}
	return falseVal
}
