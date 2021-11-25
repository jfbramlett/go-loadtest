package testscenario

import (
	"time"

	"github.com/jfbramlett/go-loadtest/pkg/metrics"
)

func TestPassed(name string, dur time.Duration, metrics ...metrics.Metric) TestResult {
	return &simpleTestResult{name: name, metrics: metrics, passed: true, duration: dur}
}

func TestFailed(name string, err error, dur time.Duration, metrics ...metrics.Metric) TestResult {
	return &simpleTestResult{name: name, error: err, metrics: metrics, passed: false, duration: dur}
}

type TestResult interface {
	Name() string
	Error() error
	Metrics() []metrics.Metric
	Duration() time.Duration
	Passed() bool
	Failed() bool
}

type simpleTestResult struct {
	name     string
	error    error
	metrics  []metrics.Metric
	duration time.Duration
	passed   bool
}

func (s *simpleTestResult) Name() string {
	return s.name
}

func (s *simpleTestResult) Error() error {
	return s.error
}

func (s *simpleTestResult) Metrics() []metrics.Metric {
	return s.metrics
}

func (s *simpleTestResult) Duration() time.Duration {
	return s.duration
}

func (s *simpleTestResult) Failed() bool {
	return !s.passed
}

func (s *simpleTestResult) Passed() bool {
	return s.passed
}

func (s *simpleTestResult) AddMetric(metric metrics.Metric) {
	s.metrics = append(s.metrics, metric)
}

func (s *simpleTestResult) SetDuration(dur time.Duration) {
	s.duration = dur
}
