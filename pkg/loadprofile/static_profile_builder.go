package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/logging"
	"github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"time"
)

// staticProfileBuilder runs each test with a fixed wait between executions
type staticProfileBuilder struct {
	concurrentUsers			int
	testLength				time.Duration
	interval				time.Duration
	rampUpStrategy			rampstrategy.RampStrategy
}


func (s *staticProfileBuilder) GetLoadProfiles(ctx context.Context, runFunc testwrapper.Test, resultCollector collector.ResultCollector) []LoadProfile {
	runners := make([]LoadProfile, 0)
	logger, ctx := logging.GetLoggerFromContext(ctx, s)

	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)

	startDelays := s.rampUpStrategy.GetStartProfile(ctx, s.testLength, s.concurrentUsers)

	for _, sd := range startDelays {
		logger.Infof(ctx,"Creating start set of %d users with initial delay of %v with execution interval %v", sd.Users, sd.Delay, s.interval)
		initialDelayStep := steps.NewInitialWaitStep(sd.Delay)
		waitBetweenRuns := steps.NewWaitStep(s.interval)

		runProfile := []steps.Step{initialDelayStep, runFuncStep, waitBetweenRuns}
		for i := 0; i < sd.Users; i++ {
			runners = append(runners, NewLoadProfile(runProfile, s.testLength - sd.Delay, false))
		}
	}

	return runners
}


func NewStaticProfileBuilder(concurrentUsers int, testLength time.Duration, interval time.Duration, rampUpStrategy  rampstrategy.RampStrategy) LoadProfileBuilder {
	return &staticProfileBuilder{concurrentUsers: concurrentUsers,
		testLength: testLength,
		interval: interval,
		rampUpStrategy: rampUpStrategy,
		}
}