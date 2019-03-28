package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type randomProfile struct {
	concurrentUsers			int
	testLength				time.Duration
}


func (s *randomProfile) GetLoad(namer naming.TestNamer, runFunc utils.RunFunc, resultCollector collector.ResultCollector) []Load {
	runners := make([]Load, 0)

	randWaitStep := NewRandomWaitStep(time.Duration(0), time.Duration(int(s.testLength/time.Millisecond/4))*time.Millisecond, utils.RandomSecondDuration)
	runFuncStep := NewRunFuncStep(runFunc, resultCollector)
	compositeStep := NewCompositeStep(time.Duration(0), false, runFuncStep, randWaitStep)
	runForStep := NewRunForStep(s.testLength, compositeStep, time.Duration(0), false)

	runProfile := []Step {randWaitStep, runForStep}

	for i := 0; i < s.concurrentUsers; i++ {
		ctx := context.WithValue(context.Background(), "testId", namer.GetName(i))
		runners = append(runners, NewLoad(ctx, runProfile))
	}

	return runners
}


func NewRandomProfile(concurrentUsers int, testLengthSec int) LoadProfile {
	testLength := time.Duration(testLengthSec)*time.Second
	return &randomProfile{concurrentUsers: concurrentUsers,
		testLength: testLength,
		}
}