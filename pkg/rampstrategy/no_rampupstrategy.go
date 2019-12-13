package rampstrategy

import "time"

// noRampStrategy starts all users together without any delay
type noRampStrategy struct {
}


func (n *noRampStrategy) GetStartProfile(testLength time.Duration, maxUsers int) []StartProfile {
	return []StartProfile{{Delay: time.Duration(0), Users: maxUsers}}
}

func NewNoRampUpStrategy() RampStrategy {
	return &noRampStrategy{}
}