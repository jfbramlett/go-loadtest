package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type randomProfileBuilder struct {
	concurrentUsers			int
	testLength				time.Duration
}


func (s *randomProfileBuilder) GetLoadProfiles(runFunc testwrapper.Test, resultCollector collector.ResultCollector) []LoadProfile {
	runners := make([]LoadProfile, 0)

	randWaitStep := steps.NewRandomWaitStep(time.Duration(0), time.Duration(int(s.testLength/time.Millisecond/4))*time.Millisecond, utils.RandomSecondDuration)
	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)
	compositeStep := steps.NewCompositeStep(time.Duration(0), false, runFuncStep, randWaitStep)
	runForStep := steps.NewRunForStep(s.testLength, compositeStep, time.Duration(0), false)

	runProfile := []steps.Step{randWaitStep, runForStep}

	for i := 0; i < s.concurrentUsers; i++ {
		runners = append(runners, NewLoadProfile(runProfile))
	}

	return runners
}


func NewRandomProfileBuilder(concurrentUsers int, testLengthSec int) LoadProfileBuilder {
	testLength := time.Duration(testLengthSec)*time.Second
	return &randomProfileBuilder{concurrentUsers: concurrentUsers,
		testLength: testLength,
		}
}