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
	t.Log(TUnixSToFormat(tt.Unix(), TFDateTime))
}

func TestTSecByDay(t *testing.T) {
	//tt := TSecByMin(5)
	//tt := TSecByDay(3)
	tt := TSecByMonth(2)
	t.Log(tt)
}

func TestTSecToAdd(t *testing.T) {
	//t.Log(TUnixSToDateTime(TSecToAdd(1, 1, 0)))

	//t.Log(DelayTimeToTomorrow(3, "00:00:00") - 1)
	//t.Log(TEndTheDayRemainSec() + 3600)
	//t.Log(TSecUntilNowToAddY(0, 0, 1, 1, 0, 0))

	//t.Log(TUnixSNowToAddY(0, 0, 0, 0, 0, 0))
	//t.Log(TSecUntilNowToAddY(0, 0, 0, 0, 45, 0))
	t.Log(TUnixSToDateTime(TUnixSNowToAdd(-1, 0, 0)))

}

func TestTFormatSec(t *testing.T) {
	t.Log(TFormatSec(3605, true))
}
