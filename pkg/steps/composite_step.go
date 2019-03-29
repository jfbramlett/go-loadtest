package steps

import (
	"context"
	"time"
)

// Step that executes a set of steps in sequence
type compositeStep struct {
	steps			[]Step
	delayBetween	time.Duration
	stopOnError 	bool
}

func (r *compositeStep) Execute(ctx context.Context) error {
	for _, step := range r.steps {
		err := step.Execute(ctx)
		if r.stopOnError && err != nil {
			return err
		}

		if r.delayBetween > 0 {
			time.Sleep(r.delayBetween)
		}
	}

	return nil
}

func NewCompositeStep(delayBetween time.Duration, stopOnError bool, steps ...Step) Step {
	return &compositeStep{delayBetween: delayBetween, stopOnError: stopOnError, steps: steps}
}

