package rampstrategy

import (
    "context"
    "fmt"
    "testing"
    "time"
)

func TestNewSmoothRampUpStrategy(t *testing.T) {
    rampStrat := NewSmoothRampUpStrategy(.10)

    strats := rampStrat.GetStartProfile(context.Background(), time.Duration(600*time.Second), 200)

    for _, s := range strats {
        fmt.Println(s)
    }
}
