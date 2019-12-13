package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

// randomProfileBuilder runs each test using a random wait period between executions
type randomProfileBuilder struct {
	concurrentUsers			int
	testLength				time.Duration
	rampUpStrategy			rampstrategy.RampStrategy
}


func (s *randomProfileBuilder) GetLoadProfiles(runFunc testwrapper.Test, resultCollector collector.ResultCollector) []LoadProfile {
	runners := make([]LoadProfile, 0)

	startDelays := s.rampUpStrategy.GetStartProfile(s.testLength, s.concurrentUsers)

	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)

	for _, sd := range startDelays {
		initialDelayStep := steps.NewInitialWaitStep(sd.Delay)

		randWaitStep := steps.NewRandomWaitStep(time.Duration(0), time.Duration(int(s.testLength/time.Millisecond/4))*time.Millisecond, utils.RandomSecondDuration)

		runProfile := []steps.Step{initialDelayStep, runFuncStep, randWaitStep}
		for i := 0; i < sd.Users; i++ {
			runners = append(runners, NewLoadProfile(runProfile, s.testLength - sd.Delay, false))
		}
	}

	return runners
}


func NewRandomProfileBuilder(concurrentUsers int, testLengthSec int, rampUpStrategy  rampstrategy.RampStrategy) LoadProfileBuilder {
	testLength := time.Duration(testLengthSec)*time.Second
	return &randomProfileBuilder{concurrentUsers: concurrentUsers,
		testLength: testLength,
		rampUpStrategy: rampUpStrategy,
	}
}