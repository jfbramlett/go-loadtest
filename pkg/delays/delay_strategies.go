package delays

import (
	"math/rand"
	"time"
)

type DelayStrategy interface {
	Wait()
}

func NewFixedDelayStrategy(delayMillis int) DelayStrategy {
	return &fixedDelayStrategy{delayMillis: delayMillis}
}

type fixedDelayStrategy struct {
	delayMillis		int
}

func (d *fixedDelayStrategy) Wait() {
	time.Sleep(time.Millisecond * time.Duration(d.delayMillis))
}

func NewRandomDelayStrategy(delayMinMillis int, delayMaxMillis int) DelayStrategy {
	return &randomDelayStrategy{delayMinMillis: delayMinMillis, delayMaxMillis: delayMaxMillis}
}

type randomDelayStrategy struct {
	delayMinMillis		int
	delayMaxMillis		int
}

func (d *randomDelayStrategy) Wait() {
	rand.Seed(time.Now().Unix())
	randInterval := int64(rand.Intn(d.delayMaxMillis - d.delayMinMillis) + d.delayMaxMillis)
	time.Sleep(time.Millisecond * time.Duration(randInterval))
}

func NewNoDelayStrategy() DelayStrategy {
	return &noDelayStrategy{}
}

type noDelayStrategy struct {

}

func (d *noDelayStrategy) Wait() {
}

