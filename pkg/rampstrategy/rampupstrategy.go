package rampstrategy

import "time"

type StartDelay struct {
	InitialDelay	time.Duration
	UsersToStart	int
}

type RampStrategy interface {
	GetStartDelay(testLength time.Duration, rampToUsers int) []StartDelay
}