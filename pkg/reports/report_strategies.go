package reports

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"time"
)


type ReportStrategy interface {
	Report(ctx context.Context, concurrentRequests int, testDurationSec time.Duration, results collector.ResultCollector)
}
