package timeutil

import "time"

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func TimeStringToTime(dateString string) (outTime time.Time, err error) {
	outTime, err = time.Parse(TimeFormat, dateString)
	return
}
