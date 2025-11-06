package utils

import (
	"github.com/sirupsen/logrus"
)

type ILogger interface {
	Info(args ...any)
	Infof(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Warnf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

var Log *logrus.Logger

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() ILogger {
	logger := logrus.New()

	logger.SetFormatter(
		&logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: "02-01-2006 15:04:05",
		},
	)

	Log = logger

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	l.logger.Warnf(format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
