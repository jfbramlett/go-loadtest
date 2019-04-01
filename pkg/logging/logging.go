package logging

import (
	"context"
	"fmt"
)


const loggerKey = "logger"

type Logger interface {
	Info(ctx context.Context, msg string)
	Infof(ctx context.Context, fmtString string, params ...interface{})
	Debug(ctx context.Context, msg string)
	Debugf(ctx context.Context, fmtString string, params ...interface{})
	Warn(ctx context.Context, msg string)
	Warnf(ctx context.Context, fmtString string, params ...interface{})
	Error(ctx context.Context, err error, msg string)
	Errorf(ctx context.Context, err error, fmtString string, params ...interface{})
}


func setLoggerInContext(ctx context.Context, loggerFor string, logger Logger) context.Context {
	return context.WithValue(ctx, loggerFor, logger)
}

func GetLoggerFromContext(ctx context.Context, tClass interface{}) (Logger, context.Context) {
	cls := fmt.Sprintf("%T", tClass)
	logKey := fmt.Sprintf("%s.%s", loggerKey, cls)

	ctxLogger := ctx.Value(logKey)
	if ctxLogger != nil {
		return ctxLogger.(Logger), ctx

	}

	logger := NewSimpleLogger(cls)
	updatedCtx := setLoggerInContext(ctx, logKey, logger)
	return logger, updatedCtx
}