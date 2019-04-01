package loadtester

import (
	"context"
	"github.com/jfbramlett/go-loadtest/pkg/collector"
	"github.com/jfbramlett/go-loadtest/pkg/loadprofile"
	"github.com/jfbramlett/go-loadtest/pkg/logging"
	"github.com/jfbramlett/go-loadtest/pkg/naming"
	"github.com/jfbramlett/go-loadtest/pkg/testwrapper"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"sync"
)

type LoadTester struct {
}

func (l *LoadTester) Run(ctx context.Context, loadProfileBuilder loadprofile.LoadProfileBuilder, runFunc testwrapper.Test, namer naming.TestNamer, resultCollector collector.ResultCollector) {
	logger, ctx := logging.GetLoggerFromContext(ctx, l)

	logger.Info(ctx, "Starting runners")
	resultCollector.Start()

	wg := sync.WaitGroup{}
	for i, r := range loadProfileBuilder.GetLoadProfiles(runFunc, resultCollector) {
		wg.Add(1)
		ctx := utils.SetTestIdInContext(ctx, namer.GetName(i))
		go l.runWrapper(r, ctx, &wg)
	}

	logger.Info(context.Background(), "Waiting for tests to end")
	wg.Wait()
	logger.Info(context.Background(), "Tests completed")
	resultCollector.Stop()
}


func (l *LoadTester) runWrapper(load loadprofile.LoadProfile, ctx context.Context, wg *sync.WaitGroup) {
	load.Run(ctx)
	wg.Done()
}

