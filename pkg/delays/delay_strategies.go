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
	delaySec	time.Duration
}

func (d *fixedDelayStrategy) GetTicker() *time.Ticker {
	fmt.Println(fmt.Sprintf("using fixed delay of %d ms\n", d.delaySec))
	return time.NewTicker(time.Duration(time.Second * d.delaySec))
}

func NewFixedDelayStrategyFactory(delaySec int64) DelayStrategy {
	fmt.Println("Using fixed delay strategy of %d s", delaySec)
	return &fixedDelayStrategy{delaySec: time.Duration(delaySec)}
}

type randomDelayStrategy struct {
	delayMinSec		int
	delayMaxSec		int
}


func (d *randomDelayStrategy) GetTicker() *time.Ticker {
	randInterval := time.Duration(rand.Intn(d.delayMaxSec - d.delayMinSec) + d.delayMinSec)
	fmt.Println(fmt.Sprintf("using random delay of %d sec", randInterval))
	return time.NewTicker(time.Duration(time.Second * randInterval))
}

func NewRandomDelayStrategy(delayMinSec int, delayMaxSec int) DelayStrategy {
	rand.Seed(time.Now().Unix())
	return &randomDelayStrategy{delayMinSec: delayMinSec, delayMaxSec: delayMaxSec}
}
