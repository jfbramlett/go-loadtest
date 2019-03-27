package rampstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/runstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
)

type RampUpStrategy interface {
	CreateRunStrategy(testId string, factory runstrategy.RunStrategyFactory, collector collector.ResultCollector, runFunc utils.RunFunc) runstrategy.RunStrategy
}