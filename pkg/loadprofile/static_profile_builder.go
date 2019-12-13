package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
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


func (s *staticProfileBuilder) GetLoadProfiles(runFunc testwrapper.Test, resultCollector collector.ResultCollector) []LoadProfile {
	runners := make([]LoadProfile, 0)

	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)

	startDelays := s.rampUpStrategy.GetStartDelay(s.testLength, s.concurrentUsers)

	for _, sd := range startDelays {
		initialDelayStep := steps.NewInitialWaitStep(sd.InitialDelay)
		waitBetweenRuns := steps.NewWaitStep(s.interval)

		runProfile := []steps.Step{initialDelayStep, runFuncStep, waitBetweenRuns}
		for i := 0; i < sd.UsersToStart; i++ {
			runners = append(runners, NewLoadProfile(runProfile, s.testLength - sd.InitialDelay, false))
		}
	}

	return runners
}


func NewStaticProfileBuilder(concurrentUsers int, testLengthSec int, intervalSec int, rampUpStrategy  rampstrategy.RampStrategy) LoadProfileBuilder {
	testLength := time.Duration(testLengthSec)*time.Second
	testInterval := time.Duration(intervalSec)*time.Second
	return &staticProfileBuilder{concurrentUsers: concurrentUsers,
		testLength: testLength,
		interval: testInterval,
		rampUpStrategy: rampUpStrategy,
		}
}