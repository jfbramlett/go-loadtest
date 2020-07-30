package reports

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/ninthwave/nwp-load-test/pkg/collector"
	"github.com/ninthwave/nwp-load-test/pkg/logging"
	"github.com/ninthwave/nwp-load-test/pkg/metrics"
	"strings"
	"time"
)

const verticle = "nwp-load-test"

func NewElasticReportStrategy(elasticUrl string, indexName string, minTimeThreshold, maxTimeThreshold time.Duration) ReportStrategy {
	return &elasticReportStrategy{ElasticUrl: elasticUrl, IndexName: indexName,
		MinTimeThreshold: minTimeThreshold, MaxTimeThreshold: maxTimeThreshold}
}

type elasticReportStrategy struct {
	MinTimeThreshold			time.Duration
	MaxTimeThreshold			time.Duration
	ElasticUrl					string
	IndexName					string
}

func (e *elasticReportStrategy) Report(ctx context.Context, concurrentRequests int, testDurationSec time.Duration, results collector.ResultCollector) {
	var totalRequests int64
	var totalTime time.Duration
	var maxTime time.Duration
	var minTime time.Duration
	var errors int
	var totalAboveThreshold int64
	var totalLessThanMin int64
	var totalInMiddle int64

	logger, ctx := logging.GetLoggerFromContext(ctx, e)

	cfg := elasticsearch.Config{
		Addresses: []string{
			e.ElasticUrl,
		},

	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Error(ctx, err, "failed to connect to elastic");
		return
	}

	for _, t := range results.GetPassedTests() {
		for _, m := range t.Metrics() {
			_ = e.writeMetric(ctx, es, m)
		}

		totalRequests++
		totalTime = totalTime + t.Duration()
		if t.Duration() > maxTime {
			maxTime = t.Duration()
		}
		if minTime == 0 || t.Duration() < minTime {
			minTime = t.Duration()
		}
		if t.Duration() <= e.MinTimeThreshold {
			totalLessThanMin++
		} else if t.Duration() > e.MinTimeThreshold && t.Duration() <= e.MaxTimeThreshold {
			totalInMiddle++
		}
		if t.Duration() > e.MaxTimeThreshold {
			totalAboveThreshold++
		}
	}
	errors = len(results.GetFailedTests())

	avgRequestTime := int64(totalTime/time.Millisecond) / totalRequests
	lessThanPercent := (float64(totalLessThanMin) / float64(totalRequests)) * float64(100)
	middlePercent := (float64(totalInMiddle) / float64(totalRequests)) * float64(100)
	thresholdPercent := (float64(totalAboveThreshold)/ float64(totalRequests)) * float64(100)

	reportTime := time.Now().Format(time.RFC3339)
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_total_requests", float64(totalRequests + int64(errors)), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_concurrent_requests", float64(concurrentRequests), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_duration", float64(testDurationSec), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_total_time", float64(int64(totalTime/time.Millisecond)), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_passed", float64(totalRequests), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_avg_time", float64(avgRequestTime), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_max_request_time", float64(int64(maxTime/time.Millisecond)), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_min_request_time", float64(int64(minTime/time.Millisecond)), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_errors", float64(errors), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_requests_below_threshold", float64(totalLessThanMin), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_requests_below_threshold_pct", lessThanPercent, reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_requests_between_threshold", float64(totalInMiddle), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_requests_between_threshold_pct", middlePercent, reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_requests_above_threshold", float64(totalAboveThreshold), reportTime))
	_ = e.writeMetric(ctx, es, metrics.NewGauge(verticle, "perf_requests_above_threshold_percent", thresholdPercent, reportTime))
}

func (e *elasticReportStrategy) writeMetric(ctx context.Context, esClient *elasticsearch.Client, metric metrics.Metric) error {
	logger, ctx := logging.GetLoggerFromContext(ctx, e)

	values := metric.Values()
	jsonString, err := json.Marshal(values)
	if err != nil {
		logger.Error(ctx, err, "failed to convert metric to json")
		return err
	}

	req := esapi.IndexRequest{
		Index:      e.IndexName,
		Body:       strings.NewReader(string(jsonString)),
		Refresh:    "true",
		DocumentType: "doc",
	}

	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		logger.Error(ctx, err, "Error getting response")
		return err
	}
	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.Infof(ctx, "Error parsing the response body: %s", err)
	} else {
		if res.IsError() {
			err = errors.New("failed to record metric in elastic")
			logger.Errorf(ctx, err, "%v", r)
			return err
		}
	}

	return nil
}


