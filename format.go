package datetime

import (
	"fmt"
	"strings"
)

// Format return the format of time with pattern/layout
func (dt *DateTime) Format(pattern ...string) string {
	patternX := DefaultFormatPattern
	if len(pattern) > 0 {
		patternX = pattern[0]
	}

	year := dt.Year()
	month := dt.Month()
	day := dt.Day()
	hour := dt.Hour()
	minute := dt.Minute()
	second := dt.Second()
	millisecond := dt.Millisecond()
	weekDay := dt.WeekDay()

	_, offset := dt.t.Zone()
	zoneOffset := offset / 3600

	r := strings.NewReplacer(
		"YYYY", fmt.Sprintf("%d", year),
		"YY", fmt.Sprintf("%02d", year),
		"MM", fmt.Sprintf("%02d", month),
		"M", fmt.Sprintf("%d", month),
		"DD", fmt.Sprintf("%02d", day),
		"D", fmt.Sprintf("%d", day),
		"HH", fmt.Sprintf("%02d", hour),
		"H", fmt.Sprintf("%d", hour),
		"hh", fmt.Sprintf("%02d", hour%12),
		"h", fmt.Sprintf("%d", hour%12),
		"mm", fmt.Sprintf("%02d", minute),
		"m", fmt.Sprintf("%d", minute),
		"ss", fmt.Sprintf("%02d", second),
		"s", fmt.Sprintf("%d", second),
		"SSS", fmt.Sprintf("%03d", millisecond),
		//
		"dddd", LocalesWeekDays[weekDay],
		"ddd", LocalesWeekDaysShort[weekDay],
		"dd", LocalesWeekDaysMin[weekDay],
		"d", fmt.Sprintf("%d", weekDay),
		//
		"a", getLowerA(hour),
		"A", getUpperA(hour),
		//
		"ZZ", getZoneOffset(zoneOffset, ""),
		"Z", getZoneOffset(zoneOffset, ":"),
	)

	return r.Replace(patternX)
}

func getLowerA(hour int) string {
	if hour < 12 {
		return "am"
	}
	return "pm"
}

func getUpperA(hour int) string {
	if hour < 12 {
		return "AM"
	}
	return "PM"
}

func getZoneOffset(zoneOffset int, seperator string) string {
	if zoneOffset > 0 {
		return fmt.Sprintf("+%02d%s00", zoneOffset, seperator)
	}

	return fmt.Sprintf("%02d%s00", zoneOffset, seperator)
}
