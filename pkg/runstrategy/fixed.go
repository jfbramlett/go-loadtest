package runstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
	"time"
)

type fixedDelayStrategyFactory struct {
	delaySec		time.Duration
}

func (fact *fixedDelayStrategyFactory) GetRunStrategy(testId string, initialDelay int, funcToRun utils.RunFunc, resultCollector collector.ResultCollector) RunStrategy {
	return &fixedDelayStrategy{testId: testId,
		initialDelay: initialDelay,
		ticker: time.NewTicker(time.Duration(time.Second * fact.delaySec)),
		funcToRun: funcToRun,
		collector: resultCollector}
}

type fixedDelayStrategy struct {
	testId				string
	initialDelay		int
	ticker				*time.Ticker
	funcToRun			utils.RunFunc
	collector			collector.ResultCollector
}

func (f *fixedDelayStrategy) Start(wg sync.WaitGroup) {
	wg.Add(1)
	if f.initialDelay > 0 {
		utils.Logt(f.testId, ": pausing for initial delay")
		time.Sleep(time.Second * time.Duration(f.initialDelay))
	}

	utils.Logt(f.testId, ": process running")
	for range f.ticker.C {
		runTest(f.testId, f.funcToRun, f.collector)
	}
	utils.Logt(f.testId, ": process stopped")
	wg.Done()
}

func (f *fixedDelayStrategy) Stop() {
	f.ticker.Stop()
}

func (f *fixedDelayStrategy) GetResults() collector.ResultCollector {
	return f.collector
}



func NewFixedDelayRunStrategyFactory(delaySec int64) RunStrategyFactory {
	return &fixedDelayStrategyFactory{delaySec: time.Duration(delaySec)}
}

