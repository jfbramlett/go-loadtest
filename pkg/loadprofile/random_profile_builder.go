package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"github.com/jfbramlett/go-loadtest/pkg/testscenario"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

// randomProfileBuilder runs each test using a random wait period between executions
type randomProfileBuilder struct {
	concurrentUsers			int
	testLength				time.Duration
	maxInterval				time.Duration
	rampUpStrategy			rampstrategy.RampStrategy
}


func (s *randomProfileBuilder) GetLoadProfiles(ctx context.Context, runFunc testscenario.Test, resultCollector collector.ResultCollector) []LoadProfile {
	runners := make([]LoadProfile, 0)
	logger := utils.LoggerFromContext(ctx)

	startDelays := s.rampUpStrategy.GetStartProfile(ctx, s.testLength, s.concurrentUsers)

	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)

	for _, sd := range startDelays {
		logger.Infof("Creating start set of %d users with initial delay of %v with max execution interval %v", sd.Users, sd.Delay, s.maxInterval)

		initialDelayStep := steps.NewInitialWaitStep(sd.Delay)

		randWaitStep := steps.NewRandomWaitStep(time.Duration(0), s.maxInterval, utils.RandomSecondDuration)

		runProfile := []steps.Step{initialDelayStep, runFuncStep, randWaitStep}
		for i := 0; i < sd.Users; i++ {
			runners = append(runners, NewLoadProfile(runProfile, s.testLength - sd.Delay, false))
		}
	}

	return runners
}


func NewRandomProfileBuilder(concurrentUsers int, testLength time.Duration, maxInterval time.Duration, rampUpStrategy rampstrategy.RampStrategy) LoadProfileBuilder {
	return &randomProfileBuilder{concurrentUsers: concurrentUsers,
		testLength: testLength,
		maxInterval: maxInterval,
		rampUpStrategy: rampUpStrategy,
	}
}