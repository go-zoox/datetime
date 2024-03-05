package datetime

// Timestamp return the miliseconds since the Unix epoch.
func (dt *DateTime) Timestamp() int64 {
	return dt.UnixMilli()
}

// Unix returns the number of seconds since the Unix epoch.
func (dt *DateTime) Unix() int64 {
	return dt.Time().Unix()
}

// UnixNano returns the number of nanoseconds since the Unix epoch.
func (dt *DateTime) UnixNano() int64 {
	return dt.Time().UnixNano()
}

// UnixMilli returns the number of milliseconds since the Unix epoch.
func (dt *DateTime) UnixMilli() int64 {
	return dt.UnixNano() / 1e6
}
