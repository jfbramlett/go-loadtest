package runstrategy

import (
	"fmt"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

// simple utility used to run a test
func runTest(funcToRun utils.RunFunc, runCollector utils.ResultCollector) {
	start := time.Now()
	_, err := funcToRun()
	end := time.Now()
	if err == nil {
		runCollector.AddRuntime(utils.TimeDiffMillis(start, end))
	} else {
		fmt.Printf("Error: %s\n", time.Since(start))
		runCollector.AddErrorRuntime(utils.TimeDiffMillis(start, end))
	}
}
