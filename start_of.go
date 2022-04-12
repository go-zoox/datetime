package datetime

import "time"

// StartOfDay returns the start of the day.
func (dt *DateTime) StartOfDay() *DateTime {
	return FromTime(dt.t.Truncate(24 * time.Hour))
}

// StartOfWeek returns the start of the week.
func (dt *DateTime) StartOfWeek() *DateTime {
	return FromTime(dt.t.Truncate(7 * 24 * time.Hour))
}

// StartOfMonth returns the start of the month.
func (dt *DateTime) StartOfMonth() *DateTime {
	return FromDate(dt.Year(), dt.Month(), 1, 0, 0, 0)
}

// StartOfYear returns the start of the year.
func (dt *DateTime) StartOfYear() *DateTime {
	return FromDate(dt.Year(), 1, 1, 0, 0, 0)
}

// StartOfHour returns the start of the hour.
func (dt *DateTime) StartOfHour() *DateTime {
	return FromTime(dt.t.Truncate(time.Hour))
}
