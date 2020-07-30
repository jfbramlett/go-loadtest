package rampstrategy

import (
	"context"
	"time"
)

type RampStrategyType int

const (
	Noop RampStrategyType = 0
	Smooth RampStrategyType = 1
	Random RampStrategyType = 2

	DefaultRampPeriod = .10
)


func GetRampStrategyType(v int) RampStrategyType {
	switch v {
	case 0:
		return Noop
	case 1:
		return Smooth
	case 2:
		return Random
	}

	return Noop
}

type StartProfile struct {
	Delay time.Duration
	Users int
}

type RampStrategy interface {
	GetStartProfile(ctx context.Context, testLength time.Duration, rampToUsers int) []StartProfile
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