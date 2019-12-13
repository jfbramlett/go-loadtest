package collector

import (
	"time"
)

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