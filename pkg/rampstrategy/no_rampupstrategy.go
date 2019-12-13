package rampstrategy

import "time"

// noRampStrategy starts all users together without any delay
type noRampStrategy struct {
}


func (n *noRampStrategy) GetStartDelay(testLength time.Duration, maxUsers int) []StartDelay {
	return []StartDelay{{InitialDelay: time.Duration(0), UsersToStart: maxUsers}}
}

func NewNoRampUpStrategy() RampStrategy {
	return &noRampStrategy{}
}