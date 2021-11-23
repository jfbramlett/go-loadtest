package utils

import (
	"context"
	"github.com/sirupsen/logrus"
)

type testIdStruct struct{}
var testIdKey = testIdStruct{}

type loggingStruct struct{}
var loggingKey = loggingStruct{}


func LoggerFromContext(ctx context.Context) *logrus.Entry {
	var entry *logrus.Entry
	logger:= ctx.Value(loggingKey)
	if logger == nil {
		entry = NewLogger()
	} else {
		entry = logger.(*logrus.Entry)
	}

	return entry
}

func ContextWithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggingKey, logger)
}

func GetTestIdFromContext(ctx context.Context) string {
	testId := ctx.Value(testIdKey)
	if testId != nil {
		return testId.(string)
	}
	return "N/A"
}

func SetTestIdInContext(ctx context.Context, testId string) context.Context {
	return context.WithValue(ctx, testIdKey, testId)
}

