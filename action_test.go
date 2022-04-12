package datetime

import (
	"testing"
	"time"
)

func TestActionAdd(t *testing.T) {
	dt, err := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	if err != nil {
		t.Error(err)
	}

	dt2 := dt.Add(1 * time.Hour)
	if dt2.Hour()-dt.Hour() != 1 {
		t.Errorf("expected(%d), but got(%d)", 1, dt2.Hour()-dt.Hour())
	}
}

func TestActionSub(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 12:39:12")

	duration := dt2.Sub(dt)
	if duration != 2*time.Hour {
		t.Errorf("expected(%d), but got(%d)", 2*time.Hour, duration)
	}
}

func TestActionBefore(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 12:39:12")

	if !dt.Before(dt2) {
		t.Errorf("expected(true), but got(false)")
	}

	if dt2.Before(dt) {
		t.Errorf("expected(false), but got(true)")
	}
}

func TestActionAfter(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 12:39:12")

	if !dt2.After(dt) {
		t.Errorf("expected(true), but got(false)")
	}

	if dt.After(dt2) {
		t.Errorf("expected(false), but got(true)")
	}
}

func TestActionEqual(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2, _ := New(2022, 04, 12, 10, 39, 12)
	dt3, _ := New(time.Date(2022, time.April, 12, 10, 39, 12, 0, time.Local))

	if !dt.Equal(dt2) {
		t.Errorf("expected(true), but got(false)")
	}

	if !dt.Equal(dt3) {
		t.Errorf("expected(true), but got(false)")
	}
}
