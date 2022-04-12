package datetime

import (
	"testing"
)

func TestSetTimeZone(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	// dt, _ := New(2022, 04, 12, 10, 39, 12)

	if dt.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 10:39:12" {
		t.Errorf("expected(2022-04-12 10:39:12), but got(%s)", dt.Format("YYYY-MM-DD HH:mm:ss"))
	}

	dt.SetTimeZone("Asia/Shanghai")
	if dt.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 10:39:12" {
		t.Errorf("expected(2022-04-12 10:39:12), but got(%s)", dt.Format("YYYY-MM-DD HH:mm:ss"))
	}

	dt.SetTimeZone("Asia/Tokyo")
	if dt.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 11:39:12" {
		t.Errorf("expected(2022-04-12 11:39:12), but got(%s)", dt.Format("YYYY-MM-DD HH:mm:ss"))
	}

	dt.SetTimeZone("Europe/Moscow")
	if dt.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 05:39:12" {
		t.Errorf("expected(2022-04-12 05:39:12), but got(%s)", dt.Format("YYYY-MM-DD HH:mm:ss"))
	}

	dt.SetTimeZone("America/New_York")
	if dt.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-11 22:39:12" {
		t.Errorf("expected(2022-04-11 22:39:12), but got(%s)", dt.Format("YYYY-MM-DD HH:mm:ss"))
	}

	dt.SetTimeZone("America/Los_Angeles")
	if dt.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-11 19:39:12" {
		t.Errorf("expected(2022-04-11 19:39:12), but got(%s)", dt.Format("YYYY-MM-DD HH:mm:ss"))
	}

	dt.SetTimeZone("Europe/London")
	if dt.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 03:39:12" {
		t.Errorf("expected(2022-04-12 03:39:12), but got(%s)", dt.Format("YYYY-MM-DD HH:mm:ss"))
	}

	dt.SetTimeZone("Europe/Paris")
	if dt.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 04:39:12" {
		t.Errorf("expected(2022-04-12 04:39:12), but got(%s)", dt.Format("YYYY-MM-DD HH:mm:ss"))
	}

	dt.SetTimeZone("Europe/Berlin")
	if dt.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 04:39:12" {
		t.Errorf("expected(2022-04-12 04:39:12), but got(%s)", dt.Format("YYYY-MM-DD HH:mm:ss"))
	}
}
