package steps

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/logging"
	"github.com/jfbramlett/go-loadtest/pkg/testscenario"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

// Step that executes our run function
type runFuncStep struct {
	test            testscenario.Test
	resultCollector collector.ResultCollector
}

func (r *runFuncStep) Execute(ctx context.Context) error {
	logger, ctx := logging.GetLoggerFromContext(ctx, r)
	testId := utils.GetTestIdFromContext(ctx)
	timerStart := time.Now()
	result := r.test.Run(ctx, testId)

	runTime := time.Since(timerStart)
	result.SetDuration(runTime)

	if result.Passed() {
		r.resultCollector.AddTestResult(result)
	} else {
		logger.Error(ctx, result.Error(), "Error in test " + result.Name() + " with id " + result.Id() + " and request id " + result.RequestId())
		r.resultCollector.AddTestResult(result)
	}

	return nil
}

func NewRunFuncStep(test testscenario.Test, resultCollector collector.ResultCollector) Step {
	return &runFuncStep{test: test, resultCollector: resultCollector}
}


