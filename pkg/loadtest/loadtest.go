package loadtest

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/reports"
	"github.com/jfbramlett/go-loadtest/pkg/runstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
	"time"
)

type LoadRunner struct {
	TestDurationSec				int64
	ConcurrentRequests			int
	RampUpStrategy				rampstrategy.RampUpStrategy
	RunStrategyFactory			runstrategy.RunStrategyFactory
	TestNamer					naming.TestNamer
	TestCollector				collector.ResultCollector
	ReportingStrategy			reports.ReportStrategy

	Target						utils.RunFunc

	endTime						time.Time
}


func (l LoadRunner) RunLoad() {
	l.endTime = time.Now().Add(time.Second * time.Duration(l.TestDurationSec))

	utils.Log("Preparing tests")
	runners := make([]runstrategy.RunStrategy, 0)
	for i := 0; i < l.ConcurrentRequests; i++ {
		r := l.RampUpStrategy.CreateRunStrategy(l.TestNamer.GetName(i), l.RunStrategyFactory, l.TestCollector, l.Target)
		runners = append(runners, r)
	}

	utils.Log("Starting runners")
	wg := sync.WaitGroup{}
	for _, r := range runners {
		go r.Start(wg)
	}

	utils.Log("Waiting for test time")
	time.Sleep(time.Second * time.Duration(l.TestDurationSec))

	utils.Log("Stopping runners")
	for _, r := range runners {
		r.Stop()
	}
	utils.Log("Waiting for tests to end")
	wg.Wait()


	l.ReportingStrategy.Report(l.ConcurrentRequests, l.TestDurationSec, l.TestCollector)
}


func RunLoad(testDurationSec int64,
	concurrentRequests int,
	rampUpStrategy rampstrategy.RampUpStrategy,
	runStrategy runstrategy.RunStrategyFactory,
	namer naming.TestNamer,
	collector collector.ResultCollector,
	reportStrategy reports.ReportStrategy,
	runFunc utils.RunFunc) {

		loadTester := LoadRunner{
			TestDurationSec: 		testDurationSec,
			ConcurrentRequests: 	concurrentRequests,
			RampUpStrategy: 		rampUpStrategy,
			RunStrategyFactory:   	runStrategy,
			TestNamer: 				namer,
			TestCollector: 			collector,
			ReportingStrategy:  	reportStrategy,
			Target:      			runFunc,
		}

	loadTester.RunLoad()
}