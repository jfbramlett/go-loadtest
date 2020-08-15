package reports

import (
	"context"
	"fmt"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/logging"
	"time"
)


func NewConsoleReportStrategy(minTimeThreshold, maxTimeThreshold time.Duration) ReportStrategy {
	return &consoleReportStrategy{MinTimeThreshold: minTimeThreshold, MaxTimeThreshold: maxTimeThreshold}
}

type consoleReportStrategy struct {
	MinTimeThreshold			time.Duration
	MaxTimeThreshold			time.Duration
}

func (c *consoleReportStrategy) Report(ctx context.Context, concurrentRequests int, testDurationSec time.Duration, results collector.ResultCollector) {
	var totalRequests int64
	var totalTime time.Duration
	var maxTime time.Duration
	var minTime time.Duration
	var errors int
	var totalAboveThreshold int64
	var totalLessThanMin int64
	var totalInMiddle int64

	for _, t := range results.GetPassedTests() {
		totalRequests++
		totalTime = totalTime + t.Duration()
		if t.Duration() > maxTime {
			maxTime = t.Duration()
		}
		if minTime == 0 || t.Duration() < minTime {
			minTime = t.Duration()
		}
		if t.Duration() <= c.MinTimeThreshold {
			totalLessThanMin++
		} else if t.Duration() > c.MinTimeThreshold && t.Duration() <= c.MaxTimeThreshold {
			totalInMiddle++
		}
		if t.Duration() > c.MaxTimeThreshold {
			totalAboveThreshold++
		}
	}
	errors = len(results.GetFailedTests())

	avgRequestTime := int64(totalTime/time.Millisecond) / totalRequests
	lessThanPercent := (float64(totalLessThanMin) / float64(totalRequests)) * float64(100)
	middlePercent := (float64(totalInMiddle) / float64(totalRequests)) * float64(100)
	thresholdPercent := (float64(totalAboveThreshold)/ float64(totalRequests)) * float64(100)

	logger, ctx := logging.GetLoggerFromContext(ctx, c)
	msg := fmt.Sprintf(`Total Concurrent Requests: %d
Test Duration %d s
Total Time %d ms
Total Requests %d
Avg Time %d ms
Max Time %d ms
Min Time %d ms
Num Errors %d
Num Below %d ms %d (%.2f %%)
Num Between %d ms and %d ms %d (%.2f %%)
Num Above %d ms %d (%.2f %%)`,
		concurrentRequests, int64(testDurationSec.Seconds()), int64(totalTime/time.Millisecond), totalRequests, avgRequestTime,
		int64(maxTime/time.Millisecond), int64(minTime/time.Millisecond), errors,
		int64(c.MinTimeThreshold/time.Millisecond), totalLessThanMin, lessThanPercent,
		int64(c.MinTimeThreshold/time.Millisecond), int64(c.MaxTimeThreshold/time.Millisecond), totalInMiddle, middlePercent,
		int64(c.MaxTimeThreshold/time.Millisecond), totalAboveThreshold, thresholdPercent)

	logger.Info(context.Background(), msg)

}

