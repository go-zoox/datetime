package datetime

import (
	"errors"
	"time"
)

// New returns a DateTime
// Support:
//   1. time.Time
//		 example: New(time.Now())
//
//	 2. YYYY/MM/DD HH:mm:ss, datetime string
//     example: New("YYYY/MM/DD HH:mm:ss", "2022/04/12 10:39:12")
//
//   3. year, month, day, hour, minute, second
//     example: New(2022, 4, 12, 10, 39, 12)
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
		if len(args) != 2 {
			return nil, errors.New("invalid arguments, should be New(pattern, datetime)")
		}

		return FromPattern(args[0].(string), args[1].(string))
	case int:
		if len(args) != 6 {
			return nil, errors.New("invalid arguments, should be New(year, month, day, hour, minute, second)")
		}

		return FromDate(args[0].(int), args[1].(int), args[2].(int), args[3].(int), args[4].(int), args[5].(int)), nil
	default:
		return nil, errors.New("unsupported arguments: " + args[0].(string))
	}
}

// FromTime returns a DateTime from time.Time.
// Example:
//	t := time.Now()
//	dt := FromTime(t)
func FromTime(t time.Time) *DateTime {
	return &DateTime{
		t: t,
	}
}

// FromPattern returns a DateTime from a pattern and a datetime string.
// Example:
//	pattern := "YYYY/MM/DD HH:mm:ss"
//	datetime := "2022/04/12 10:39:12"
//	dt, err := FromPattern(pattern, datetime)
func FromPattern(pattern string, datetime string) (*DateTime, error) {
	gobirth := time.Date(2006, time.January, 2, 15, 04, 05, 0, time.Local)
	dt := &DateTime{t: gobirth}
	f := dt.Format(pattern)
	t, err := time.ParseInLocation(f, datetime, time.Local)
	if err != nil {
		return nil, err
	}

	return FromTime(t), nil
}

// FromDate returns a DateTime from a Date.
// Example:
//  FromDate(2022, 4, 12, 10, 39, 12)
func FromDate(year, month, day, hour, minute, second int) *DateTime {
	return FromTime(time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local))
}
