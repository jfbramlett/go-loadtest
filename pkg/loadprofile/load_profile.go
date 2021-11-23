package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
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


func (r *defaultLoadProfile) Run(ctx context.Context) {
	logger := utils.LoggerFromContext(ctx)
	logger.Info("starting run")

	testStart := time.Now()
	for r.testLength > time.Since(testStart) {
		for _, step := range r.runSteps {
			if r.testLength > time.Since(testStart) {
				err := step.Execute(ctx)
				if r.stopOnError && err != nil {
					logger.WithError(err).Error("run failed with error")
					return
				}
			}
		}
	}
	logger.Info("run complete")
}

func NewLoadProfile(runSteps []steps.Step, testLength time.Duration, stopOnError bool) LoadProfile {
	return &defaultLoadProfile{runSteps: runSteps, testLength: testLength, stopOnError: stopOnError}
}