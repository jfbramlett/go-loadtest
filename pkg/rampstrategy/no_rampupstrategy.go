package rampstrategy

import "time"

type noRampUpStrategy struct {
}


func (n *noRampUpStrategy) GetStartDelay(testLength time.Duration, maxUsers int) []StartDelay {
	return []StartDelay{{InitialDelay: time.Duration(0), UsersToStart: maxUsers}}
}

func NewNoRampUpStrategy() RampStrategy {
	return &noRampUpStrategy{}
}