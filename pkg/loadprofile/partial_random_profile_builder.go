package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/logging"
	"github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/testscenario"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

// partialRandomProfileBuilder runs each test using a random wait period between executions for some set of users
// and a fixed interval for others
type partialRandomProfileBuilder struct {
	concurrentUsers			int
	testLength				time.Duration
	interval				time.Duration
	rampUpStrategy			rampstrategy.RampStrategy
}


func (p *partialRandomProfileBuilder) GetLoadProfiles(ctx context.Context, runFunc testscenario.Test, resultCollector collector.ResultCollector) []LoadProfile {
	runners := make([]LoadProfile, 0)
	logger, ctx := logging.GetLoggerFromContext(ctx, p)

	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)


	startDelays := p.rampUpStrategy.GetStartProfile(ctx, p.testLength, p.concurrentUsers)

	for i, sd := range startDelays {
		logger.Infof(ctx,"Creating start set of %d users with initial delay of %v with max execution interval %v", sd.Users, sd.Delay, p.interval)
		initialDelayStep := steps.NewInitialWaitStep(sd.Delay)

		var waitStep steps.Step
		if i % 2 == 0 {
			waitStep = steps.NewWaitStep(p.interval)
		 } else {
			waitStep = steps.NewRandomWaitStep(time.Duration(0), p.interval, utils.RandomSecondDuration)
		}

		runProfile := []steps.Step{initialDelayStep, runFuncStep, waitStep}
		for i := 0; i < sd.Users; i++ {
			runners = append(runners, NewLoadProfile(runProfile, p.testLength - sd.Delay, false))
		}
	}

	return runners
}


func NewPartialRandomProfileBuilder(concurrentUsers int, testLength time.Duration, interval time.Duration, rampstrategy rampstrategy.RampStrategy) LoadProfileBuilder {
	return &partialRandomProfileBuilder{concurrentUsers: concurrentUsers,
		testLength: testLength,
		interval: interval,
		rampUpStrategy: rampstrategy,
	}
}