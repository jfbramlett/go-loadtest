package testscenario

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/metrics"
	"time"
)

func TestPassed(testId string, name string, requestId string) TestResult {
	return &simpleTestResult{id: testId, name: name, requestId: requestId, metrics: []metrics.Metric{}, passed: true}
}

func TestPassedWithMetrics(testId string, name string, requestId string, metrics []metrics.Metric) TestResult {
	return &simpleTestResult{id: testId, name: name, requestId: requestId, metrics: metrics, passed: true}
}

func TestFailed(testId string, name string, requestId string, err error) TestResult {
	return &simpleTestResult{id: testId, name: name, requestId: requestId, error: err, metrics: []metrics.Metric{}, passed: false}
}

func TestFailedWithMetrics(testId string, name string, requestId string, err error, metrics []metrics.Metric) TestResult {
	return &simpleTestResult{id: testId, name: name, requestId: requestId, error: err, metrics: metrics, passed: false}
}


type TestResult interface {
	Id() string
	Name() string
	RequestId() string
	Error() error
	Metrics() []metrics.Metric
	Duration() time.Duration
	Passed() bool
	Failed() bool
	AddMetric(metric metrics.Metric)
	SetDuration(dur time.Duration)
}

type simpleTestResult struct {
	id			string
	name		string
	requestId	string
	error		error
	metrics 	[]metrics.Metric
	duration	time.Duration
	passed		bool
}

func (s *simpleTestResult) Id() string {
	return s.id
}

func (s *simpleTestResult) Name() string {
	return s.name
}

func (s *simpleTestResult) RequestId() string {
	return s.requestId
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

type Test interface {
	Run(ctx context.Context, testId string) TestResult
}
