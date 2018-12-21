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

	outchan := make(chan reports.RunTimes)

	for i := 0; i < l.ConcurrentRequests; i++ {
		go l.runFunc(outchan)
	}

	results := make([]reports.RunTimes, 0)

	for ; len(results) < l.ConcurrentRequests; {
		results = append(results, <-outchan)
	}

	l.ReportingStrategy.Report(l.ConcurrentRequests, l.TestDurationSec, results)
}


func (l LoadRunner) runFunc(outchan chan<- reports.RunTimes) {
	rt := reports.NewRunTimes()
	for  start := time.Now(); start.Before(l.endTime); start = time.Now() {
		_, err := l.Target.Run()
		end := time.Now()
		if err == nil {
			fmt.Println(time.Since(start))
			rt.Times = append(rt.Times, utils.TimeDiffMillis(start, end))
		} else {
			fmt.Printf("Error: %s\n", time.Since(start))
			rt.Errors = append(rt.Errors, utils.TimeDiffMillis(start, end))
		}
		l.DelayingStrategy.Wait()
	}

	outchan <- rt
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