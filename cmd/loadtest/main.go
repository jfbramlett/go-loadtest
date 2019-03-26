package main

import (
    "fmt"
    "github.com/jfbramlett/go-loadtest/pkg/loadtest"
    "github.com/jfbramlett/go-loadtest/pkg/reports"
    "github.com/jfbramlett/go-loadtest/pkg/runstrategy"
)

// our main function
func main() {
    loadtest.RunLoad(60,
        100,
        runstrategy.NewRandomDelayRunStrategyFactory(2, 10),
        reports.NewConsoleReportStrategy(int64(500), int64(1500)),
        TestFunc)

}


func TestFunc() (interface{}, error) {
    fmt.Println("Blah blah blah")

    return nil, nil
}