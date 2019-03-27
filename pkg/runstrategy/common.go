package runstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

// simple utility used to run a test
func runTest(testId string, funcToRun utils.RunFunc, runCollector collector.ResultCollector) {
	start := time.Now()
	_, err := funcToRun()
	end := time.Now()
	if err == nil {
		runCollector.AddPassedTest(testId, utils.TimeDiffMillis(start, end))
	} else {
		utils.Logtf(testId, "Error - %s\n", time.Since(start))
		runCollector.AddFailedTest(testId, utils.TimeDiffMillis(start, end), err)
	}
}
