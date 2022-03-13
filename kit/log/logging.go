package logger

import (
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	Environment string
	LogFormat   string
	LogLevel    string
}

// type envs struct {
// 	environment string
// 	host        string
// 	port        string
// 	allowedCORS string
// 	wait        time.Duration
// 	logJSON     bool
// 	logLevel    string
// }

// func main() {
// 	envs := loadEnvs()
// 	log := configLog(envs).WithFields(logrus.Fields{
// 		"app": "pre-separacao",
// 	})

// 	log.Info("Hello, World!")
// }

// func ConfigLog(envs envs.Envs) *logrus.Logger {
// func Config(cfg Config) *logrus.Logger {
func Config(environment, logFormat, logLevel string) *logrus.Logger {

	logger := logrus.New()
	timestampFormat := "2006-01-02 15:04:05.0000"

	formatter := logrus.Formatter(&logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := " " + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			//return frame.Function, fileName
			return "", fileName
		},
		DisableColors:    false,
		FullTimestamp:    true,
		TimestampFormat:  timestampFormat,
		QuoteEmptyFields: true,
	})

	if environment == "prd" || logFormat == "json" {
		formatter = &logrus.JSONFormatter{
			TimestampFormat: timestampFormat,
		}
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)
	logger.SetFormatter(formatter)
	logger.SetReportCaller(level == logrus.DebugLevel)
	logger.Out = os.Stdout

	return logger
}

// func loadEnvs() envs {

// 	godotenv.Load(".env")

// 	envs := envs{
// 		environment: "dev",
// 		allowedCORS: os.Getenv("ALLOWED_CORS"),
// 		host:        os.Getenv("HOST_NAME"),
// 		port:        "8080",
// 		logLevel:    "info",
// 		logJSON:     false,
// 		wait:        time.Second * 15,
// 	}

// 	flag.StringVar(&envs.allowedCORS, "cors", envs.allowedCORS, "Cross-Origin Resource Sharing")
// 	flag.StringVar(&envs.host, "host", envs.host, "HTTP service host name")

// 	if env, ok := os.LookupEnv("ENVIRONMENT"); ok && env != "" {
// 		envs.environment = env
// 	}
// 	flag.StringVar(&envs.environment, "environment", envs.environment, "Runtime environment [dev, hlg, stg, prd]")

// 	if env, ok := os.LookupEnv("HOST_PORT"); ok && env != "" {
// 		envs.port = env
// 	}
// 	flag.StringVar(&envs.port, "port", envs.port, "HTTP service port")

// 	if env, ok := os.LookupEnv("LOG_LEVEL"); ok && env != "" {
// 		envs.logLevel = env
// 	}
// 	flag.StringVar(&envs.logLevel, "log-level", envs.logLevel, "Log output level for the server [debug, info, trace]")

// 	if env, ok := os.LookupEnv("LOG_JSON"); ok && env != "" {
// 		b, _ := strconv.ParseBool(env)
// 		envs.logJSON = b
// 	}
// 	flag.BoolVar(&envs.logJSON, "log-json", envs.logJSON, "Log output level for the server [debug, info, trace]")

// 	if env, ok := os.LookupEnv("GRACEFUL_TIMEOUT"); ok && env != "" {
// 		if w, err := time.ParseDuration(env); err == nil {
// 			envs.wait = w
// 		}
// 	}
// 	flag.DurationVar(&envs.wait, "graceful-timeout", envs.wait, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

// 	flag.Parse()
// 	return envs
// }
