package tests

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello World")
}

func TestLoggerAllLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Hello World")
	logger.Debug("Hello World")
	logger.Info("Hello World")
	logger.Warn("Hello World")
	logger.Error("Hello World")
	logger.Fatal("Hello World")
	logger.Panic("Hello World")
}

func TestLogToFile(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile(
		"../application.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	logger.SetOutput(file)

	logger.Info("Hello World")
}

func TestLogWithJsonFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Info("Hello World")
}

func TestLoggerWithField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("username", "arfanxn").WithField("email", "arf@gm.com").Info("Has logged in")

	logger.WithFields(logrus.Fields{
		"username": "arfanxn",
		"email":    "arf@gm.com",
	}).Infof("Has logged in")
}

func TestLogManuallyWithEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithFields(logrus.Fields{
		"username": "arfanxn",
		"email":    "arf@gm.com",
	}).Info("Has logged in")
}

func TestLogHook(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithFields(logrus.Fields{
		"username": "arfanxn",
		"email":    "arf@gm.com",
	}).Info("Has logged in")
}
