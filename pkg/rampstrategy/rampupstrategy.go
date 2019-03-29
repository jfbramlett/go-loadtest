package rampstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/runstrategy"
)

type RampUpStrategy interface {
	CreateRunStrategy(testId string, factory runstrategy.RunStrategyFactory, collector collector.ResultCollector, runFunc testwrapper.RunFunc) runstrategy.RunStrategy
}