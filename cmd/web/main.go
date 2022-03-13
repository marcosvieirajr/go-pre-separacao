package main

import (
	"github.com/sirupsen/logrus"
	"github.com/viavarejo-internal/pre-separacao-api/kit/config"
	logger "github.com/viavarejo-internal/pre-separacao-api/kit/log"
)

var (
	cfg config.Config
	log *logrus.Entry
)

func init() {
	cfg = config.Load()
	log = logger.Config(cfg.Environment, cfg.LogFormat, cfg.LogLevel).
		WithFields(logrus.Fields{
			"app": "pre-separacao",
		})
}

func main() {
	log.Info("Hello, RESTful World!")
	log.Info(cfg.Environment)
}
