package steps

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

// Step that executes our run function
type runFuncStep struct {
	runFunc				utils.RunFunc
	resultCollector 	collector.ResultCollector
}

func (r *runFuncStep) Execute(ctx context.Context) error {
	timerStart := time.Now()
	err := r.runFunc(ctx)

	if err == nil {
		r.resultCollector.AddTestResult(collector.NewPassedTest(utils.GetTestId(ctx), time.Since(timerStart)))
	} else {
		utils.Logtf(utils.GetTestId(ctx), "Error - %s\n", err)
		r.resultCollector.AddTestResult(collector.NewFailedTest(utils.GetTestId(ctx), time.Since(timerStart), err))
	}

	return nil
}

func NewRunFuncStep(runFunc utils.RunFunc, resultCollector collector.ResultCollector) Step {
	return &runFuncStep{runFunc: runFunc, resultCollector: resultCollector}
}


