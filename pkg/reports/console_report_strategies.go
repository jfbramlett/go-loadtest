package reports

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
)


func NewConsoleReportStrategy(minTimeThreshold, maxTimeThreshold int64) ReportStrategy {
	return &consoleReportStrategy{MinTimeThreshold: minTimeThreshold, MaxTimeThreshold: maxTimeThreshold}
}

type consoleReportStrategy struct {
	MinTimeThreshold			int64
	MaxTimeThreshold			int64
}

func (c *consoleReportStrategy) Report(concurrentRequests int, testDurationSec int64, results collector.ResultCollector) {
	var totalRequests int64
	var totalTime int64
	var maxTime int64
	var minTime int64
	var errors int
	var totalAboveThreshold int64
	var totalLessThanMin int64
	var totalInMiddle int64
	for _, t := range results.GetPassedTests() {
		totalRequests++
		totalTime = totalTime + t.Duration
		if t.Duration > maxTime {
			maxTime = t.Duration
		}
		if minTime == 0 || t.Duration < minTime {
			minTime = t.Duration
		}
		if t.Duration <= c.MinTimeThreshold {
			totalLessThanMin++
		} else if t.Duration > c.MinTimeThreshold && t.Duration <= c.MaxTimeThreshold {
			totalInMiddle++
		}
		if t.Duration > c.MaxTimeThreshold {
			totalAboveThreshold++
		}
	}
	errors = len(results.GetFailedTests())

	avgRequestTime := totalTime / totalRequests
	lessThanPercent := (float64(totalLessThanMin) / float64(totalRequests)) * float64(100)
	middlePercent := (float64(totalInMiddle) / float64(totalRequests)) * float64(100)
	thresholdPercent := (float64(totalAboveThreshold)/ float64(totalRequests)) * float64(100)

	utils.Logf(`Total Concurrent Requests: %d
Test Duration %d s
Total Time %d ms
Total Requests %d
Avg Time %d ms
Max Time %d ms
Min Time %d ms
Num Errors %d
Num Below %d ms %d (%.2f %%)
Num Between %d ms and %d ms %d (%.2f %%)
Num Above %d ms %d (%.2f %%)\n`,
		concurrentRequests, testDurationSec, totalTime, totalRequests, avgRequestTime, maxTime, minTime, errors,
		c.MinTimeThreshold, totalLessThanMin, lessThanPercent,
		c.MinTimeThreshold, c.MaxTimeThreshold, totalInMiddle, middlePercent,
		c.MinTimeThreshold, totalAboveThreshold, thresholdPercent)

}

