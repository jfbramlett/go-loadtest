package runstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
)


type RunStrategy interface {
	Start(wg sync.WaitGroup)
	Stop()
	GetResults() utils.ResultCollector
}

type RunStrategyFactory interface {
	GetRunStrategy(runId string, initialDelay int, funcToRun utils.RunFunc, resultCollector utils.ResultCollector) RunStrategy
}

