package datetime

// StartOfMinute returns the start of the hour.
func (dt *DateTime) StartOfMinute() *DateTime {
	dt2, _ := FromDate(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), 0, dt.t.Location().String())
	return dt2
}

// StartOfHour returns the start of the hour.
func (dt *DateTime) StartOfHour() *DateTime {
	dt2, _ := FromDate(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), 0, 0, dt.t.Location().String())
	return dt2
}

// StartOfDay returns the start of the day.
func (dt *DateTime) StartOfDay() *DateTime {
	dt2, _ := FromDate(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, dt.t.Location().String())
	return dt2
}

// StartOfWeek returns the start of the week.
func (dt *DateTime) StartOfWeek() *DateTime {
	weekday := dt.WeekDay()
	if weekday == 0 {
		weekday = 7
	}
	return dt.AddDate(0, 0, 1-weekday).StartOfDay()
}

// StartOfMonth returns the start of the month.
func (dt *DateTime) StartOfMonth() *DateTime {
	dt2, _ := FromDate(dt.Year(), dt.Month(), 1, 0, 0, 0, dt.t.Location().String())
	return dt2
}

// StartOfYear returns the start of the year.
func (dt *DateTime) StartOfYear() *DateTime {
	dt2, _ := FromDate(dt.Year(), 1, 1, 0, 0, 0, dt.t.Location().String())
	return dt2
}

// StartOfQuarter returns the end of the hour.
func (dt *DateTime) StartOfQuarter() *DateTime {
	month := dt.StartOfMonth()
	offset := (month.Month() - 1) % 3
	return month.AddDate(0, -offset, 0)
}
