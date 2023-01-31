package tests

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type LoggerHook struct {
}

func (this *LoggerHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (this *LoggerHook) Fire(entry *logrus.Entry) error {
	fmt.Println("LoggerHook:", entry.Level, entry.Message)
	return nil
}

func TestLoggerHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&LoggerHook{})

	logger.Info("Hello World")
	logger.Warn("Hello World")
	logger.Errorf("Hello World")
}
