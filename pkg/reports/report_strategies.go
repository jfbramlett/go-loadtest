package reports

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
)


type ReportStrategy interface {
	Report(concurrentRequests int, testDurationSec int64, results collector.ResultCollector)
}
