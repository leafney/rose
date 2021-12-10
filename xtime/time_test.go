package xtime

import (
	"testing"
)

func TestTimeNow10(t *testing.T) {
	t.Log(NowS())
}

func TestTimeNow13(t *testing.T) {
	t.Log(NowMs())
}

func TestParseGMTTime(t *testing.T) {
	tt, _ := ParseGMTTimeOfRFC1123("Mon, 20 Jul 2020 06:09:21 GMT")
	t.Log(tt.Unix())
	t.Log(GetTimeByUnix(tt.Unix()))
}

func TestGetMonth(t *testing.T) {
	t.Log(GetMonth())
	t.Log(GetMonthLast())
	t.Log(GetMonthNext())
	t.Log(GetMonthByUnix(1595225361))
	t.Log(GetTimeMHByUnix(1595225361))
	t.Log(MonthStart())
	t.Log(MonthEnd())
	t.Log(TodayStart())
	t.Log(TodayEnd())
}
