package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)

type RunStep struct {
	startDelay		time.Duration
	runTime			time.Duration
	invocationDelay	time.Duration
}


type RunProfile struct {
	steps	[]RunStep
}

type LoadProfile interface {
	GetRunners(namer naming.TestNamer, runFunc utils.RunFunc, resultCollector collector.ResultCollector) []Runner
}
