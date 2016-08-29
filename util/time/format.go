package time

import (
	"time"
)

const (
	DateTimeLayout = "2006-01-02 15:04:05"
)

// DateTimeFormat ...
func DateTimeFormat(t time.Time) string {
	return t.Format(DateTimeLayout)
}

// ParseDateTimeFormat returns
func ParseDateTimeFormat(v string) (time.Time, error) {
	return time.Parse(DateTimeLayout, v)
}

// MustParseDateTimeFormat returns
func MustParseDateTimeFormat(v string) time.Time {
	t, _ := ParseDateTimeFormat(v)
	return t
}
