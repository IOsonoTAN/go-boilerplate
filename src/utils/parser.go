package utils

import (
	"time"
)

// ConvertTimezoneDatetime a function to convert timezone to location in config
func ConvertTimezoneDatetime(dt time.Time, timezone string) time.Time {
	var convToTimezone string
	if len(timezone) > 0 {
		convToTimezone = timezone
	} else {
		convToTimezone = "Asia/Bangkok"
	}
	location, _ := time.LoadLocation(convToTz)
	datetime := time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), dt.Nanosecond(), time.UTC).In(location)

	return datetime
}
