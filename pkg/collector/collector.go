package collector

import "time"

type TestResult struct {
	TestId			string
	Duration		time.Duration
	Error			error
	Passed			bool
}


func NewPassedTest(testId string, dur time.Duration) TestResult {
	return TestResult{TestId: testId, Duration: dur, Passed: true}
}


func NewFailedTest(testId string, dur time.Duration, err error) TestResult {
	return TestResult{TestId: testId, Duration: dur, Error: err, Passed: false}
}


type ResultCollector interface {
	AddTestResult(result TestResult)

	Start()
	Stop()

	GetPassedTests() []TestResult
	GetFailedTests() []TestResult
}

type inmemoryResultCollector struct {
	passed 				[]TestResult
	failed				[]TestResult
	collector			chan TestResult
	collectionComplete 	bool
}


func (i *inmemoryResultCollector) AddTestResult(result TestResult) {
	i.collector <- result
}

func (i *inmemoryResultCollector) GetPassedTests() []TestResult {
	return i.passed
}

func (i *inmemoryResultCollector) GetFailedTests() []TestResult {
	return i.failed
}

func (i *inmemoryResultCollector) Start() {
	go func() {
		for t := range i.collector {
			if t.Passed {
				i.passed = append(i.passed, t)
			} else {
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
	collector :=  &inmemoryResultCollector{passed: make([]TestResult, 0), failed: make([]TestResult, 0), collector: make(chan TestResult)}
	return collector
}
