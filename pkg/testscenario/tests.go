package testscenario

import (
	"context"
	"time"

	wr "github.com/mroth/weightedrand"
)

type TestFunc func(ctx context.Context, resultsCollector ResultCollector)

type SetupFunc func(ctx context.Context) (context.Context, error)

type TeardownFunc func(ctx context.Context) error

func NewTest(name string, underlying func(ctx context.Context) error) TestFunc {
	return func(ctx context.Context, resultsCollector ResultCollector) {
		start := time.Now()
		err := underlying(ctx)
		if err != nil {
			resultsCollector.AddTestResult(TestFailed(name, err, time.Since(start)))
		}
		resultsCollector.AddTestResult(TestPassed(name, time.Since(start)))
	}
}

type WeightedTest struct {
	Test   TestFunc
	Weight uint
}

func NewWeightedTest(test TestFunc, weight uint) WeightedTest {
	return WeightedTest{Test: test, Weight: weight}
}

func NewWeightedTestFunc(tests ...WeightedTest) TestFunc {

	choices := make([]wr.Choice, len(tests))
	for idx, test := range tests {
		choices[idx] = wr.Choice{Item: test, Weight: test.Weight}
	}

	chooser, _ := wr.NewChooser(
		choices...,
	)

	return func(ctx context.Context, resultsCollector ResultCollector) {
		step := chooser.Pick().(WeightedTest)
		step.Test(ctx, resultsCollector)
	}
}
