package loadprofile

import (
	"context"
	"github.com/ninthwave/nwp-load-test/pkg/collector"
	"github.com/ninthwave/nwp-load-test/pkg/logging"
	"github.com/ninthwave/nwp-load-test/pkg/rampstrategy"
	"github.com/ninthwave/nwp-load-test/pkg/testscenario"
	"github.com/ninthwave/nwp-load-test/pkg/steps"
	"github.com/ninthwave/nwp-load-test/pkg/utils"
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
	logger, ctx := logging.GetLoggerFromContext(ctx, s)

	startDelays := s.rampUpStrategy.GetStartProfile(ctx, s.testLength, s.concurrentUsers)

	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)

	for _, sd := range startDelays {
		logger.Infof(ctx,"Creating start set of %d users with initial delay of %v with max execution interval %v", sd.Users, sd.Delay, s.maxInterval)

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