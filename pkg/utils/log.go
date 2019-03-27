package utils

import (
	"fmt"
	"time"
)

func Log(msg string) {
	fmt.Println(time.Now().Format(time.StampMilli) + " : " + msg)
}

func Logf(msg string, params ...interface{}) {
	fmt.Println(time.Now().Format(time.StampMilli) + " " + fmt.Sprintf(msg, params...))
}

func Logt(testId, msg string) {
	fmt.Println(time.Now().Format(time.StampMilli) + " [" + testId + "] " + msg)
}

func Logtf(testId, msg string, params ...interface{}) {
	fmt.Println(time.Now().Format(time.StampMilli) + " [" + testId + "] " + fmt.Sprintf(msg, params...))
}