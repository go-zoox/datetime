package datetime

import "time"

// StartOfDay returns the start of the day.
func (dt *DateTime) StartOfDay() *DateTime {
	dt2, _ := FromDate(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0)
	return dt2
}

// StartOfWeek returns the start of the week.
func (dt *DateTime) StartOfWeek() *DateTime {
	return dt.AddDate(0, 0, 1-dt.WeekDay()).StartOfDay()
}

// StartOfMonth returns the start of the month.
func (dt *DateTime) StartOfMonth() *DateTime {
	dt2, _ := FromDate(dt.Year(), dt.Month(), 1, 0, 0, 0)
	return dt2
}

// StartOfYear returns the start of the year.
func (dt *DateTime) StartOfYear() *DateTime {
	dt2, _ := FromDate(dt.Year(), 1, 1, 0, 0, 0)
	return dt2
}

// StartOfHour returns the start of the hour.
func (dt *DateTime) StartOfHour() *DateTime {
	return FromTime(dt.t.Truncate(time.Hour))
}
