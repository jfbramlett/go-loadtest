package rampstrategy

import "time"

const MaxRamps = 10

type smoothRampUpStrategy struct {
	rampPeriodPct float64
}


func (s *smoothRampUpStrategy) GetStartDelay(testLength time.Duration, maxUsers int) []StartDelay {
	rampPeriodSec := int64(testLength.Seconds() * s.rampPeriodPct)

	rampIntervals := rampPeriodSec / MaxRamps

	usersPerRamp := maxUsers / MaxRamps
	lastRamp := maxUsers % MaxRamps

	strategies := make([]StartDelay, 0)
	wait := int64(0)
	for i := 0; i < MaxRamps; i++ {
		strategies = append(strategies, StartDelay{InitialDelay: time.Duration(wait)*time.Second, UsersToStart: usersPerRamp})
		wait = wait + rampIntervals
	}
	if lastRamp > 0 {
		strategies = append(strategies, StartDelay{InitialDelay: time.Duration(wait)*time.Second, UsersToStart: lastRamp})
	}

	return strategies
}

func NewSmoothRampUpStrategy(rampPeriod float64) RampStrategy {
	return &smoothRampUpStrategy{rampPeriodPct: rampPeriod}
}