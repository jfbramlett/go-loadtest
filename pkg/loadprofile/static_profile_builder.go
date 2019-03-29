package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type staticProfileBuilder struct {
	concurrentUsers			int
	testLength				time.Duration
	interval				time.Duration
}


func (s *staticProfileBuilder) GetLoadProfiles(runFunc utils.RunFunc, resultCollector collector.ResultCollector) []LoadProfile {
	runners := make([]LoadProfile, 0)

	runFuncStep := steps.NewRunFuncStep(runFunc, resultCollector)
	runForStep := steps.NewRunForStep(s.testLength, runFuncStep, s.interval, false)
	runProfile := []steps.Step{runForStep}

	for i := 0; i < s.concurrentUsers; i++ {
		runners = append(runners, NewLoadProfile(runProfile))
	}

	return runners
}


func NewStaticProfileBuilder(concurrentUsers int, testLengthSec int, intervalSec int) LoadProfileBuilder {
	testLength := time.Duration(testLengthSec)*time.Second
	testInterval := time.Duration(intervalSec)*time.Second
	return &staticProfileBuilder{concurrentUsers: concurrentUsers,
		testLength: testLength,
		interval: testInterval,
		}
}