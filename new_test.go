package datetime

import (
	"testing"
	"time"
)

func TestFromTime(t *testing.T) {
	ti := time.Now()
	dt := FromTime(ti)
	if dt.Time().String() != ti.String() {
		t.Errorf("expected(%s), but got(%s)", ti.String(), dt.Time().String())
	}
}

func TestFromPattern(t *testing.T) {
	pattern := "YYYY/MM/DD HH:mm:ss"
	datetime := "2016/03/06 22:04:05"
	dt, err := FromPattern(pattern, datetime)
	if err != nil {
		t.Error(err)
	}

	if dt.Format(pattern) != datetime {
		t.Errorf("expected(%s), but got(%s)", datetime, dt.Format(pattern))
	}

	if dt.Format(pattern) != dt.Time().Format("2006/01/02 15:04:05") {
		t.Errorf("expected(%s), but got(%s)", dt.Time().Format("2006/01/02 01:02:03"), dt.Format(pattern))
	}
}

func TestNew(t *testing.T) {
	ti := time.Now()
	dt, err := New(ti)
	if err != nil {
		t.Error(err)
	}
	if dt.Time().String() != ti.String() {
		t.Errorf("expected(%s), but got(%s)", ti.String(), dt.Time().String())
	}

	pattern := "YYYY/MM/DD HH:mm:ss"
	datetime := "2016/03/06 22:04:05"
	dt2, err2 := New(pattern, datetime)
	if err2 != nil {
		t.Error(err)
	}

	if dt2.Format(pattern) != datetime {
		t.Errorf("expected(%s), but got(%s)", datetime, dt2.Format(pattern))
	}

	if dt2.Format(pattern) != dt2.Time().Format("2006/01/02 15:04:05") {
		t.Errorf("expected(%s), but got(%s)", dt2.Time().Format("2006/01/02 01:02:03"), dt2.Format(pattern))
	}

	sourceTimezone := "Europe/London"
	targetTimezone := "Asia/Shanghai"
	dt3, err3 := New(pattern, datetime, sourceTimezone)
	if err3 != nil {
		t.Error(err)
	}
	if dt3.Format(pattern) != datetime {
		t.Errorf("expected(%s), but got(%s)", datetime, dt3.Format(pattern))
	}
	if err := dt3.SetTimeZone(targetTimezone); err != nil {
		t.Fatal(err)
	}
	if dt3.Format(pattern) != "2016/03/07 06:04:05" {
		t.Errorf("expected(%s), but got(%s)", dt3.Format(pattern), "2016/03/07 06:04:05")
	}
}
