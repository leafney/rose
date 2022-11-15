package rose

import "testing"

func TestStr(t *testing.T) {
	t.Log(StrToInt("3"))

	t.Log(StrRemoveAny("hello world", "l", "o"))

	t.Log(StrPrefixAny("hello", "H"))
}
