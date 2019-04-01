package steps

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/logging"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

// Step that executes our run function
type runFuncStep struct {
	test            testwrapper.Test
	resultCollector collector.ResultCollector
}

func (r *runFuncStep) Execute(ctx context.Context) error {
	logger, ctx := logging.GetLoggerFromContext(ctx, r)
	timerStart := time.Now()
	err := r.test.Run(ctx)

	if err == nil {
		r.resultCollector.AddTestResult(collector.NewPassedTest(utils.GetTestIdFromContext(ctx), time.Since(timerStart)))
	} else {
		logger.Error(ctx, err, "Error")
		r.resultCollector.AddTestResult(collector.NewFailedTest(utils.GetTestIdFromContext(ctx), time.Since(timerStart), err))
	}

	return nil
}

func NewRunFuncStep(test testwrapper.Test, resultCollector collector.ResultCollector) Step {
	return &runFuncStep{test: test, resultCollector: resultCollector}
}


