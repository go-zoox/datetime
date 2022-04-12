package datetime

import "time"

// SetTimeZone sets the time zone of the datetime.
func (dt *DateTime) SetTimeZone(timezone string) error {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	dt.t = dt.t.In(loc)
	return nil
}
