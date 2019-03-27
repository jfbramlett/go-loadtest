package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)


type Runner interface {
	Run()
}

type defaultRunner struct {
	testId			string
	runSteps		[]RunStep
	runFunc			utils.RunFunc
	resultCollector	collector.ResultCollector
}


// runs the loop that executes our run steps around running the test
func (r *defaultRunner) Run() {
	for _, step := range r.runSteps {
		if step.startDelay > 0 {
			utils.Logt(r.testId, "pausing for initial delay")
			time.Sleep(step.startDelay)
		}

		utils.Logt(r.testId, "starting test")
		stepStart := time.Now()
		for step.runTime > time.Since(stepStart) {
			funcStart := time.Now()
			_, err := r.runFunc()
			if err == nil {
				r.resultCollector.AddTestResult(collector.NewPassedTest(r.testId, time.Since(funcStart)))
			} else {
				utils.Logtf(r.testId, "Error - %s\n", time.Since(funcStart))
				r.resultCollector.AddTestResult(collector.NewFailedTest(r.testId, time.Since(funcStart), err))
			}

			if step.invocationDelay > 0 {
				time.Sleep(step.invocationDelay)
			}
		}
	}
	utils.Logt(r.testId, "test complete")
}

func NewRunner(testId string, runSteps []RunStep, runFunc utils.RunFunc, resultCollector collector.ResultCollector) Runner {
	return &defaultRunner{testId: testId, runSteps: runSteps, runFunc: runFunc, resultCollector: resultCollector}
}