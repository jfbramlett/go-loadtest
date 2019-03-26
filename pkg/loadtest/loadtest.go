package loadtest

import (
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/reports"
	"github.com/jfbramlett/go-loadtest/pkg/runstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
	"time"
)

type LoadRunner struct {
	TestDurationSec				int64
	ConcurrentRequests			int
	RunStrategyFactory			runstrategy.RunStrategyFactory
	TestNamer					naming.TestNamer
	ReportingStrategy			reports.ReportStrategy

	Target						utils.RunFunc

	endTime						time.Time
}


func (l LoadRunner) RunLoad() {
	l.endTime = time.Now().Add(time.Second * time.Duration(l.TestDurationSec))

	runners := make([]runstrategy.RunStrategy, 0)
	wg := sync.WaitGroup{}
	for i := 0; i < l.ConcurrentRequests; i++ {
		r := l.RunStrategyFactory.GetRunStrategy(l.TestNamer.GetName(i), 0, l.Target, utils.NewInMemoryRunCollector())
		runners = append(runners, r)
		go r.Start(wg)
	}

	time.Sleep(time.Second * time.Duration(l.TestDurationSec))

	for _, r := range runners {
		r.Stop()
	}
	wg.Wait()

	results := make([]utils.ResultCollector, 0)

	for _, r := range runners {
		results = append(results, r.GetResults())
	}

	l.ReportingStrategy.Report(l.ConcurrentRequests, l.TestDurationSec, results)
}


func RunLoad(testDurationSec int64,
	concurrentRequests int,
	runStrategy runstrategy.RunStrategyFactory,
	namer naming.TestNamer,
	reportStrategy reports.ReportStrategy,
	runFunc utils.RunFunc) {

		loadTester := LoadRunner{TestDurationSec: testDurationSec,
		ConcurrentRequests: concurrentRequests,
		RunStrategyFactory:   runStrategy,
		TestNamer: namer,
		ReportingStrategy:  reportStrategy,
		Target:      		runFunc,
	}

	loadTester.RunLoad()
}