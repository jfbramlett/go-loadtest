package rampstrategy

import "time"

type StartProfile struct {
	Delay time.Duration
	Users int
}

type RampStrategy interface {
	GetStartProfile(testLength time.Duration, rampToUsers int) []StartProfile
}