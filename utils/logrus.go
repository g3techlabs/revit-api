package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

type CustomFormatter struct {
	logrus.TextFormatter
}

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	Log.SetFormatter(&CustomFormatter{
		TextFormatter: logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: "02-01-2006 15:04:05",
		},
	})

	Log.SetOutput(os.Stdout)
}
