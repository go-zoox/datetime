package datetime

import (
	"testing"
	"time"
)

func TestLocation(t *testing.T) {
	d := FromTime(time.Now())
	location := d.Location()
	if location.String() == "" {
		t.Error("location is empty")
	}
}

func TestDate(t *testing.T) {
	dt := FromTime(time.Now())
	year, month, day := dt.Date()
	if year != dt.Year() || month != dt.Month() || day != dt.Day() {
		t.Error("date is not equal")
	}
}

func TestString(t *testing.T) {
	dt := FromTime(time.Now())
	if dt.String() != dt.Format() {
		t.Errorf("string is not equal to format: %s != %s", dt.String(), dt.Format())
	}
}
