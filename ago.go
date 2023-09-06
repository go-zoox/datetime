package datetime

import (
	"fmt"
	"time"
)

// AgoPattern is the pattern for Ago.
var AgoPattern = map[string]string{
	"now":    "just now",
	"second": "%d second ago",
	"minute": "%d minute ago",
	"hour":   "%d hour ago",
	"day":    "%d day ago",
	"week":   "%d week ago",
	"month":  "%d month ago",
	"year":   "%d year ago",
}

// AgoOptions is the options for Ago.
type AgoOptions struct {
	Diff *DateTime
}

// Ago return time ago from now.
func (dt *DateTime) Ago(opts ...*AgoOptions) string {
	var diff *DateTime
	if len(opts) > 0 && opts[0] != nil {
		diff = opts[0].Diff
	} else {
		diff = Now()
	}

	delta := diff.Sub(dt)
	if delta < time.Second {
		return AgoPattern["now"]
	}

	if delta < time.Minute {
		return fmt.Sprintf(AgoPattern["second"], delta/time.Second)
	}

	if delta < time.Hour {
		return fmt.Sprintf(AgoPattern["minute"], delta/time.Minute)
	}

	if delta < time.Hour*24 {
		return fmt.Sprintf(AgoPattern["hour"], delta/time.Hour)
	}

	if delta < time.Hour*24*7 {
		return fmt.Sprintf(AgoPattern["day"], delta/(time.Hour*24))
	}

	if delta < time.Hour*24*30 {
		return fmt.Sprintf(AgoPattern["week"], delta/(time.Hour*24*7))
	}

	if delta < time.Hour*24*30*12 {
		return fmt.Sprintf(AgoPattern["month"], delta/(time.Hour*24*30))
	}

	return fmt.Sprintf(AgoPattern["year"], delta/(time.Hour*24*30*12))
}
