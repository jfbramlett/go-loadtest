package testscenario

import (
	"context"
	"time"
)

type TestFunc func(ctx context.Context, resultsCollector ResultCollector)

type SetupFunc func(ctx context.Context) (context.Context, error)

type TeardownFunc func(ctx context.Context) error

func NewTest(name string, underlying func(ctx context.Context) error) TestFunc {
	return func(ctx context.Context, resultsCollector ResultCollector) {
		start := time.Now()
		err := underlying(ctx)
		if err != nil {
			resultsCollector.AddTestResult(TestFailed(name, err, time.Since(start)))
		}
		resultsCollector.AddTestResult(TestPassed(name, time.Since(start)))
	}
}
