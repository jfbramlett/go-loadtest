package runstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
)


type RunStrategy interface {
	Start(wg sync.WaitGroup)
	Stop()
	GetResults() collector.ResultCollector
}

type RunStrategyFactory interface {
	GetRunStrategy(testId string, initialDelay int, funcToRun utils.RunFunc, resultCollector collector.ResultCollector) RunStrategy
}

