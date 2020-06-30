package utils

import (
	"time"
)

// GetSecondsToMidNight a function to calculate a seconds from now to midnight
func GetSecondsToMidNight() int {
	location, _ := time.LoadLocation("Asia/Bangkok")
	timeNow := time.Now().In(location)
	timeMidnight := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 23, 59, 59, 999, location)
	diffSeconds := timeMidnight.Sub(timeNow).Seconds()
	seconds := int(diffSeconds)

	if seconds > 0 {
		return seconds
	}
	return 0
}
