package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"time"
)

type LoadProfileType int

const (
	StaticProfile LoadProfileType = 1
	RandomProfile LoadProfileType = 2
	PartialRandomProfile LoadProfileType = 3
)


type LoadProfileBuilder interface {
	GetLoadProfiles(runFunc testwrapper.Test, resultCollector collector.ResultCollector) []LoadProfile
}


func NewLoadProfileBuilder(profileType LoadProfileType, users int, testLength time.Duration, interval time.Duration, rampStrategy rampstrategy.RampStrategyType) LoadProfileBuilder {
	ramping := rampstrategy.NewRampStrategy(rampStrategy)

	switch profileType {
	case StaticProfile:
		return NewStaticProfileBuilder(users, testLength, interval, ramping)
	case RandomProfile:
		return NewRandomProfileBuilder(users, testLength, interval, ramping)
	case PartialRandomProfile:
		return NewPartialRandomProfileBuilder(users, testLength, interval, ramping)
	}

	return nil
}