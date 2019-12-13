package rampstrategy

import "time"

type RampStrategyType int

const (
	Noop RampStrategyType = 0
	Smooth RampStrategyType = 1
	Random RampStrategyType = 2

	DefaultRampPeriod = .10
)


type StartProfile struct {
	Delay time.Duration
	Users int
}

type RampStrategy interface {
	GetStartProfile(testLength time.Duration, rampToUsers int) []StartProfile
}


func NewRampStrategy(rampType RampStrategyType) RampStrategy {
	switch rampType {
	case Noop:
		return NewNoRampUpStrategy()
	case Smooth:
		return NewSmoothRampUpStrategy(DefaultRampPeriod)
	case Random:
		return NewRandomRampUpStrategy(DefaultRampPeriod)
	}

	return nil
}