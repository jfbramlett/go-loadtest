package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger() *logrus.Entry {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	return logger.WithField("service", "load-test")
}
