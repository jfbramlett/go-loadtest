package loadtester

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/loadprofile"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
)

type LoadTester struct {
}

func (l *LoadTester) Run(loadProfileBuilder loadprofile.LoadProfileBuilder, runFunc utils.RunFunc, namer naming.TestNamer, resultCollector collector.ResultCollector) {
	utils.Log("Starting runners")
	resultCollector.Start()

	wg := sync.WaitGroup{}
	for i, r := range loadProfileBuilder.GetLoadProfiles(runFunc, resultCollector) {
		wg.Add(1)
		ctx := context.WithValue(context.Background(), "testId", namer.GetName(i))
		go l.runWrapper(r, ctx, &wg)
	}

	utils.Log("Waiting for tests to end")
	wg.Wait()
	utils.Log("Tests completed")
	resultCollector.Stop()
}


func (l *LoadTester) runWrapper(load loadprofile.LoadProfile, ctx context.Context, wg *sync.WaitGroup) {
	load.Run(ctx)
	wg.Done()
}

