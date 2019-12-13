package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type partialRandomProfileBuilder struct {
	concurrentUsers			int
	testLength				time.Duration
	interval				time.Duration
	rampUpStrategy			rampstrategy.RampStrategy
}


func (p *partialRandomProfileBuilder) GetLoadProfiles(runFunc testwrapper.Test, resultCollector collector.ResultCollector) []LoadProfile {
	runners := make([]LoadProfile, 0)

	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)


	startDelays := p.rampUpStrategy.GetStartDelay(p.testLength, p.concurrentUsers)

	for i, sd := range startDelays {
		initialDelayStep := steps.NewInitialWaitStep(sd.InitialDelay)

		var waitStep steps.Step
		if i % 2 == 0 {
			waitStep = steps.NewWaitStep(p.interval)
		 } else {
			waitStep = steps.NewRandomWaitStep(time.Duration(0), time.Duration(int(p.testLength/time.Millisecond/4))*time.Millisecond, utils.RandomSecondDuration)
		}

		runProfile := []steps.Step{initialDelayStep, runFuncStep, waitStep}
		for i := 0; i < sd.UsersToStart; i++ {
			runners = append(runners, NewLoadProfile(runProfile, p.testLength - sd.InitialDelay, false))
		}
	}

	return runners
}


func NewPartialRandomProfileBuilder(concurrentUsers int, testLengthSec int, intervalSec int, rampstrategy rampstrategy.RampStrategy) LoadProfileBuilder {
	testLength := time.Duration(testLengthSec)*time.Second
	testInterval := time.Duration(intervalSec)*time.Second
	return &partialRandomProfileBuilder{concurrentUsers: concurrentUsers,
		testLength: testLength,
		interval: testInterval,
		rampUpStrategy: rampstrategy,
	}
}