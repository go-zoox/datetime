package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	d := FromTime(time.Now())
	timestamp := d.Timestamp()
	if len(fmt.Sprintf("%d", timestamp)) != 13 {
		t.Error("timestamp length is not 13")
	}
}

func TestLocation(t *testing.T) {
	d := FromTime(time.Now())
	location := d.Location()
	if location.String() == "" {
		t.Error("location is empty")
	}
}
