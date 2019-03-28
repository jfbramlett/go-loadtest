package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type staticProfile struct {
	concurrentUsers			int
	testLength				time.Duration
	interval				time.Duration
}


func (s *staticProfile) GetLoad(namer naming.TestNamer, runFunc utils.RunFunc, resultCollector collector.ResultCollector) []Load {
	runners := make([]Load, 0)

	runFuncStep := NewRunFuncStep(runFunc, resultCollector)
	runForStep := NewRunForStep(s.testLength, runFuncStep, s.interval, false)
	runProfile := []Step {runForStep}

	for i := 0; i < s.concurrentUsers; i++ {
		ctx := context.WithValue(context.Background(), "testId", namer.GetName(i))
		runners = append(runners, NewLoad(ctx, runProfile))
	}

	return runners
}


func NewStaticProfile(concurrentUsers int, testLengthSec int, intervalSec int) LoadProfile {
	testLength := time.Duration(testLengthSec)*time.Second
	testInterval := time.Duration(intervalSec)*time.Second
	return &staticProfile{concurrentUsers: concurrentUsers,
		testLength: testLength,
		interval: testInterval,
		}
}