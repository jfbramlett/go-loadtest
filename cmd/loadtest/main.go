package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/jfbramlett/go-loadtest/pkg/testscenario"

	"github.com/jfbramlett/go-loadtest/pkg/reports"
)

// our main function
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	scenarioPtr := flag.String("scenario", "dev-scenario.json", "the file containing our test config")
	concurrentUsersPtr := flag.Int("concurrency", 50, "the number of concurrent requests (i.e. users")
	testLengthPtr := flag.Int("length", 300, "the number of seconds to run the test")
	intervalPtr := flag.Int("interval", 2, "number of seconds between calls")
	minTimePtr := flag.Int("minTime", 1000, "milliseconds for our min threshold (for reporting)")
	maxTimePtr := flag.Int("maxTime", 1500, "milliseconds for our min threshold (for reporting)")
	publishMetricsToPrometheus := flag.Bool("publishToPrometheus", false, "publish metrics to Prometheus")
	prometheusUrl := flag.String("prometheusUrl", "http://localhost:9091", "url for prometheus")
	publishMetricsToElastic := flag.Bool("publishToElastic", false, "publish metrics to elastic")
	elasticUrl := flag.String("elasticUrl", "http://alias-rsrv:9200", "url for elastic")
	elasticIndex := flag.String("elasticIndex", "metrics-2020-07", "index used when writting to elastic")
	flag.Parse()

	concurrentUsers := *concurrentUsersPtr
	testLength := time.Duration(*testLengthPtr) * time.Second
	testInterval := time.Duration(*intervalPtr) * time.Second
	minThreshold := time.Duration(*minTimePtr) * time.Millisecond
	maxThreshold := time.Duration(*maxTimePtr) * time.Millisecond

	fmt.Println("================================================")
	fmt.Println("Running load tester")
	fmt.Printf("Concurrent users: %v\n", concurrentUsers)
	fmt.Printf("Test Length: %v\n", testLength)
	fmt.Printf("Test Interval: %v\n", testInterval)
	fmt.Printf("Reporting Min Threshold: %v\n", minThreshold)
	fmt.Printf("Reporting Max Threshold: %v\n", maxThreshold)
	fmt.Printf("Scenario: %v\n", *scenarioPtr)
	fmt.Println("================================================")

	resultsCollector := testscenario.NewInMemoryRunCollector()
	resultsCollector.Start()

	scenario := testscenario.NewTestScenario(
		testscenario.NewWeightedTestFunc(
			testscenario.NewWeightedTest(testFunc("test 1"), 2),
			testscenario.NewWeightedTest(testFunc("test 2"), 8),
		),
		nil, nil, concurrentUsers, testLength, testscenario.RandomStartStrategy, testscenario.FixedPauseStrategy(testInterval),
	)

	_ = scenario.Run(context.Background(), resultsCollector)

	if *publishMetricsToPrometheus {
		prometheusReporter := reports.NewPrometheusReportStrategy("perf-test", minThreshold, maxThreshold, *prometheusUrl)
		prometheusReporter.Report(context.Background(), concurrentUsers, testLength, resultsCollector)
	}

	if *publishMetricsToElastic {
		elasticReport := reports.NewElasticReportStrategy(*elasticUrl, *elasticIndex, minThreshold, maxThreshold)
		elasticReport.Report(context.Background(), concurrentUsers, testLength, resultsCollector)
	}

	reporter := reports.NewConsoleReportStrategy(minThreshold, maxThreshold)
	reporter.Report(context.Background(), concurrentUsers, testLength, resultsCollector)
}

func testFunc(name string) testscenario.TestFunc {
	return testscenario.NewTest(name, func(ctx context.Context) error {
		return nil
	})
}
