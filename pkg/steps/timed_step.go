package steps

import (
	"context"
	"time"
)

// Step that runs a step for a given duration
type runForStep struct {
	timeToRun		time.Duration
	stepToRepeat	Step
	delayBetween	time.Duration
	stopOnError		bool
}

func (r *runForStep) Execute(ctx context.Context) error {
	stepStart := time.Now()
	for r.timeToRun > time.Since(stepStart) {
		err := r.stepToRepeat.Execute(ctx)
		if r.stopOnError && err != nil {
			return err
		}

		if r.delayBetween > 0 {
			time.Sleep(r.delayBetween)
		}
	}

	return nil
}

func NewRunForStep(timeToRun time.Duration, stepToRepeat Step, delayBetween time.Duration, stopOnError bool) Step {
	return &runForStep{timeToRun: timeToRun, stepToRepeat: stepToRepeat, delayBetween: delayBetween, stopOnError: stopOnError}
}

