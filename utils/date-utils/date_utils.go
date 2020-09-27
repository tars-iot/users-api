package dateutils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNowUTC is function to return current time in Time format
func GetNowUTC() time.Time {
	return time.Now().UTC()
}

//GetNowString is fucntion to return current date time in string format
func GetNowString() string {
	return GetNowUTC().Format(apiDateLayout)
}
