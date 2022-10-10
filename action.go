package datetime

import "time"

// Add adds the duration to the datetime.
func (dt *DateTime) Add(d time.Duration) *DateTime {
	t := dt.t.Add(d)
	return &DateTime{
		t: t,
	}
}

// Sub subtracts the duration from the datetime.
func (dt *DateTime) Sub(t *DateTime) time.Duration {
	return dt.t.Sub(t.t)
}

// Before returns true if the datetime is before the other datetime.
func (dt *DateTime) Before(t *DateTime) bool {
	return dt.t.Before(t.t)
}

// After returns true if the datetime is after the other datetime.
func (dt *DateTime) After(t *DateTime) bool {
	return dt.t.After(t.t)
}

// AddDate adds the specified years, months, and days to the datetime.
func (dt *DateTime) AddDate(years int, months int, days int) *DateTime {
	t := dt.t.AddDate(years, months, days)
	return &DateTime{
		t: t,
	}
}

// Equal reports whether t and dt represent the same time instant.
func (dt *DateTime) Equal(t *DateTime) bool {
	return dt.t.Equal(t.t)
}
