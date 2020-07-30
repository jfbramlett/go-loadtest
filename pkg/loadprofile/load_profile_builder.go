package loadprofile

import (
	"context"
	"github.com/ninthwave/nwp-load-test/pkg/collector"
	"github.com/ninthwave/nwp-load-test/pkg/rampstrategy"
	"github.com/ninthwave/nwp-load-test/pkg/testscenario"
	"time"
)

type LoadProfileType int

const (
	StaticProfile LoadProfileType = 1
	RandomProfile LoadProfileType = 2
	PartialRandomProfile LoadProfileType = 3
)


func GetLoadProfileType(v int) LoadProfileType {
	switch v {
	case 1:
		return StaticProfile
	case 2:
		return RandomProfile
	case 3:
		return PartialRandomProfile
	}

	return StaticProfile
}

type LoadProfileBuilder interface {
	GetLoadProfiles(ctx context.Context, runFunc testscenario.Test, resultCollector collector.ResultCollector) []LoadProfile
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