package logging

import (
	"context"
	"fmt"
	"github.com/jfbramlett/go-loadtest/pkg/utils"
	"time"
)


type simpleLogger struct {
	cls		string
}

func (l *simpleLogger) Info(ctx context.Context, msg string) {
	l.log("INFO", ctx, msg)
}

func (l *simpleLogger) Infof(ctx context.Context, fmtString string, params ...interface{}) {
	l.log("INFO", ctx, l.toMsg(fmtString, params))
}

func (l *simpleLogger) Debug(ctx context.Context, msg string) {
	l.log("DEBUG", ctx, msg)
}

func (l *simpleLogger) Debugf(ctx context.Context, fmtString string, params ...interface{}) {
	l.log("DEBUG", ctx, l.toMsg(fmtString, params))
}


func (l *simpleLogger) Warn(ctx context.Context, msg string) {
	l.log("WARN", ctx, msg)
}

func (l *simpleLogger) Warnf(ctx context.Context, fmtString string, params ...interface{}) {
	l.log("WARN", ctx, l.toMsg(fmtString, params))
}


func (l *simpleLogger) Error(ctx context.Context, err error, msg string) {
	l.log("ERROR", ctx, fmt.Sprintf("%s: %s", msg, err))

}

func (l *simpleLogger) Errorf(ctx context.Context, err error, fmtString string, params ...interface{}) {
	l.log("ERROR", ctx, l.toMsg("%s: %s", l.toMsg(fmtString, params), err))
}

func (l* simpleLogger) toMsg(fmtString string, params ...interface{}) string {
	return fmt.Sprintf(fmtString, params)
}

func (l *simpleLogger) log(level string, ctx context.Context, msg string) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	fmtMsg := fmt.Sprintf("%s [%s] [%s] [%s] %s", ts, l.cls, utils.GetTestId(ctx), level, msg)
	fmt.Println(fmtMsg)
}

func (l *simpleLogger) NewLogger(cls interface{}) Logger {
	return NewSimpleLogger(cls)
}


func NewSimpleLogger(cls interface{}) Logger {
	return &simpleLogger{cls: fmt.Sprintf("%T", cls)}
}