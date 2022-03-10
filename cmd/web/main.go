package main

import (
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type envs struct {
	environment string
	host        string
	port        string
	allowedCORS string
	wait        time.Duration
	logJSON     bool
	logLevel    string
}

func main() {
	envs := loadEnvs()
	log := configLog(envs)

	log.Info("Hello, World!")
}

func configLog(envs envs) *logrus.Entry {

	ttFormat := "2006-01-02 15:04:05.123"
	if envs.environment == "prd" || envs.logJSON {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: ttFormat,
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors:    true,
			FullTimestamp:    true,
			TimestampFormat:  ttFormat,
			QuoteEmptyFields: true,
		})
	}

	level, err := logrus.ParseLevel(envs.logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	logrus.SetLevel(level)
	logrus.SetReportCaller(logrus.GetLevel() == logrus.DebugLevel)

	return logrus.WithFields(logrus.Fields{
		"app": "pre-separacao",
	})
}

func loadEnvs() envs {

	godotenv.Load(".env")

	envs := envs{
		environment: "dev",
		allowedCORS: os.Getenv("ALLOWED_CORS"),
		host:        os.Getenv("HOST_NAME"),
		port:        "8080",
		logLevel:    "info",
		logJSON:     false,
		wait:        time.Second * 15,
	}

	flag.StringVar(&envs.allowedCORS, "cors", envs.allowedCORS, "Cross-Origin Resource Sharing")
	flag.StringVar(&envs.host, "host", envs.host, "HTTP service host name")

	if env, ok := os.LookupEnv("ENVIRONMENT"); ok && env != "" {
		envs.environment = env
	}
	flag.StringVar(&envs.environment, "environment", envs.environment, "Runtime environment [dev, hlg, stg, prd]")

	if env, ok := os.LookupEnv("HOST_PORT"); ok && env != "" {
		envs.port = env
	}
	flag.StringVar(&envs.port, "port", envs.port, "HTTP service port")

	if env, ok := os.LookupEnv("LOG_LEVEL"); ok && env != "" {
		envs.logLevel = env
	}
	flag.StringVar(&envs.logLevel, "log-level", envs.logLevel, "Log output level for the server [debug, info, trace]")

	if env, ok := os.LookupEnv("LOG_JSON"); ok && env != "" {
		b, _ := strconv.ParseBool(env)
		envs.logJSON = b
	}
	flag.BoolVar(&envs.logJSON, "log-json", envs.logJSON, "Log output level for the server [debug, info, trace]")

	if env, ok := os.LookupEnv("GRACEFUL_TIMEOUT"); ok && env != "" {
		if w, err := time.ParseDuration(env); err == nil {
			envs.wait = w
		}
	}
	flag.DurationVar(&envs.wait, "graceful-timeout", envs.wait, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	flag.Parse()
	return envs
}
