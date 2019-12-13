package rampstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

const RampUsersPct = .10

type randomRampUpStrategy struct {
	rampPeriodPct float64
}


func (r *randomRampUpStrategy) GetStartDelay(testLength time.Duration, maxUsers int) []StartDelay {
	rampPeriod := time.Duration(int64(testLength.Seconds() * r.rampPeriodPct)) * time.Second

	usersPerRamp := int(float32(maxUsers) * RampUsersPct)

	strategies := make([]StartDelay, 0)

	var assignedUsers int
	for assignedUsers = 0; assignedUsers < maxUsers; assignedUsers += usersPerRamp {
		strategies = append(strategies, StartDelay{InitialDelay: utils.RandomSecondDuration(time.Duration(0), rampPeriod), UsersToStart: usersPerRamp})
	}

	if maxUsers - assignedUsers > 0 {
		strategies = append(strategies, StartDelay{InitialDelay: utils.RandomSecondDuration(time.Duration(0), rampPeriod), UsersToStart: maxUsers - assignedUsers})
	}

	return strategies
}

func NewRandomRampUpStrategy(rampPeriod float64) RampStrategy {
	return &randomRampUpStrategy{rampPeriodPct: rampPeriod}
}