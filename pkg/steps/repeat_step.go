package steps

import (
	"context"
	"time"
)

// Step that repeats another step a given number of times
type repeatStep struct {
	timesToRepeat 	int
	stepToRepeat	Step
	delayBetween	time.Duration
	stopOnError		bool
}

func (r *repeatStep) Execute(ctx context.Context) error {
	for i := 0; i < r.timesToRepeat; i++ {
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

func NewRepeatStep(timesToRepeat int, stepToRepease Step, delayBetween time.Duration, stopOnError bool) Step {
	return &repeatStep{timesToRepeat: timesToRepeat, stepToRepeat: stepToRepease, delayBetween: delayBetween, stopOnError: stopOnError}
}


