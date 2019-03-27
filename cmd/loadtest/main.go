package main

import (
    "github.com/jfbramlett/go-loadtest/pkg/collector"
    "github.com/jfbramlett/go-loadtest/pkg/loadprofile"
    "github.com/jfbramlett/go-loadtest/pkg/naming"
    "github.com/jfbramlett/go-loadtest/pkg/reports"
    "github.com/jfbramlett/go-loadtest/pkg/utils"
    "math/rand"
    "time"
)

// our main function
func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    concurrentUsers := 5
    testLengthSec := 60
    //testInterval := 2
    //loadProfile := loadprofile.NewStaticProfile(concurrentUsers, testLengthSec, testInterval)
    loadProfile := loadprofile.NewRandomProfile(concurrentUsers, testLengthSec)


    resultCollector := collector.NewInMemoryRunCollector()

    loadRunner := loadprofile.LoadRunner{}

    loadRunner.Run(loadProfile,
        TestFunc,
        naming.NewSimpleTestNamer(),
        resultCollector)

    reporter := reports.NewConsoleReportStrategy(time.Duration(500) * time.Millisecond, time.Duration(750) * time.Millisecond)

    reporter.Report(concurrentUsers, time.Duration(testLengthSec) * time.Second, resultCollector)
}


func TestFunc() (interface{}, error) {
    utils.Log("Blah blah blah")
    time.Sleep(time.Duration(utils.RandomIntBetween(0, 5)) * time.Second)
    return nil, nil
}