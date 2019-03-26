package reports

import (
	"github.com/jfbramlett/go-loadtest/pkg/utils"
)


type ReportStrategy interface {
	Report(concurrentRequests int, testDurationSec int64, results []utils.ResultCollector)
}
