package loadtest

import (
	"fmt"
	"github.com/jfbramlett/go-loadtest/pkg/delays"
	"github.com/jfbramlett/go-loadtest/pkg/reports"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type RunStrategy interface {
	Run() (interface{}, error)
}

type LoadRunner struct {
	TestDurationSec				int64
	ConcurrentRequests			int
	DelayingStrategy			delays.DelayStrategy
	ReportingStrategy			reports.ReportStrategy

	Target						RunStrategy

	endTime						time.Time
}


func (l LoadRunner) RunLoad() {
	l.endTime = time.Now().Add(time.Second * time.Duration(l.TestDurationSec))

	runners := make([]*runner, 0)
	for i := 0; i < l.ConcurrentRequests; i++ {
		ticker := l.DelayingStrategy.GetTicker()

		r := &runner{Target: l.Target, Ticker: ticker, Results: utils.NewRunTimes()}
		runners = append(runners, r)
		go r.runFunc()
	}


	time.Sleep(time.Second * time.Duration(l.TestDurationSec))

	for _, r := range runners {
		r.Ticker.Stop()
	}

	results := make([]*utils.RunTimes, 0)

	for _, r := range runners {
		results = append(results, r.Results)
	}

	l.ReportingStrategy.Report(l.ConcurrentRequests, l.TestDurationSec, results)
}


type runner struct {
	Target			RunStrategy
	Results			*utils.RunTimes
	Ticker			*time.Ticker
}


func (r runner) runFunc() {
	for start := range r.Ticker.C {
		_, err := r.Target.Run()
		end := time.Now()
		if err == nil {
			//fmt.Println(time.Since(start))
			r.Results.Times = append(r.Results.Times, utils.TimeDiffMillis(start, end))
		} else {
			fmt.Printf("Error: %s\n", time.Since(start))
			r.Results.Errors = append(r.Results.Errors, utils.TimeDiffMillis(start, end))
		}
	}
}


func RunLoad(testDurationSec int64, concurrentRequests int,
	delayStrategy delays.DelayStrategy, reportStrategy reports.ReportStrategy,
	runStrategy RunStrategy) {

		loadTester := LoadRunner{TestDurationSec: testDurationSec,
		ConcurrentRequests: concurrentRequests,
		DelayingStrategy:   delayStrategy,
		ReportingStrategy:  reportStrategy,
		Target:      		runStrategy,
	}

	loadTester.RunLoad()
}