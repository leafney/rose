package rose

import "testing"

func TestTimeNow10(t *testing.T) {
	t.Log(TNowS())
}

func TestTimeNow13(t *testing.T) {
	t.Log(TNowMs())
}

func TestParseGMTTime(t *testing.T) {
	tt, _ := ParseGMTTimeOfRFC1123("Mon, 20 Jul 2020 06:09:21 GMT")
	t.Log(tt.Unix())
	t.Log(GetTimeFromUnix(tt.Unix()))
}
