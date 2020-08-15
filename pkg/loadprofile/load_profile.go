package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/logging"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"time"
)


type LoadProfile interface {
	Run(ctx context.Context)
}

type defaultLoadProfile struct {
	runSteps		[]steps.Step
	testLength		time.Duration
	stopOnError		bool
}


// runs the loop that executes our run steps around running the test
func (r *defaultLoadProfile) Run(ctx context.Context) {
	logger, ctx := logging.GetLoggerFromContext(ctx, r)
	logger.Info(ctx, "starting run")

	testStart := time.Now()
	for r.testLength > time.Since(testStart) {
		for _, step := range r.runSteps {
			if r.testLength > time.Since(testStart) {
				err := step.Execute(ctx)
				if r.stopOnError && err != nil {
					logger.Error(ctx, err, "run failed with error")
					return
				}
			}
		}
	}
	logger.Info(ctx, "run complete")
}

func NewLoadProfile(runSteps []steps.Step, testLength time.Duration, stopOnError bool) LoadProfile {
	return &defaultLoadProfile{runSteps: runSteps, testLength: testLength, stopOnError: stopOnError}
}