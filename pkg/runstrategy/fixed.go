package runstrategy

import (
	"fmt"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
	"time"
)

type fixedDelayStrategyFactory struct {
	delaySec		time.Duration
}

func (fact *fixedDelayStrategyFactory) GetRunStrategy(runId string, initialDelay int, funcToRun utils.RunFunc, resultCollector utils.ResultCollector) RunStrategy {
	fmt.Println(fmt.Sprintf("%s: using fixed delay of %d s with initial delay %d s", runId, fact.delaySec, initialDelay))
	return &fixedDelayStrategy{runId: runId,
		initialDelay: initialDelay,
		ticker: time.NewTicker(time.Duration(time.Second * fact.delaySec)),
		funcToRun: funcToRun,
		collector: resultCollector}
}

type fixedDelayStrategy struct {
	runId				string
	initialDelay		int
	ticker				*time.Ticker
	funcToRun			utils.RunFunc
	collector			utils.ResultCollector
}

func (f *fixedDelayStrategy) Start(wg sync.WaitGroup) {
	if f.initialDelay > 0 {
		fmt.Println(f.runId + ": pausing for initial delay")
		time.Sleep(time.Second * time.Duration(f.initialDelay))
	}

	fmt.Println(f.runId + ": process running")
	wg.Add(1)
	for range f.ticker.C {
		runTest(f.funcToRun, f.collector)
	}
	fmt.Println(f.runId + ": process stopped")
	wg.Done()
}

func (f *fixedDelayStrategy) Stop() {
	f.ticker.Stop()
}

func (f *fixedDelayStrategy) GetResults() utils.ResultCollector {
	return f.collector
}



func NewFixedDelayRunStrategyFactory(delaySec int64) RunStrategyFactory {
	fmt.Println(fmt.Sprintf("Using fixed delay strategy of %d s", delaySec))
	return &fixedDelayStrategyFactory{delaySec: time.Duration(delaySec)}
}

