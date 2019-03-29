package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
)

type LoadProfileBuilder interface {
	GetLoadProfiles(runFunc utils.RunFunc, resultCollector collector.ResultCollector) []LoadProfile
}
