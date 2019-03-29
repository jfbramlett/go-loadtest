package steps

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

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


