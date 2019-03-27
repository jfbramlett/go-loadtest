package rampstrategy

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/runstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
)

type smoothRampUpStrategy struct {
	interval				int
	currentDelay			int
	usersPerInterval		int
	currentIntervalUsers	int
}


func (s *smoothRampUpStrategy) CreateRunStrategy(testId string, factory runstrategy.RunStrategyFactory, collector collector.ResultCollector, runFunc utils.RunFunc) runstrategy.RunStrategy {
	rs := factory.GetRunStrategy(testId, s.currentDelay, runFunc, collector)
	s.currentIntervalUsers++
	if s.currentIntervalUsers == s.usersPerInterval {
		s.currentDelay = s.currentDelay + s.interval
		s.currentIntervalUsers = 0
	}

	return rs
}



func NewSmoothRampUpStrategy(rampUpPeriod int, maxUsers int, minDelay int) RampUpStrategy {
	maxRamps := rampUpPeriod / minDelay
	minRamps := maxRamps/2

	ramps := utils.RandomIntBetween(minRamps, maxRamps)
	rampInterval := rampUpPeriod / ramps

	usersPerInterval := maxUsers / ramps

	return &smoothRampUpStrategy{
		interval: rampInterval,
		usersPerInterval: usersPerInterval}
}