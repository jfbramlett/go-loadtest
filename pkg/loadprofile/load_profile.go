package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/steps"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
)


type LoadProfile interface {
	Run(ctx context.Context)
}

type defaultLoadProfile struct {
	runSteps		[]steps.Step
}


// runs the loop that executes our run steps around running the test
func (r *defaultLoadProfile) Run(ctx context.Context) {
	utils.Logt(utils.GetTestId(ctx), "starting run")
	for _, step := range r.runSteps {
		err := step.Execute(ctx)
		if err != nil {
			return
		}
	}
	utils.Logt(utils.GetTestId(ctx), "run complete")
}

func NewLoadProfile(runSteps []steps.Step) LoadProfile {
	return &defaultLoadProfile{runSteps: runSteps}
}