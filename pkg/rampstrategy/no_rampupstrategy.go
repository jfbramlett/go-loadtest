package rampstrategy

import (
	"context"
	"time"
)

// noRampStrategy starts all users together without any delay
type noRampStrategy struct {
}


func (n *noRampStrategy) GetStartProfile(ctx context.Context, testLength time.Duration, maxUsers int) []StartProfile {
	return []StartProfile{{Delay: time.Duration(0), Users: maxUsers}}
}

func NewNoRampUpStrategy() RampStrategy {
	return &noRampStrategy{}
}