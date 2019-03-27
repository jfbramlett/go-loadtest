package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type randomProfile struct {
	concurrentUsers			int
	testLength				time.Duration
}


func (s *randomProfile) GetRunners(namer naming.TestNamer, runFunc utils.RunFunc, resultCollector collector.ResultCollector) []Runner {
	runners := make([]Runner, 0)

	for i := 0; i < s.concurrentUsers; i++ {
		startDelay := utils.RandomIntBetween(0, int(s.testLength/time.Millisecond/4))
		runInterval := utils.RandomIntBetween(0, int(s.testLength/time.Millisecond/10))
		step := RunStep{startDelay: time.Duration(startDelay)*time.Millisecond, runTime: s.testLength, invocationDelay: time.Duration(runInterval)*time.Millisecond}
		steps := []RunStep {step}

		runners = append(runners, NewRunner(namer.GetName(i), steps, runFunc, resultCollector))
	}

	return runners
}


func NewRandomProfile(concurrentUsers int, testLengthSec int) LoadProfile {
	testLength := time.Duration(testLengthSec)*time.Second
	return &randomProfile{concurrentUsers: concurrentUsers,
		testLength: testLength,
		}
}