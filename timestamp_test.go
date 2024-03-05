package datetime

import (
	"testing"
)

func TestTimestamp(t *testing.T) {
	d, _ := FromPattern("YYYY/MM/DD HH:mm:ss", "2024/03/05 20:11:32", "Asia/Shanghai")
	var expected int64 = 1709640692000
	recieved := d.Timestamp()
	if recieved != expected {
		t.Fatalf("expected(%d), but got(%d)", expected, recieved)
	}
}

func TestUnix(t *testing.T) {
	d, _ := FromPattern("YYYY/MM/DD HH:mm:ss", "2024/03/05 20:11:32", "Asia/Shanghai")
	var expected int64 = 1709640692
	recieved := d.Unix()
	if recieved != expected {
		t.Fatalf("expected(%d), but got(%d)", expected, recieved)
	}
}

func TestUnixNano(t *testing.T) {
	d, _ := FromPattern("YYYY/MM/DD HH:mm:ss", "2024/03/05 20:11:32", "Asia/Shanghai")
	var expected int64 = 1709640692000000000
	recieved := d.UnixNano()
	if recieved != expected {
		t.Fatalf("expected(%d), but got(%d)", expected, recieved)
	}
}

func TestUnixMilli(t *testing.T) {
	d, _ := FromPattern("YYYY/MM/DD HH:mm:ss", "2024/03/05 20:11:32", "Asia/Shanghai")
	var expected int64 = 1709640692000
	recieved := d.UnixMilli()
	if recieved != expected {
		t.Fatalf("expected(%d), but got(%d)", expected, recieved)
	}
}
