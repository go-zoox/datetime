package datetime

import "time"

// SetTimeZone sets the time zone of the datetime.
func (dt *DateTime) SetTimeZone(timezone string) error {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	dt.t = dt.t.In(loc)
	dt.isTimeZoneSet = true
	return nil
}

// GetTimeZone gets the time zone of the datetime.
func (dt *DateTime) GetTimeZone() string {
	loc := dt.t.Location()
	return loc.String()
}

// SetTimeZone sets the global time zone of the datetime.
func SetTimeZone(timezone string) {
	TimeZoneForce = timezone
}
