package str

import "testing"

func TestStrToInt(t *testing.T) {
	t.Log(StrToInt("3"))
}

func TestJoinString(t *testing.T) {
	t.Log(JoinString("aaa", "bbb", "ccc"))
}
