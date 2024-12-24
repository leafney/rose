/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose
 * @Date:        2024-12-24 11:18
 * @Description:
 */

package rose

import (
	"fmt"
	"strconv"
)

func AnyToStr(a interface{}) string {
	switch v := a.(type) {
	case nil:
		return ""
	case string:
		return v
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(int64(v.(int)), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(uint64(v.(uint)), 10)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", a)
	}
}

func AnyToInt(a interface{}) (int, error) {
	switch v := a.(type) {
	case int:
		return v, nil
	case int8, int16, int32, int64:
		return int(v.(int64)), nil // 统一转换为 int
	case string:
		return s2i(v)
	case float32, float64:
		return int(v.(float64)), nil // 将浮点数转换为整数
	case bool:
		if v {
			return 1, nil // true 转为 1
		}
		return 0, nil // false 转为 0
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

// 辅助函数：将字符串转换为整数
func s2i(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("cannot convert string to int: %s", s)
	}
	return i, nil
}

// AnyToInt64 将任意类型的值转换为 int64
func AnyToInt64(a interface{}) (int64, error) {
	switch v := a.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case string:
		return s2i64(v)
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case bool:
		if v {
			return 1, nil // true 转为 1
		}
		return 0, nil // false 转为 0
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

// 辅助函数：将字符串转换为 int64
func s2i64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot convert string to int64: %s", s)
	}
	return i, nil
}
