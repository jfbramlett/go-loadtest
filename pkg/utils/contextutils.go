package utils

import (
	"context"

	"github.com/sirupsen/logrus"
)

type loggingStruct struct{}

var loggingKey = loggingStruct{}

func LoggerFromContext(ctx context.Context) *logrus.Entry {
	var entry *logrus.Entry
	logger := ctx.Value(loggingKey)
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
