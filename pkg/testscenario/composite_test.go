package testscenario

import (
	"context"
	"time"
)

func NewCompositeTest(delayBetween time.Duration, tests ...TestFunc) TestFunc {
	return func(ctx context.Context, resultsCollector ResultCollector) {
		for _, t := range tests {
			t(ctx, resultsCollector)
			time.Sleep(delayBetween)
		}
	}
}
