package rose

import "strconv"

func BoolToStr(b bool) string {
	return strconv.FormatBool(b)
}

func BoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func BoolToInt64(b bool) int64 {
	return int64(BoolToInt(b))
}
