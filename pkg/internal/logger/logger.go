package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: &logrus.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	},
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}
