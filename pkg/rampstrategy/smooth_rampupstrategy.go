package rampstrategy

import "time"

const MaxRamps = 10

// smoothRampUpStrategy starts sets of users at fixed intervals
type smoothRampUpStrategy struct {
	rampPeriodPct float64
}


func (s *smoothRampUpStrategy) GetStartProfile(testLength time.Duration, maxUsers int) []StartProfile {
	rampPeriodSec := int64(testLength.Seconds() * s.rampPeriodPct)

	rampIntervals := rampPeriodSec / MaxRamps

	usersPerRamp := maxUsers / MaxRamps
	lastRamp := maxUsers % MaxRamps

	strategies := make([]StartProfile, 0)
	wait := int64(0)
	for i := 0; i < MaxRamps; i++ {
		strategies = append(strategies, StartProfile{Delay: time.Duration(wait)*time.Second, Users: usersPerRamp})
		wait = wait + rampIntervals
	}
	if lastRamp > 0 {
		strategies = append(strategies, StartProfile{Delay: time.Duration(wait)*time.Second, Users: lastRamp})
	}

	return strategies
}

func NewSmoothRampUpStrategy(rampPeriod float64) RampStrategy {
	return &smoothRampUpStrategy{rampPeriodPct: rampPeriod}
}