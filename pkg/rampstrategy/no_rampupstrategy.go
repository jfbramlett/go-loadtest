package rampstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/runstrategy"
)

type noRampUpStrategy struct {
	interval				int
	currentDelay			int
	usersPerInterval		int
	currentIntervalUsers	int
}


func (s *noRampUpStrategy) CreateRunStrategy(testId string, factory runstrategy.RunStrategyFactory, collector collector.ResultCollector, runFunc testwrapper.RunFunc) runstrategy.RunStrategy {
	return factory.GetRunStrategy(testId, 0, runFunc, collector)
}


func NewNoRampUpStrategy() RampUpStrategy {
	return &noRampUpStrategy{}
}