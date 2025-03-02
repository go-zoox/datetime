package datetime

import (
	"testing"
)

func TestStartOfDay(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.StartOfDay()

	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 00:00:00" {
		t.Errorf("expected(2022-04-12 00:00:00), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestStartOfWeek(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.StartOfWeek()

	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-11 00:00:00" {
		t.Errorf("expected(2022-04-11 00:00:00), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-24"); d.StartOfWeek().Format("YYYY-MM-DD") != "2025-02-24" {
		t.Errorf("expected(2025-02-24), but got(%s)", d.StartOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-25"); d.StartOfWeek().Format("YYYY-MM-DD") != "2025-02-24" {
		t.Errorf("expected(2025-02-24), but got(%s)", d.StartOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-26"); d.StartOfWeek().Format("YYYY-MM-DD") != "2025-02-24" {
		t.Errorf("expected(2025-02-24), but got(%s)", d.StartOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-27"); d.StartOfWeek().Format("YYYY-MM-DD") != "2025-02-24" {
		t.Errorf("expected(2025-02-24), but got(%s)", d.StartOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-28"); d.StartOfWeek().Format("YYYY-MM-DD") != "2025-02-24" {
		t.Errorf("expected(2025-02-24), but got(%s)", d.StartOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-03-01"); d.StartOfWeek().Format("YYYY-MM-DD") != "2025-02-24" {
		t.Errorf("expected(2025-02-24), but got(%s)", d.StartOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-03-02"); d.StartOfWeek().Format("YYYY-MM-DD") != "2025-02-24" {
		t.Errorf("expected(2025-02-24), but got(%s)", d.StartOfWeek().Format("YYYY-MM-DD"))
	}
}

func TestStartOfMonth(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.StartOfMonth()

	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-01 00:00:00" {
		t.Errorf("expected(2022-04-01 00:00:00), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestStartOfYear(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.StartOfYear()

	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-01-01 00:00:00" {
		t.Errorf("expected(2022-01-01 00:00:00), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestStartOfMinute(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.StartOfMinute()

	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 10:39:00" {
		t.Errorf("expected(2022-04-12 10:39:00), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestStartOfHour(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.StartOfHour()

	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 10:00:00" {
		t.Errorf("expected(2022-04-12 10:00:00), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestStartOfQuarter(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.StartOfQuarter()

	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-01 00:00:00" {
		t.Errorf("expected(2022-04-01 00:00:00), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}
