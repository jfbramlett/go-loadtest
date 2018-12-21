package main

import (
    "database/sql"
    "fmt"
    "github.com/jfbramlett/go-loadtest/pkg/db"
)

// our main function
func main() {
    dbConnection, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/asapp_dev_companies1")

    if err != nil {
        fmt.Println(err)
    }
    repActivity := db.RepActivity{DBConnection: dbConnection}
    repActivity.AddActivity()

/*
    loadtest.RunLoad(60,
        100,
        delays.NewRandomDelayStrategy(2000, 4000),
        reports.NewConsoleReportStrategy(int64(500), int64(1500)),
        &functionWrapper{})
*/
}


type functionWrapper struct {
}

func (g *functionWrapper) Run() (interface{}, error) {
    fmt.Println("Blah blah blah")

    return nil, nil
}