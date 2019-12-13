package rampstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

const RampUsersPct = .10

// randomRampStrategy starts sets of users with random delays
type randomRampStrategy struct {
	rampPeriodPct float64
}


func (r *randomRampStrategy) GetStartProfile(testLength time.Duration, maxUsers int) []StartProfile {
	rampPeriod := time.Duration(int64(testLength.Seconds() * r.rampPeriodPct)) * time.Second

	usersPerRamp := int(float32(maxUsers) * RampUsersPct)

	strategies := make([]StartProfile, 0)

	var assignedUsers int
	for assignedUsers = 0; assignedUsers < maxUsers; assignedUsers += usersPerRamp {
		strategies = append(strategies, StartProfile{Delay: utils.RandomSecondDuration(time.Duration(0), rampPeriod), Users: usersPerRamp})
	}

	if maxUsers - assignedUsers > 0 {
		strategies = append(strategies, StartProfile{Delay: utils.RandomSecondDuration(time.Duration(0), rampPeriod), Users: maxUsers - assignedUsers})
	}

	return strategies
}

func NewRandomRampUpStrategy(rampPeriod float64) RampStrategy {
	return &randomRampStrategy{rampPeriodPct: rampPeriod}
}