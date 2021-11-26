package testscenario

import (
	"context"
	"time"

	"github.com/jfbramlett/go-loadtest/pkg/utils"
)

type PauseStrategyFunc func(ctx context.Context)

func FixedPauseStrategy(dur time.Duration) PauseStrategyFunc {
	return func(ctx context.Context) {
		time.Sleep(dur)
	}
}

func RandomPauseStrategy(min time.Duration, max time.Duration) PauseStrategyFunc {
	return func(ctx context.Context) {
		time.Sleep(utils.RandomMilliSecondDuration(min, max))
	}
}
