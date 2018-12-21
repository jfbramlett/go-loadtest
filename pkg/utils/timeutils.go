package utils

import "time"

func TimeDiffMillis(start time.Time, end time.Time) int64 {
	return toMillis(end) - toMillis(start)
}

func toMillis(t time.Time) int64 {
	return t.UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}


