package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type Step interface {
	Execute(ctx context.Context) error
}


// Step that pauses execution for a given duration
type waitStep struct {
	waitTime	time.Duration
}

func (w *waitStep) Execute(ctx context.Context) error {
	time.Sleep(w.waitTime)
	return nil
}

func NewWaitStep(waitTime time.Duration) Step {
	return &waitStep{waitTime: waitTime}
}


// step that does a random wait
type randomWaitStep struct {
	minDuration		time.Duration
	maxDuration		time.Duration
	durationGen		utils.DurationGenator
}

func (w *randomWaitStep) Execute(ctx context.Context) error {
	time.Sleep(w.durationGen(w.minDuration, w.maxDuration))
	return nil
}

func NewRandomWaitStep(minDuration time.Duration, maxDuration time.Duration, durationGenerator utils.DurationGenator) Step {
	return &randomWaitStep{minDuration: minDuration, maxDuration: maxDuration, durationGen: durationGenerator}
}




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
