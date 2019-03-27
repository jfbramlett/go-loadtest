package main

import (
    "github.com/jfbramlett/go-loadtest/pkg/collector"
    "github.com/jfbramlett/go-loadtest/pkg/loadtest"
    "github.com/jfbramlett/go-loadtest/pkg/naming"
    "github.com/jfbramlett/go-loadtest/pkg/rampstrategy"
    "github.com/jfbramlett/go-loadtest/pkg/reports"
    "github.com/jfbramlett/go-loadtest/pkg/runstrategy"
    "github.com/jfbramlett/go-loadtest/pkg/utils"
    "math/rand"
    "time"
)

// our main function
func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    loadtest.RunLoad(300,
        100,
        rampstrategy.NewSmoothRampUpStrategy(60, 100, 5),
        runstrategy.NewRandomDelayRunStrategyFactory(2, 10),
        naming.NewSimpleTestNamer(),
        collector.NewInMemoryRunCollector(),
        reports.NewConsoleReportStrategy(int64(500), int64(1500)),
        TestFunc)

}


func TestFunc() (interface{}, error) {
    utils.Log("Blah blah blah")

    return nil, nil
}