package loadtester

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/loadprofile"
	"github.com/jfbramlett/go-loadtest/pkg/logging"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
	"time"
)

type LoadTester interface {
	Run(ctx context.Context, runFunc testwrapper.Test) collector.ResultCollector
}


type DefaultLoadTester struct {
	loadProfileBuilder		loadprofile.LoadProfileBuilder
	namer					naming.TestNamer
	resultCollector			collector.ResultCollector
}

func (l *DefaultLoadTester) Run(ctx context.Context, runFunc testwrapper.Test) collector.ResultCollector {
	logger, ctx := logging.GetLoggerFromContext(ctx, l)

	logger.Info(ctx, "Starting runners")
	l.resultCollector.Start()

	wg := sync.WaitGroup{}
	for i, r := range l.loadProfileBuilder.GetLoadProfiles(runFunc, l.resultCollector) {
		wg.Add(1)
		ctx := utils.SetTestIdInContext(ctx, l.namer.GetName(i))
		go l.runWrapper(r, ctx, &wg)
	}

	logger.Info(context.Background(), "Waiting for tests to end")
	wg.Wait()
	logger.Info(context.Background(), "Tests completed")
	l.resultCollector.Stop()

	return l.resultCollector
}


func (l *DefaultLoadTester) runWrapper(load loadprofile.LoadProfile, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	load.Run(ctx)
}


func NewDefaultLoadTester(loadProfileBuilder loadprofile.LoadProfileBuilder, namer naming.TestNamer, resultCollector collector.ResultCollector) LoadTester {
	return &DefaultLoadTester{loadProfileBuilder: loadProfileBuilder,
		namer: namer,
		resultCollector: resultCollector,
	}
}


func NewLoadTester(concurrentUsers int, testLength time.Duration, testInterval time.Duration,
	profileType loadprofile.LoadProfileType, strategyType rampstrategy.RampStrategyType) LoadTester {

	loadProfileBuilder := loadprofile.NewLoadProfileBuilder(profileType, concurrentUsers, testLength, testInterval, strategyType)

	return &DefaultLoadTester{loadProfileBuilder: loadProfileBuilder,
		namer: naming.NewSimpleTestNamer(),
		resultCollector: collector.NewInMemoryRunCollector(),
	}
}
