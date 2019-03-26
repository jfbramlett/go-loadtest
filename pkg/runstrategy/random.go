package runstrategy

import (
	"fmt"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"math/rand"
	"sync"
	"time"
)

type randomDelayStrategyFactory struct {
	delayMinSec		int
	delayMaxSec		int
}

func (rfact *randomDelayStrategyFactory) GetRunStrategy(funcToRun utils.RunFunc, resultCollector utils.ResultCollector) RunStrategy {
	fmt.Println(fmt.Sprintf("using random delay between%d ms and %d ms\n", rfact.delayMinSec, rfact.delayMaxSec))
	return &randomDelayStrategy{delayMinSec: rfact.delayMinSec,
		delayMaxSec: rfact.delayMaxSec,
		funcToRun: funcToRun,
		collector: resultCollector}
}

type randomDelayStrategy struct {
	delayMinSec		int
	delayMaxSec		int
	ticker			*time.Ticker
	funcToRun		utils.RunFunc
	collector		utils.ResultCollector
	stopped			bool
}

func (r *randomDelayStrategy) Start(wg sync.WaitGroup) {
	wg.Add(1)
	ticker := r.newTicker()
	for !r.stopped {
		select {
		case <-ticker.C:
			runTest(r.funcToRun, r.collector)
			ticker = r.newTicker()
		}
	}
	wg.Done()
}

func (r *randomDelayStrategy) Stop() {
	r.stopped = true
}

func (r *randomDelayStrategy) GetResults() utils.ResultCollector {
	return r.collector
}


func (r *randomDelayStrategy) newTicker() *time.Ticker {
	randInterval := time.Duration(rand.Intn(r.delayMaxSec - r.delayMinSec) + r.delayMinSec)
	fmt.Println(fmt.Sprintf("using random delay of %d sec", randInterval))
	return time.NewTicker(time.Duration(time.Second * randInterval))
}

func NewRandomDelayRunStrategyFactory(minDelaySec int, maxDelaySec int) RunStrategyFactory {
	fmt.Println(fmt.Sprintf("Using random delay strategy betweem %d s and %d s", minDelaySec, maxDelaySec))
	return &randomDelayStrategyFactory{delayMinSec: minDelaySec, delayMaxSec: maxDelaySec}
}

