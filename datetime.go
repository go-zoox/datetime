package datetime

import (
	"time"
)

// DateTime is a struct that represents a datetime.
type DateTime struct {
	t time.Time
	//
	isTimeZoneSet bool
}

// Year returns the year.
func (dt *DateTime) Year() int {
	return dt.t.Year()
}

// Month returns the month.
func (dt *DateTime) Month() int {
	return int(dt.t.Month())
}

// Day returns the day of the month.
func (dt *DateTime) Day() int {
	return dt.t.Day()
}

// Hour returns the hour.
func (dt *DateTime) Hour() int {
	return dt.t.Hour()
}

// Minute returns the minutes.
func (dt *DateTime) Minute() int {
	return dt.t.Minute()
}

// Second returns the seconds.
func (dt *DateTime) Second() int {
	return dt.t.Second()
}

// Millisecond returns the milliseconds.
func (dt *DateTime) Millisecond() int {
	return dt.t.Nanosecond() / 1e6
}

// WeekDay returns the day of week
//
//	0 - Sunday
//	1 - Monday
//	2 - Tuesday
//	3 - Wednesday
//	4 - Thursday
//	5 - Friday
//	6 - Sataurday
func (dt *DateTime) WeekDay() int {
	return int(dt.t.Weekday())
}

// Location returns the location.
func (dt *DateTime) Location() *time.Location {
	return dt.t.Location()
}

// Time returns the time.Time.
func (dt *DateTime) Time() time.Time {
	return dt.t
}

// Date returns the date.
func (dt *DateTime) Date() (year, month, day int) {
	return dt.Year(), dt.Month(), dt.Day()
}

// String returns the string representation of the datetime.
func (dt *DateTime) String() string {
	return dt.Format()
}
