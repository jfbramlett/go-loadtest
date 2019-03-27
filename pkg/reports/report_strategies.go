package reports

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"time"
)


type ReportStrategy interface {
	Report(concurrentRequests int, testDurationSec time.Duration, results collector.ResultCollector)
}
