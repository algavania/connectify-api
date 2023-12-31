package pkg

import (
	"time"
)

func ParseTime(timeStr string) (time.Time, error) {
	layout := "2006-01-02T15:04:05+07:00"
	birthdayString := timeStr
	parsedBirthday, err := time.Parse(layout, birthdayString)
	return parsedBirthday, err
}
