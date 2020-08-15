package collector

import (
	"github.com/jfbramlett/go-loadtest/pkg/testscenario"
)

type ResultCollector interface {
	AddTestResult(result testscenario.TestResult)

	Start()
	Stop()

	GetPassedTests() []testscenario.TestResult
	GetFailedTests() []testscenario.TestResult
}