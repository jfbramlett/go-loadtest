package collector

import (
	"github.com/ninthwave/nwp-load-test/pkg/testscenario"
)

type ResultCollector interface {
	AddTestResult(result testscenario.TestResult)

	Start()
	Stop()

	GetPassedTests() []testscenario.TestResult
	GetFailedTests() []testscenario.TestResult
}