package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var l = initLogger()

func initLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}

func getLogger() *logrus.Logger {
	return l
}

func Infof(message string, args ...interface{}) {
	getLogger().Infof(message, args...)
}

func Errorf(message string, args ...interface{}) {
	getLogger().Errorf(message, args...)
}

func Debugf(message string, args ...interface{}) {
	getLogger().Debugf(message, args...)
}

func Fatalf(message string, args ...interface{}) {
	getLogger().Fatalf(message, args...)
}

func Info(message string) {
	getLogger().Info(message)
}

func Error(message string) {
	getLogger().Error(message)
}
