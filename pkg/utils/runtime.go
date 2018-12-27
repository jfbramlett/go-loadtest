package utils


type RunTimes struct {
	Times 			[]int64
	Errors			[]int64
}

func NewRunTimes() *RunTimes {
	return &RunTimes{Times: make([]int64, 0), Errors: make([]int64, 0)}
}
