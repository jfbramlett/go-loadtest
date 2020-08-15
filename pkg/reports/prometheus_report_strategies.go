package reports

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/push"
	"time"
)

func NewPrometheusReportStrategy(reportName string, minTimeThreshold, maxTimeThreshold time.Duration, prometheusUrl string) ReportStrategy {
	return &prometheusReportStrategy{MinTimeThreshold: minTimeThreshold,
		MaxTimeThreshold: maxTimeThreshold,
		PrometheusUrl: prometheusUrl,
		Name: reportName}
}

type prometheusReportStrategy struct {
	MinTimeThreshold			time.Duration
	MaxTimeThreshold			time.Duration
	PrometheusUrl				string
	Name						string
}

func (p *prometheusReportStrategy) Report(ctx context.Context, concurrentRequests int, testDurationSec time.Duration, results collector.ResultCollector) {

	logger, ctx := logging.GetLoggerFromContext(ctx, p)
	pusher := push.New(p.PrometheusUrl, p.Name)

	var totalRequests int64
	var totalTime time.Duration
	var maxTime time.Duration
	var minTime time.Duration
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
		if t.Duration() <= p.MinTimeThreshold {
			totalLessThanMin++
		} else if t.Duration() > p.MinTimeThreshold && t.Duration() <= p.MaxTimeThreshold {
			totalInMiddle++
		}
		if t.Duration() > p.MaxTimeThreshold {
			totalAboveThreshold++
		}
	}

	avgRequestTime := int64(totalTime/time.Millisecond) / totalRequests
	lessThanPercent := (float64(totalLessThanMin) / float64(totalRequests)) * float64(100)
	middlePercent := (float64(totalInMiddle) / float64(totalRequests)) * float64(100)
	thresholdPercent := (float64(totalAboveThreshold)/ float64(totalRequests)) * float64(100)

	concurrentRequestsGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_concurrent_requests",
	})
	pusher.Collector(concurrentRequestsGuage)
	concurrentRequestsGuage.Set(float64(concurrentRequests))

	totalRequestsGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_total_requests",
	})
	pusher.Collector(totalRequestsGuage)
	totalRequestsGuage.Set(float64(totalRequests))

	durationGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_duration_s",
	})
	pusher.Collector(durationGuage)
	durationGuage.Set(float64(testDurationSec))

	totalTimeGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_total_time_ms",
	})
	pusher.Collector(totalTimeGuage)
	totalTimeGuage.Set(float64(totalTime))

	passedGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_passed_tests",
	})
	pusher.Collector(passedGuage)
	passedGuage.Set(float64(len(results.GetPassedTests())))

	failedGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_failed_tests",
	})
	pusher.Collector(failedGuage)
	failedGuage.Set(float64(len(results.GetFailedTests())))

	avgTimeGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_avg_request_time_ms",
	})
	pusher.Collector(avgTimeGuage)
	avgTimeGuage.Set(float64(avgRequestTime))

	minTimeGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_min_request_time_ms",
	})
	pusher.Collector(minTimeGuage)
	minTimeGuage.Set(float64(minTime/time.Millisecond))

	maxTimeGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_max_request_time_ms",
	})
	pusher.Collector(maxTimeGuage)
	maxTimeGuage.Set(float64(maxTime/time.Millisecond))

	countBelowGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_count_below_threshold",
	})
	pusher.Collector(countBelowGuage)
	countBelowGuage.Set(float64(totalLessThanMin))

	countMiddleGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_count_within_threshold",
	})
	pusher.Collector(countMiddleGuage)
	countMiddleGuage.Set(float64(totalInMiddle))

	countAboveGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_count_above_threshold",
	})
	pusher.Collector(countAboveGuage)
	countAboveGuage.Set(float64(totalAboveThreshold))

	percentBelowGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_percent_below_threshold",
	})
	pusher.Collector(percentBelowGuage)
	percentBelowGuage.Set(lessThanPercent)

	percentMiddleGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_percent_within_threshold",
	})
	pusher.Collector(percentMiddleGuage)
	percentMiddleGuage.Set(middlePercent)

	percentAboveGuage := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "perf_percent_above_threshold",
	})
	pusher.Collector(percentAboveGuage)
	percentAboveGuage.Set(thresholdPercent)

	logger.Info(ctx, "sending metrics to prometheus gateway")
	err := pusher.Push()
	if err != nil {
		logger.Error(ctx, err, "failed to send metrics to prometheus gateway")
		return
	}

	logger.Info(ctx, "successfully sent metrics to prometheus gateway")
}


