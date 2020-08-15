package collector

import (
    "github.com/jfbramlett/go-loadtest/pkg/testscenario"
    "time"
)

type inmemoryResultCollector struct {
    passed 				[]testscenario.TestResult
    failed				[]testscenario.TestResult
    collector			chan testscenario.TestResult
    collectionComplete 	bool
    MinLatencyThreshold	time.Duration
    MaxLatencyThreshold	time.Duration
}


func (i *inmemoryResultCollector) AddTestResult(result testscenario.TestResult) {
    i.collector <- result
}

func (i *inmemoryResultCollector) GetPassedTests() []testscenario.TestResult {
    return i.passed
}

func (i *inmemoryResultCollector) GetFailedTests() []testscenario.TestResult {
    return i.failed
}

func (i *inmemoryResultCollector) Start() {
    go func() {
        //opsPassed := metrics.NewCounter("perf_tests_passed", "The total number of tests that passed")
        //opsFailed := metrics.NewCounter("perf_tests_failed", "The total number of tests that passed")
        //latency := metrics.NewSummaryVec("perf_test_latency", "The latency for the tests", i.MinLatencyThreshold, i.MaxLatencyThreshold)

        for t := range i.collector {
            //latency.WithLabelValues(t.TestName).Observe(float64(t.Duration.Nanoseconds()))
            if t.Passed() {
                //opsPassed.WithLabelValues(t.TestName).Inc()
                i.passed = append(i.passed, t)
            } else {
                //opsFailed.WithLabelValues(t.TestName).Inc()
                i.failed = append(i.failed, t)
            }
        }
        i.collectionComplete = true
    }()
}

func (i *inmemoryResultCollector) Stop() {
    close(i.collector)
    for !i.collectionComplete {
        time.Sleep(time.Duration(1) * time.Second)
    }
}

func NewInMemoryRunCollector() ResultCollector {
    collector :=  &inmemoryResultCollector{passed: make([]testscenario.TestResult, 0),
        failed: make([]testscenario.TestResult, 0), collector: make(chan testscenario.TestResult)}
    return collector
}
