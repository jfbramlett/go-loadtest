package main

import (
    "fmt"
    "github.com/jfbramlett/go-loadtest/pkg/delays"
    "github.com/jfbramlett/go-loadtest/pkg/loadtest"
    "github.com/jfbramlett/go-loadtest/pkg/reports"
)

// our main function
func main() {
    loadtest.RunLoad(60,
        100,
        delays.NewRandomDelayStrategy(2000, 4000),
        reports.NewConsoleReportStrategy(int64(500), int64(1500)),
        &functionWrapper{})
}


type functionWrapper struct {
}

func (g *functionWrapper) Run() (interface{}, error) {
    fmt.Println("Blah blah blah")

    return nil, nil
}