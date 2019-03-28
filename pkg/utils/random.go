package utils

import (
	"math/rand"
	"time"
)


type DurationGenator func(min, max time.Duration) time.Duration


func RandomIntBetween(min int, max int) int {
	return rand.Intn(max - min) + min
}

func RandomInt64Between(min int64, max int64) int64 {
	return rand.Int63n(max - min) + min
}


func RandomDuration(min time.Duration, max time.Duration) time.Duration {
	return time.Duration(rand.Int63n(max.Nanoseconds() - min.Nanoseconds()) + min.Nanoseconds())
}


func RandomSecondDuration(min time.Duration, max time.Duration) time.Duration {
	randSecs := RandomIntBetween(int(min.Seconds()), int(max.Seconds()))
	return time.Duration(time.Duration(randSecs) * time.Second)
}


func RandomMilliSecondDuration(min time.Duration, max time.Duration) time.Duration {
	randMilliSecs := RandomInt64Between(DurationMillis(min), DurationMillis(max))
	return time.Duration(time.Duration(randMilliSecs) * time.Millisecond)
}