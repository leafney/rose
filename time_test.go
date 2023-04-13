package rose

import "testing"

func TestTimeNow10(t *testing.T) {
	t.Log(TNowS())
}

func TestTimeNow13(t *testing.T) {
	t.Log(TNowMs())
}

func TestParseGMTTime(t *testing.T) {
	tt, _ := TParseGMTTimeOfRFC1123("Mon, 20 Jul 2020 06:09:21 GMT")
	t.Log(tt.Unix())
	t.Log(TUnixSFormat(tt.Unix(), TFLongYMDHS))
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
	t.Log(TSecUntilNowToAddY(0, 0, 0, 0, 45, 0))
	//t.Log(TUnixSToDateTime(TNowAddUnixSDHM(-1, 0, 0)))

	//t.Log(TMonthEndStr())

	t.Log(TNowAddIntervalSecDHM(0, 0, 45))

}

func TestTFormatSec(t *testing.T) {
	t.Log(TFormatSec(3505, true))
	t.Log(TSecFormatMS(3505))
	t.Log(TSecFormatHMS(3505))
}
