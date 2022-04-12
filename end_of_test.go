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

func TestEndOfHour(t *testing.T) {
	dt, _ := New("YYYY-MM-DD HH:mm:ss", "2022-04-12 10:39:12")
	dt2 := dt.EndOfHour()

	// fmt.Println(dt2.Format("YYYY-MM-DD HH:mm:ss"))
	if dt2.Format("YYYY-MM-DD HH:mm:ss") != "2022-04-12 10:59:59" {
		t.Errorf("expected(2022-04-12 10:59:59), but got(%s)", dt2.Format("YYYY-MM-DD HH:mm:ss"))
	}
}
