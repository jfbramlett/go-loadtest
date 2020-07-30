package steps

import (
	"context"
	"github.com/ninthwave/nwp-load-test/pkg/utils"
	"time"
)

// Step that pauses execution for a given duration
type waitStep struct {
	waitTime	time.Duration
	repeat		bool
	run			bool
}

func (w *waitStep) Execute(ctx context.Context) error {
	if w.run {
		time.Sleep(w.waitTime)
		w.run = w.repeat
	}
	return nil
}

func NewWaitStep(waitTime time.Duration) Step {
	return &waitStep{waitTime: waitTime, repeat: true, run: true}
}

func NewInitialWaitStep(waitTime time.Duration) Step {
	return &waitStep{waitTime: waitTime, repeat: false, run: true}
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


