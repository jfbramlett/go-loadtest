package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
)

type LoadProfileBuilder interface {
	GetLoadProfiles(runFunc testwrapper.Test, resultCollector collector.ResultCollector) []LoadProfile
}
