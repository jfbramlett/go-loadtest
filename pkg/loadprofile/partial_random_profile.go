package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type partialRandomProfile struct {
	concurrentUsers			int
	testLength				time.Duration
	interval				time.Duration
}


func (p *partialRandomProfile) GetLoad(namer naming.TestNamer, runFunc utils.RunFunc, resultCollector collector.ResultCollector) []Load {
	runners := make([]Load, 0)

	randWaitStep := NewRandomWaitStep(time.Duration(0), time.Duration(int(p.testLength/time.Millisecond/4))*time.Millisecond, utils.RandomSecondDuration)
	runFuncStep := NewRunFuncStep(runFunc, resultCollector)
	compositeStep := NewCompositeStep(time.Duration(0), false, runFuncStep, randWaitStep)
	runForStep := NewRunForStep(p.testLength, compositeStep, time.Duration(0), false)
	randomRunProfile := []Step {randWaitStep, runForStep}

	fixedRunForStep := NewRunForStep(p.testLength, runFuncStep, p.interval, false)
	fixedRunProfile := []Step {fixedRunForStep}


	for i := 0; i < p.concurrentUsers; i++ {
		ctx := context.WithValue(context.Background(), "testId", namer.GetName(i))

		if i % 2 == 0 {
			runners = append(runners, NewLoad(ctx, randomRunProfile))
		} else {
			runners = append(runners, NewLoad(ctx, fixedRunProfile))
		}
	}

	return runners
}


func NewPartialRandomProfile(concurrentUsers int, testLengthSec int, intervalSec int) LoadProfile {
	testLength := time.Duration(testLengthSec)*time.Second
	testInterval := time.Duration(intervalSec)*time.Second
	return &partialRandomProfile{concurrentUsers: concurrentUsers,
		testLength: testLength,
		interval: testInterval,
	}
}