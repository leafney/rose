package rose

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func MustBool(s string, defVal ...bool) bool {
	if s == "" {
		return getBoolDefault(defVal...)
	}

	b, err := strconv.ParseBool(strings.TrimSpace(s))
	if err != nil {
		return getBoolDefault(defVal...)
	}

	return b
}

func getBoolDefault(defVal ...bool) bool {
	if len(defVal) > 0 {
		return defVal[0]
	}
	return false
}

// ************

func getFloatDefault(defVals ...float64) float64 {
	if len(defVals) > 0 {
		return defVals[0]
	}
	return 0.0
}

func mustFloat(s string, defVals ...float64) float64 {
	if s == "" {
		return getFloatDefault()
	}

	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return getFloatDefault()
	}

	return f
}

func MustFloat(inter interface{}, defaultVals ...float64) float64 {
	switch v := inter.(type) {
	case float64:
		return v
	case string:
		return mustFloat(v, defaultVals...)
	case int64:
		return float64(v)
	case float32:
		return float64(v)
	default:
		return getFloatDefault(defaultVals...)
	}
}

// MustInt 字符串转int
func MustInt(s string, defVal ...int) int {
	getDefault := func() int {
		if len(defVal) > 0 {
			return defVal[0]
		}
		return 0
	}

	if s == "" {
		return getDefault()
	}

	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		msg := "goutils MustInt strconv.Atoi error:" + err.Error()
		// 加上文件调用和行号
		_, callerFile, line, ok := runtime.Caller(1)
		if ok {
			msg += fmt.Sprintf("file:%s,line:%d", callerFile, line)
		}
		return getDefault()
	}

	return i
}

// MustInt64 字符串转int64
func MustInt64(s string, defVal ...int64) int64 {
	getDefault := func() int64 {
		if len(defVal) > 0 {
			return defVal[0]
		}
		return 0
	}

	if s == "" {
		return getDefault()
	}

	i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		msg := "goutils MustInt64 strconv.ParseInt error:" + err.Error()
		// 加上文件调用和行号
		_, callerFile, line, ok := runtime.Caller(1)
		if ok {
			msg += fmt.Sprintf("file:%s,line:%d", callerFile, line)
		}
		return getDefault()
	}

	return i
}
