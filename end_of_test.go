package datetime

import (
	"testing"
)

func TestEndOfDay(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.EndOfDay()

	// fmt.Println(dt2.Format("YYYY-MM-DD HH:mm:ss"))
	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 23:59:59" {
		t.Errorf("expected(2022-04-12 23:59:59), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestEndOfWeek(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.EndOfWeek()

	// fmt.Println(dt2.Format("YYYY-MM-DD HH:mm:ss"))
	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-17 23:59:59" {
		t.Errorf("expected(2022-04-17 23:59:59), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-24"); d.EndOfWeek().Format("YYYY-MM-DD") != "2025-03-02" {
		t.Errorf("expected(2025-03-02), but got(%s)", d.EndOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-25"); d.EndOfWeek().Format("YYYY-MM-DD") != "2025-03-02" {
		t.Errorf("expected(2025-03-02), but got(%s)", d.EndOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-26"); d.EndOfWeek().Format("YYYY-MM-DD") != "2025-03-02" {
		t.Errorf("expected(2025-03-02), but got(%s)", d.EndOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-27"); d.EndOfWeek().Format("YYYY-MM-DD") != "2025-03-02" {
		t.Errorf("expected(2025-03-02), but got(%s)", d.EndOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-02-28"); d.EndOfWeek().Format("YYYY-MM-DD") != "2025-03-02" {
		t.Errorf("expected(2025-03-02), but got(%s)", d.EndOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-03-01"); d.EndOfWeek().Format("YYYY-MM-DD") != "2025-03-02" {
		t.Errorf("expected(2025-03-02), but got(%s)", d.EndOfWeek().Format("YYYY-MM-DD"))
	}

	if d, _ := New("YYYY-MM-DD", "2025-03-02"); d.EndOfWeek().Format("YYYY-MM-DD") != "2025-03-02" {
		t.Errorf("expected(2025-03-02), but got(%s)", d.EndOfWeek().Format("YYYY-MM-DD"))
	}
}

func TestEndOfMonth(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.EndOfMonth()

	// fmt.Println(dt2.Format("YYYY-MM-DD HH:mm:ss"))
	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-30 23:59:59" {
		t.Errorf("expected(2022-04-30 23:59:59), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestEndOfYear(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.EndOfYear()

	// fmt.Println(dt2.Format("YYYY-MM-DD HH:mm:ss"))
	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-12-31 23:59:59" {
		t.Errorf("expected(2022-12-31 23:59:59), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestEndOfMinute(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.EndOfMinute()

	// fmt.Println(dt2.Format("YYYY-MM-DD HH:mm:ss"))
	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 10:39:59" {
		t.Errorf("expected(2022-04-12 10:39:59), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestEndOfHour(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.EndOfHour()

	// fmt.Println(dt2.Format("YYYY-MM-DD HH:mm:ss"))
	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 10:59:59" {
		t.Errorf("expected(2022-04-12 10:59:59), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}

func TestEndOfQuarter(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.EndOfQuarter()

	// fmt.Println(dt2.Format("YYYY-MM-DD HH:mm:ss"))
	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-06-30 23:59:59" {
		t.Errorf("expected(2022-06-30 23:59:59), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}
