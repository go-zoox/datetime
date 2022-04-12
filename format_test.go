package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatDefaultPattern(t *testing.T) {
	ti := time.Now()
	d := FromTime(ti)
	// fmt.Println(d.Format(), ti.Format("2006-01-02 15:04:05"))
	if d.Format() != ti.Format("2006-01-02 15:04:05") {
		t.Errorf("expected(%s), but got(%s)", ti.Format("2006-01-02 15:04:05"), d.Format())
	}
}

func TestFormatCustomPattern(t *testing.T) {
	ti := time.Now()
	d := FromTime(ti)
	pattern := "YYYY-MM-DD HH:mm:ss.SSS"
	// fmt.Println(d.Format(pattern), ti.Format("2006-01-02 15:04:05.000"))
	if d.Format(pattern) != ti.Format("2006-01-02 15:04:05.000") {
		t.Errorf("expected(%s), but got(%s)", ti.Format("2006-01-02 15:04:05"), d.Format())
	}
}

func TestFormatWeekDay(t *testing.T) {
	ti := time.Now()
	d := FromTime(ti)

	if d.WeekDay() != int(ti.Weekday()) {
		t.Errorf("expected(%d), but got(%d)", ti.Weekday(), d.WeekDay())
	}

	if d.Format("dddd") != ti.Weekday().String() {
		t.Errorf("expected(%s), but got(%s)", ti.Weekday().String(), d.Format("dddd"))
	}

	// fmt.Println(d.Format("d"), ti.Weekday().String())
	// fmt.Println(d.Format("dd"))
	// fmt.Println(d.Format("ddd"))
	// fmt.Println(d.Format("dddd"))
}

func TestFormatTimeZone(t *testing.T) {
	ti := time.Now()
	d := FromTime(ti)

	_, offset := ti.Zone()

	if d.Format("Z") != fmt.Sprintf("+%02d:00", offset/3600) {
		t.Errorf("expected(%s), but got(%s)", fmt.Sprintf("+%02d:00", offset/3600), d.Format("Z"))
	}

	if d.Format("ZZ") != fmt.Sprintf("+%02d00", offset/3600) {
		t.Errorf("expected(%s), but got(%s)", fmt.Sprintf("+%02d00", offset/3600), d.Format("ZZ"))
	}
}
