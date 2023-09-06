package datetime

// DefaultFormatPattern is the default format pattern.
var DefaultFormatPattern = "YYYY-MM-DD HH:mm:ss"

// TimeZoneForce will force to override datetime instance time zone
//
//		default is empty string, means no force time zone, use datetime self timezone, it is local timezone
//	 for example, set as Asia/Shanghai,
//			Format will use Asia/Shanghai override internal time zone
//		  unless it has call SetTimeZone before
var TimeZoneForce = ""

// LocalesWeekDays is the week days.
var LocalesWeekDays = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

// LocalesWeekDaysShort is the week days in short.
var LocalesWeekDaysShort = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

// LocalesWeekDaysMin is the week days in minium.
var LocalesWeekDaysMin = []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}

// Locales is the locales.
var Locales = I18nLocales{}

// Language is the language.
var Language = "en-US"
