package testscenario

import (
	"context"
	"sync"
	"time"

	"github.com/jfbramlett/go-loadtest/pkg/utils"
)

type TestScenario struct {
	name          string
	test          TestFunc
	setup         SetupFunc
	teardown      TeardownFunc
	maxUsers      int
	testLength    time.Duration
	startStrategy StartStrategyFunc
	pauseFunc     PauseStrategyFunc
}

func NewTestScenario(name string, test TestFunc, setup SetupFunc, teardown TeardownFunc, maxUsers int, testLength time.Duration, startStrategy StartStrategyFunc, pause PauseStrategyFunc) *TestScenario {
	return &TestScenario{
		name:          name,
		test:          test,
		setup:         setup,
		teardown:      teardown,
		maxUsers:      maxUsers,
		startStrategy: startStrategy,
		pauseFunc:     pause,
		testLength:    testLength,
	}
}

func (ts *TestScenario) Run(ctx context.Context, resultsCollector ResultCollector) error {
	var err error

	logger := utils.LoggerFromContext(ctx)
	logger = logger.WithField("name", ts.name)

	if ts.setup != nil {
		ctx, err = ts.setup(ctx)
		if err != nil {
			return err
		}
	}

	startProfiles := ts.startStrategy(ctx, ts.testLength, ts.maxUsers)

	wg := sync.WaitGroup{}
	for _, sp := range startProfiles {
		logger.Infof("starting %d users with start delay %dms", sp.Users, sp.Delay.Milliseconds())
		for i := 0; i < sp.Users; i++ {
			wg.Add(1)
			go func() {
				start := time.Now()
				time.Sleep(sp.Delay)
				for time.Since(start) < ts.testLength {
					ts.test(ctx, resultsCollector)
					if time.Since(start) < ts.testLength && ts.pauseFunc != nil {
						ts.pauseFunc(ctx)
					}
				}
				wg.Done()
			}()
		}
	}

	wg.Wait()

	if ts.teardown != nil {
		err = ts.teardown(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
