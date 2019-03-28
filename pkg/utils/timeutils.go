package utils

import (
	"time"
)

func TimeDiffMillis(start time.Time, end time.Time) int64 {
	return ToMillis(end) - ToMillis(start)
}

func ToMillis(t time.Time) int64 {
	return t.UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}

func DurationMillis(dur time.Duration ) int64 {
	return int64(dur/time.Millisecond)
}
