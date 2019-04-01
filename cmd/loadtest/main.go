package main

import (
    "context"
    "github.com/jfbramlett/go-loadtest/pkg/collector"
    "github.com/jfbramlett/go-loadtest/pkg/loadprofile"
    "github.com/jfbramlett/go-loadtest/pkg/loadtester"
    "github.com/jfbramlett/go-loadtest/pkg/logging"
    "github.com/jfbramlett/go-loadtest/pkg/metrics"
    "github.com/jfbramlett/go-loadtest/pkg/naming"
    "github.com/jfbramlett/go-loadtest/pkg/reports"
    "github.com/jfbramlett/go-loadtest/pkg/utils"
    "math/rand"
    "time"
)

// our main function
func main() {
    metrics.Start()

    rand.Seed(time.Now().UTC().UnixNano())

    concurrentUsers := 5
    testLengthSec := 60
    testInterval := 2
    //loadProfileBuilder := loadprofile.NewStaticProfileBuilder(concurrentUsers, testLengthSec, testInterval)
    //loadProfileBuilder := loadprofile.NewRandomProfileBuilder(concurrentUsers, testLengthSec)
    loadProfileBuilder := loadprofile.NewPartialRandomProfileBuilder(concurrentUsers, testLengthSec, testInterval)


    resultCollector := collector.NewInMemoryRunCollector()

    test := loadtester.LoadTester{}

    test.Run(context.Background(), loadProfileBuilder, &Tester{}, naming.NewSimpleTestNamer(), resultCollector)

    reporter := reports.NewConsoleReportStrategy(time.Duration(500) * time.Millisecond, time.Duration(750) * time.Millisecond)

    reporter.Report(context.Background(), concurrentUsers, time.Duration(testLengthSec) * time.Second, resultCollector)
}


type Tester struct {
}

func (t *Tester) Run(ctx context.Context) error {
    logger, ctx := logging.GetLoggerFromContext(ctx, t)
    logger.Info(ctx, "Blah blah blah")
    time.Sleep(time.Duration(utils.RandomIntBetween(0, 5)) * time.Second)
    return nil
}