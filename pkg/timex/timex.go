package timex

import (
	"errors"
	"time"
)

var (
	ErrInvalidTimeFormat = errors.New("invalid date format: %s, expected YYYY-MM-DD")
)

// ParseDateOnly parses a string into a time.Time object using the "YYYY-MM-DD" format.
// It returns an error if the input is not in the expected format.
func ParseDateOnly(dateStr string) (time.Time, error) {
	parsedTime, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return time.Time{}, ErrInvalidTimeFormat
	}
	return parsedTime, nil
}
