package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

type ILogger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type Logger struct {
	logrus.TextFormatter
}

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	Log.SetFormatter(&Logger{
		TextFormatter: logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: "02-01-2006 15:04:05",
		},
	})

	Log.SetOutput(os.Stdout)
}
