package testscenario

import (
	"context"

	wr "github.com/mroth/weightedrand"
)

type WeightedTest struct {
	Test   TestFunc
	Weight uint
}

func NewWightedTest(test TestFunc, weight uint) WeightedTest {
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
