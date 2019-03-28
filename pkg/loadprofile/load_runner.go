package loadprofile

import (
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
)

type LoadRunner struct {
}

func (l *LoadRunner) Run(loadProfile LoadProfile, runFunc utils.RunFunc, namer naming.TestNamer, resultCollector collector.ResultCollector) {
	utils.Log("Starting runners")
	resultCollector.Start()

	wg := sync.WaitGroup{}
	for _, r := range loadProfile.GetLoad(namer, runFunc, resultCollector) {
		wg.Add(1)
		go l.runWrapper(r, &wg)
	}

	utils.Log("Waiting for tests to end")
	wg.Wait()
	utils.Log("Tests completed")
	resultCollector.Stop()
}


func (l *LoadRunner) runWrapper(load Load, wg *sync.WaitGroup) {
	load.Run()
	wg.Done()
}

