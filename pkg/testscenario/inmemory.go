package testscenario

import (
	"time"
)

var empty = struct{}{}

type inmemoryResultCollector struct {
	passed              []TestResult
	failed              []TestResult
	collector           chan TestResult
	collectionComplete  bool
	MinLatencyThreshold time.Duration
	MaxLatencyThreshold time.Duration
}

func (i *inmemoryResultCollector) AddTestResult(result TestResult) {
	i.collector <- result
}

func (i *inmemoryResultCollector) GetPassedTests(name string) []TestResult {
	results := make([]TestResult, 0)
	for _, t := range i.passed {
		if t.Name() == name {
			results = append(results, t)
		}
	}
	return results
}

func (i *inmemoryResultCollector) GetFailedTests(name string) []TestResult {
	results := make([]TestResult, 0)
	for _, t := range i.failed {
		if t.Name() == name {
			results = append(results, t)
		}
	}
	return results
}

func (i *inmemoryResultCollector) GetTests() []string {
	names := make(map[string]struct{})
	for _, t := range i.passed {
		names[t.Name()] = empty
	}
	for _, t := range i.failed {
		names[t.Name()] = empty
	}

	results := make([]string, 0, len(names))
	for nm := range names {
		results = append(results, nm)
	}
	return results
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
	collector := &inmemoryResultCollector{passed: make([]TestResult, 0),
		failed: make([]TestResult, 0), collector: make(chan TestResult)}
	return collector
}
