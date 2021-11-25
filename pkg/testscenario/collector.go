package testscenario

type ResultCollector interface {
	AddTestResult(result TestResult)

	Start()
	Stop()

	GetTests() []string

	GetPassedTests(name string) []TestResult
	GetFailedTests(name string) []TestResult
}
