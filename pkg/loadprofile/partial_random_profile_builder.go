package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type partialRandomProfileBuilder struct {
	concurrentUsers			int
	testLength				time.Duration
	interval				time.Duration
}


func (p *partialRandomProfileBuilder) GetLoadProfiles(runFunc testwrapper.Test, resultCollector collector.ResultCollector) []LoadProfile {
	runners := make([]LoadProfile, 0)

	randWaitStep := steps.NewRandomWaitStep(time.Duration(0), time.Duration(int(p.testLength/time.Millisecond/4))*time.Millisecond, utils.RandomSecondDuration)
	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)
	compositeStep := steps.NewCompositeStep(time.Duration(0), false, runFuncStep, randWaitStep)
	runForStep := steps.NewRunForStep(p.testLength, compositeStep, time.Duration(0), false)
	randomRunProfile := []steps.Step{randWaitStep, runForStep}

	fixedRunForStep := steps.NewRunForStep(p.testLength, runFuncStep, p.interval, false)
	fixedRunProfile := []steps.Step{fixedRunForStep}


	for i := 0; i < p.concurrentUsers; i++ {
		if i % 2 == 0 {
			runners = append(runners, NewLoadProfile(randomRunProfile))
		} else {
			runners = append(runners, NewLoadProfile(fixedRunProfile))
		}
	}

	return runners
}


func NewPartialRandomProfileBuilder(concurrentUsers int, testLengthSec int, intervalSec int) LoadProfileBuilder {
	testLength := time.Duration(testLengthSec)*time.Second
	testInterval := time.Duration(intervalSec)*time.Second
	return &partialRandomProfileBuilder{concurrentUsers: concurrentUsers,
		testLength: testLength,
		interval: testInterval,
	}
}