package datetime

import "time"

// Now returns the time.Time of DateTime.
func Now() *DateTime {
	return FromTime(time.Now())
}
