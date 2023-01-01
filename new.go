package datetime

import (
	"errors"
	"time"
)

// New returns a DateTime
// Support:
//
//  1. time.Time
//     example: New(time.Now())
//
//  2. YYYY/MM/DD HH:mm:ss, datetime string
//     example: New("YYYY/MM/DD HH:mm:ss", "2022/04/12 10:39:12")
//     New("YYYY/MM/DD HH:mm:ss", "2022/04/12 10:39:12", "Asia/Shanghai")
//
//  3. year, month, day, hour, minute, second
//     example: New(2022, 4, 12, 10, 39, 12)
//     New(2022, 4, 12, 10, 39, 12, "Asia/Shanghai")
func New(args ...interface{}) (*DateTime, error) {
	if len(args) == 0 {
		return FromTime(time.Now()), nil
	}

	switch args[0].(type) {
	case time.Time:
		if len(args) != 1 {
			return nil, errors.New("invalid arguments, should be New(time.Time)")
		}

		return FromTime(args[0].(time.Time)), nil
	case string:
		if len(args) == 2 {
			return FromPattern(args[0].(string), args[1].(string))
		} else if len(args) == 3 {
			return FromPattern(args[0].(string), args[1].(string), args[2].(string))
		}

		return nil, errors.New("invalid arguments, should be New(pattern, datetime[, timezone])")
	case int:
		if len(args) == 6 {
			return FromDate(args[0].(int), args[1].(int), args[2].(int), args[3].(int), args[4].(int), args[5].(int))
		} else if len(args) == 7 {
			return FromDate(args[0].(int), args[1].(int), args[2].(int), args[3].(int), args[4].(int), args[5].(int), args[6].(string))
		}

		return nil, errors.New("invalid arguments, should be New(year, month, day, hour, minute, second[, timezone])")
	default:
		return nil, errors.New("unsupported arguments: " + args[0].(string))
	}
}

// FromTime returns a DateTime from time.Time.
// Example:
//
//	t := time.Now()
//	dt := FromTime(t)
func FromTime(t time.Time) *DateTime {
	return &DateTime{
		t: t,
	}
}

// FromPattern returns a DateTime from a pattern(layout) and a datetime string.
// Example:
//
//		 pattern := "YYYY/MM/DD HH:mm:ss"
//		 datetime := "2022/04/12 10:39:12"
//		 dt, err := FromPattern(pattern, datetime)
//		 dt, err := FromPattern(pattern, datetime, "Asia/Shanghai")
//
//	Available pattern(layout): (file: /usr/local/go/src/time/format.go)
//	   time.Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
//	   time.ANSIC       = "Mon Jan _2 15:04:05 2006"
//	   time.UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
//	   time.RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
//	   time.RFC822      = "02 Jan 06 15:04 MST"
//	   time.RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
//	   time.RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
//	   time.RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
//	   time.RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
//	   time.RFC3339     = "2006-01-02T15:04:05Z07:00"
//	   time.RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
//	   time.Kitchen     = "3:04PM"
//	   // Handy time stamps.
//	   time.Stamp      = "Jan _2 15:04:05"
//	   time.StampMilli = "Jan _2 15:04:05.000"
//	   time.StampMicro = "Jan _2 15:04:05.000000"
//	   time.StampNano  = "Jan _2 15:04:05.000000000"
func FromPattern(pattern string, datetime string, timezone ...string) (*DateTime, error) {
	tz := time.Local
	if len(timezone) > 0 {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return nil, err
		}

		tz = loc
	}

	gobirth := time.Date(2006, time.January, 2, 15, 04, 05, 0, time.Local)
	dt := &DateTime{t: gobirth}
	f := dt.Format(pattern)
	t, err := time.ParseInLocation(f, datetime, tz)
	if err != nil {
		return nil, err
	}

	return FromTime(t), nil
}

// FromDate returns a DateTime from a Date.
// Example:
//
//	FromDate(2022, 4, 12, 10, 39, 12)
//	FromDate(2022, 4, 12, 10, 39, 12, "Asia/Shanghai")
func FromDate(year, month, day, hour, minute, second int, timezone ...string) (*DateTime, error) {
	tz := time.Local
	if len(timezone) > 0 {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return nil, err
		}

		tz = loc
	}

	return FromTime(time.Date(year, time.Month(month), day, hour, minute, second, 0, tz)), nil
}
