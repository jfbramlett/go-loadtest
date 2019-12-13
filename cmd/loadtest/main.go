package main

import (
    "context"
    "github.com/jfbramlett/go-loadtest/pkg/loadprofile"
    "github.com/jfbramlett/go-loadtest/pkg/loadtester"
    "github.com/jfbramlett/go-loadtest/pkg/logging"
    "github.com/jfbramlett/go-loadtest/pkg/metrics"
    "github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
    "github.com/jfbramlett/go-loadtest/pkg/reports"
    "math/rand"
    "time"
)

// our main function
func main() {
    metrics.Start()

    rand.Seed(time.Now().UTC().UnixNano())

    concurrentUsers := 100
    testLength := 300 * time.Second
    testInterval := 2*time.Second

    runner := loadtester.NewLoadTester(concurrentUsers, testLength, testInterval, loadprofile.StaticProfile, rampstrategy.Smooth)
    //runner := loadtester.NewLoadTester(concurrentUsers, testLength, testInterval, loadprofile.RandomProfile, rampstrategy.Smooth)
    //runner := loadtester.NewLoadTester(concurrentUsers, testLength, testInterval, loadprofile.PartialRandomProfile, rampstrategy.Smooth)

    collector := runner.Run(context.Background(), &Tester{})

    reporter := reports.NewConsoleReportStrategy(time.Duration(500) * time.Millisecond, time.Duration(750) * time.Millisecond)

    reporter.Report(context.Background(), concurrentUsers, testLength, collector)
}


type Tester struct {
}

func (t *Tester) Run(ctx context.Context) error {
    logger, ctx := logging.GetLoggerFromContext(ctx, t)
    logger.Info(ctx, "Blah blah blah")
    return nil
}