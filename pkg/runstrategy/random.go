package runstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
	"time"
)

type randomDelayStrategyFactory struct {
	delayMinSec		int
	delayMaxSec		int
}

func (fact *randomDelayStrategyFactory) GetRunStrategy(testId string, initialDelay int, funcToRun utils.RunFunc, resultCollector collector.ResultCollector) RunStrategy {
	return &randomDelayStrategy{testId: testId,
		initialDelay: initialDelay,
		delayMinSec: fact.delayMinSec,
		delayMaxSec: fact.delayMaxSec,
		funcToRun: funcToRun,
		collector: resultCollector}
}

type randomDelayStrategy struct {
	testId			string
	initialDelay	int
	delayMinSec		int
	delayMaxSec		int
	ticker			*time.Ticker
	funcToRun		utils.RunFunc
	collector		collector.ResultCollector
	stopped			bool
}

func (r *randomDelayStrategy) Start(wg sync.WaitGroup) {
	wg.Add(1)

	if r.initialDelay > 0 {
		utils.Logt(r.testId, ": pausing for initial delay")
		time.Sleep(time.Second * time.Duration(r.initialDelay))
	}

	utils.Logt(r.testId, ": starting test")
	ticker := r.newTicker()
	for !r.stopped {
		select {
		case <-ticker.C:
			runTest(r.testId, r.funcToRun, r.collector)
			ticker = r.newTicker()
		}
	}
	utils.Logt(r.testId, ": process stopped")
	wg.Done()
}

func (r *randomDelayStrategy) Stop() {
	r.stopped = true
}

func (r *randomDelayStrategy) GetResults() collector.ResultCollector {
	return r.collector
}


func (r *randomDelayStrategy) newTicker() *time.Ticker {
	randInterval := time.Duration(utils.RandomIntBetween(r.delayMinSec, r.delayMaxSec))
	return time.NewTicker(time.Duration(time.Second * randInterval))
}

func NewRandomDelayRunStrategyFactory(minDelaySec int, maxDelaySec int) RunStrategyFactory {
	return &randomDelayStrategyFactory{delayMinSec: minDelaySec, delayMaxSec: maxDelaySec}
}

