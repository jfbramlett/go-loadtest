package utils


type ResultCollector interface {
	AddRuntime(dur int64)
	GetRuntimes() []int64
	GetRunCount() int

	AddErrorRuntime(dur int64)
	GetErrorRuntimes() []int64
	GetErrorCount() int

	GetTotalRuns() int
}

type inmemoryResultCollector struct {
	times 			[]int64
	errors			[]int64
}

func (i *inmemoryResultCollector) AddRuntime(dur int64) {
	i.times = append(i.times, dur)
}

func (i *inmemoryResultCollector) GetRuntimes() []int64 {
	return i.times
}

func (i *inmemoryResultCollector) GetRunCount() int {
	return len(i.times)
}

func (i *inmemoryResultCollector) AddErrorRuntime(dur int64) {
	i.errors = append(i.errors, dur)
}

func (i *inmemoryResultCollector) GetErrorRuntimes() []int64 {
	return i.errors
}

func (i *inmemoryResultCollector) GetErrorCount() int {
	return len(i.errors)
}

func (i *inmemoryResultCollector) GetTotalRuns() int {
	return i.GetRunCount() + i.GetErrorCount()
}

func NewInMemoryRunCollector() ResultCollector {
	return &inmemoryResultCollector{times: make([]int64, 0), errors: make([]int64, 0)}
}
