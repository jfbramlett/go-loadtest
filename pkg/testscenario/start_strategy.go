package testscenario

import (
	"context"
	"time"

	"github.com/jfbramlett/go-loadtest/pkg/utils"
)

const (
	DefaultRampPeriod = .10
	RampUsersPct      = .10
	MaxRamps          = 10
)

type StartProfile struct {
	Delay time.Duration
	Users int
}

type StartStrategyFunc func(ctx context.Context, testLength time.Duration, rampToUsers int) []StartProfile

func NoopStartStrategy(ctx context.Context, testLength time.Duration, rampToUsers int) []StartProfile {
	return []StartProfile{
		{Delay: time.Duration(0), Users: rampToUsers},
	}
}

func RandomStartStrategy(ctx context.Context, testLength time.Duration, rampToUsers int) []StartProfile {
	rampPeriod := time.Duration(int64(testLength.Seconds()*DefaultRampPeriod)) * time.Second

	usersPerRamp := int(float32(rampToUsers) * RampUsersPct)

	strategies := make([]StartProfile, 0)

	startDelays := make(map[time.Duration]struct{})
	var assignedUsers int
	for assignedUsers = 0; assignedUsers < rampToUsers; assignedUsers += usersPerRamp {
		delay := getStartDelay(rampPeriod, startDelays)
		startDelays[delay] = empty
		strategies = append(strategies, StartProfile{Delay: delay, Users: usersPerRamp})
	}

	if rampToUsers-assignedUsers > 0 {
		delay := getStartDelay(rampPeriod, startDelays)
		strategies = append(strategies, StartProfile{Delay: delay, Users: rampToUsers - assignedUsers})
	}

	return strategies
}

func getStartDelay(rampPeriod time.Duration, existingDelays map[time.Duration]struct{}) time.Duration {
	for {
		delay := utils.RandomMilliSecondDuration(time.Duration(0), rampPeriod)
		if _, found := existingDelays[delay]; !found {
			return delay
		}
	}
}

func SmoothStartStrategy(ctx context.Context, testLength time.Duration, rampToUsers int) []StartProfile {
	rampPeriodSec := int64(testLength.Seconds() * DefaultRampPeriod)

	rampIntervals := rampPeriodSec / MaxRamps

	usersPerRamp := rampToUsers / MaxRamps
	lastRamp := rampToUsers % MaxRamps

	strategies := make([]StartProfile, 0)
	wait := int64(0)
	for i := 0; i < MaxRamps; i++ {
		strategies = append(strategies, StartProfile{Delay: time.Duration(wait) * time.Second, Users: usersPerRamp})
		wait = wait + rampIntervals
	}
	if lastRamp > 0 {
		strategies = append(strategies, StartProfile{Delay: time.Duration(wait) * time.Second, Users: lastRamp})
	}

	return strategies
}
