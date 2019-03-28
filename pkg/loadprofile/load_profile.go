package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
)

type LoadProfile interface {
	GetLoad(namer naming.TestNamer, runFunc utils.RunFunc, resultCollector collector.ResultCollector) []Load
}
