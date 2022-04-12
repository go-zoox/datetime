package datetime

import "time"

// EndOfDay returns the end of the day.
func (dt *DateTime) EndOfDay() *DateTime {
	return dt.StartOfDay().Add(24*time.Hour - time.Nanosecond)
}

// EndOfWeek returns the end of the week.
func (dt *DateTime) EndOfWeek() *DateTime {
	return dt.StartOfWeek().Add(7*24*time.Hour - time.Nanosecond)
}

// EndOfMonth returns the end of the month.
func (dt *DateTime) EndOfMonth() *DateTime {
	return dt.StartOfMonth().AddDate(0, 1, 0).Add(-1 * time.Nanosecond)
}

// EndOfYear returns the end of the year.
func (dt *DateTime) EndOfYear() *DateTime {
	return dt.StartOfYear().AddDate(1, 0, 0).Add(-1 * time.Nanosecond)
}

// EndOfMinute returns the end of the hour.
func (dt *DateTime) EndOfMinute() *DateTime {
	return dt.StartOfMinute().Add(time.Minute - time.Nanosecond)
}

// EndOfHour returns the end of the hour.
func (dt *DateTime) EndOfHour() *DateTime {
	return dt.StartOfHour().Add(time.Hour - time.Nanosecond)
}

// EndOfQuarter returns the end of the quarter.
func (dt *DateTime) EndOfQuarter() *DateTime {
	return dt.StartOfQuarter().AddDate(0, 3, 0).Add(-1 * time.Nanosecond)
}
