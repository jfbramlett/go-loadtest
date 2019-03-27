package collector


type TestResult struct {
	TestId			string
	Duration		int64
	Error			error
	Passed			bool
}


type ResultCollector interface {
	AddPassedTest(testId string, duration int64)
	AddFailedTest(testId string, duration int64, error error)

	GetPassedTests() []TestResult
	GetFailedTests() []TestResult
}

type inmemoryResultCollector struct {
	passed 			[]TestResult
	failed			[]TestResult
}

func (i *inmemoryResultCollector) AddPassedTest(testId string, duration int64) {
	i.passed = append(i.passed, TestResult{TestId: testId, Duration: duration, Passed: true})
}

func (i *inmemoryResultCollector) AddFailedTest(testId string, duration int64, error error) {
	i.failed = append(i.failed, TestResult{TestId: testId, Duration: duration, Error: error, Passed: false})
}

func (i *inmemoryResultCollector) GetPassedTests() []TestResult {
	return i.passed
}

func (i *inmemoryResultCollector) GetFailedTests() []TestResult {
	return i.failed
}


func NewInMemoryRunCollector() ResultCollector {
	return &inmemoryResultCollector{passed: make([]TestResult, 0), failed: make([]TestResult, 0)}
}
