package datetime

import "testing"

func TestNow(t *testing.T) {
	d := Now()
	if d.Timestamp() == 0 {
		t.Error("timestamp is 0")
	}
	if d.Location().String() == "" {
		t.Error("location is empty")
	}
}
