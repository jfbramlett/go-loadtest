package loadprofile

import (
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


func (s *staticProfile) GetRunners(namer naming.TestNamer, runFunc utils.RunFunc, resultCollector collector.ResultCollector) []Runner {
	runners := make([]Runner, 0)
	step := RunStep{startDelay: time.Duration(0), runTime: s.testLength, invocationDelay: s.interval}
	steps := []RunStep {step}

	for i := 0; i < s.concurrentUsers; i++ {
		runners = append(runners, NewRunner(namer.GetName(i), steps, runFunc, resultCollector))
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