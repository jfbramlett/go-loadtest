package loadprofile

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
)


type Load interface {
	Run()
}

type defaultLoad struct {
	ctx				context.Context
	runSteps		[]Step
}


// runs the loop that executes our run steps around running the test
func (r *defaultLoad) Run() {
	utils.Logt(utils.GetTestId(r.ctx), "starting run")
	for _, step := range r.runSteps {
		err := step.Execute(r.ctx)
		if err != nil {
			return
		}
	}
	utils.Logt(utils.GetTestId(r.ctx), "run complete")
}

func NewLoad(ctx context.Context, runSteps []Step) Load {
	return &defaultLoad{ctx: ctx, runSteps: runSteps}
}