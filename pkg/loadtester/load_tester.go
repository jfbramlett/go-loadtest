package loadtester

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/idgenerator"
	"github.com/jfbramlett/go-loadtest/pkg/loadprofile"
	"github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/testscenario"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
	"time"
)

type LoadTester interface {
	Run(ctx context.Context, runFunc testscenario.Test) collector.ResultCollector
}


type DefaultLoadTester struct {
	loadProfileBuilder loadprofile.LoadProfileBuilder
	idGenerator        idgenerator.TestIdGenerator
	resultCollector    collector.ResultCollector
}

func (l *DefaultLoadTester) Run(ctx context.Context, runFunc testscenario.Test) collector.ResultCollector {
	logger := utils.LoggerFromContext(ctx)

	logger.Info( "Starting results collector")
	l.resultCollector.Start()

	logger.Info( "Starting runners")
	wg := sync.WaitGroup{}
	for i, r := range l.loadProfileBuilder.GetLoadProfiles(ctx, runFunc, l.resultCollector) {
		wg.Add(1)
		ctx := utils.SetTestIdInContext(ctx, l.idGenerator.GetId(i))
		go l.runWrapper(r, ctx, &wg)
	}

	logger.Info("Waiting for tests to end")
	wg.Wait()
	logger.Info("Tests completed")
	l.resultCollector.Stop()

	return l.resultCollector
}


func (l *DefaultLoadTester) runWrapper(load loadprofile.LoadProfile, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	load.Run(ctx)
}


func NewDefaultLoadTester(loadProfileBuilder loadprofile.LoadProfileBuilder, idGenerator idgenerator.TestIdGenerator, resultCollector collector.ResultCollector) LoadTester {
	return &DefaultLoadTester{loadProfileBuilder: loadProfileBuilder,
		idGenerator: idGenerator,
		resultCollector: resultCollector,
	}
}


func NewLoadTester(concurrentUsers int, testLength time.Duration, testInterval time.Duration,
	profileType loadprofile.LoadProfileType, strategyType rampstrategy.RampStrategyType) LoadTester {

	loadProfileBuilder := loadprofile.NewLoadProfileBuilder(profileType, concurrentUsers, testLength, testInterval, strategyType)

	return &DefaultLoadTester{loadProfileBuilder: loadProfileBuilder,
		idGenerator:     idgenerator.NewSimpleTestIdGenerator(),
		resultCollector: collector.NewInMemoryRunCollector(),
	}
}
