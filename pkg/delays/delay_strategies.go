package delays

import (
	"fmt"
	"math/rand"
	"time"
)


type DelayStrategy interface {
	GetTicker() *time.Ticker
}

type fixedDelayStrategy struct {
	delayMillis	time.Duration
}

func (d *fixedDelayStrategy) GetTicker() *time.Ticker {
	fmt.Printf("using fixed delay of %d ms\n", d.delayMillis)
	return time.NewTicker(time.Duration(time.Millisecond * d.delayMillis))
}

func NewFixedDelayStrategyFactory(delayMillis int64) DelayStrategy {
	fmt.Println("Using fixed delay strategy of %d ms", delayMillis)
	return &fixedDelayStrategy{delayMillis: time.Duration(delayMillis)}
}

type randomDelayStrategy struct {
	delayMinMillis		int
	delayMaxMillis		int
}


func (d *randomDelayStrategy) GetTicker() *time.Ticker {
	randInterval := time.Duration(rand.Intn(d.delayMaxMillis - d.delayMinMillis) + d.delayMaxMillis)
	fmt.Printf("using random delay of %d ms\n", randInterval)
	return time.NewTicker(time.Duration(time.Millisecond * randInterval))
}

func NewRandomDelayStrategy(delayMinMillis int, delayMaxMillis int) DelayStrategy {
	rand.Seed(time.Now().Unix())
	return &randomDelayStrategy{delayMinMillis: delayMinMillis, delayMaxMillis: delayMaxMillis}
}
