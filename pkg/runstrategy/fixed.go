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

func (fact *fixedDelayStrategyFactory) GetRunStrategy(funcToRun utils.RunFunc, resultCollector utils.ResultCollector) RunStrategy {
	fmt.Println(fmt.Sprintf("using fixed delay of %d ms\n", fact.delaySec))
	return &fixedDelayStrategy{ticker: time.NewTicker(time.Duration(time.Second * fact.delaySec)),
		funcToRun: funcToRun,
		collector: resultCollector}
}

type fixedDelayStrategy struct {
	ticker		*time.Ticker
	funcToRun	utils.RunFunc
	collector	utils.ResultCollector
}

func (f *fixedDelayStrategy) Start(wg sync.WaitGroup) {
	wg.Add(1)
	for range f.ticker.C {
		runTest(f.funcToRun, f.collector)
	}
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

