package logging

import (
	"context"
	"fmt"
)


type Logger interface {
	Info(ctx context.Context, msg string)
	Infof(ctx context.Context, fmtString string, params ...interface{})
	Debug(ctx context.Context, msg string)
	Debugf(ctx context.Context, fmtString string, params ...interface{})
	Warn(ctx context.Context, msg string)
	Warnf(ctx context.Context, fmtString string, params ...interface{})
	Error(ctx context.Context, err error, msg string)
	Errorf(ctx context.Context, err error, fmtString string, params ...interface{})

	NewLogger(tClass interface{}) Logger
}


func GetLogger(ctx context.Context, tClass interface{}) Logger {
	ctxLogger := ctx.Value("logger")
	if ctxLogger != nil {
		logger := ctxLogger.(Logger)
		return logger.NewLogger(fmt.Sprintf("%T", tClass))
	}
	return NewSimpleLogger(tClass)
}