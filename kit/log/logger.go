package logger

import (
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Environment string
	LogLevel    string
	LogFields   logrus.Fields
}

func New(cfg Config) *logrus.Logger {

	logger := logrus.New()
	timestampFormat := "2006-01-02 15:04:05.0000"
	cPrettyfier := func(frame *runtime.Frame) (function string, file string) {
		fileName := " " + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
		//return frame.Function, fileName
		return "", fileName
	}

	formatter := logrus.Formatter(&logrus.TextFormatter{
		CallerPrettyfier: cPrettyfier,
		DisableColors:    false,
		FullTimestamp:    true,
		TimestampFormat:  timestampFormat,
		QuoteEmptyFields: true,
	})

	if cfg.Environment == "prod" {
		formatter = &logrus.JSONFormatter{
			TimestampFormat:  timestampFormat,
			CallerPrettyfier: cPrettyfier,
		}
	}

	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)
	logger.SetFormatter(formatter)
	logger.SetReportCaller(level == logrus.DebugLevel)
	logger.Out = os.Stdout

	return logger
}
