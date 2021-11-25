package reports

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/testscenario"
	"time"
)


type ReportStrategy interface {
	Report(ctx context.Context, concurrentRequests int, testDurationSec time.Duration, results testscenario.ResultCollector)
}
