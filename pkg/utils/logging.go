package utils

import (
	"context"
	"os"

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

func NewLogger() *logrus.Entry {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	return logger.WithField("service", "load-test")
}
