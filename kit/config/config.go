package config

import (
	"flag"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	Host        string
	Port        string
	AllowedCORS string
	LogFormat   string
	LogLevel    string
	Wait        time.Duration
}

func Load() Config {

	godotenv.Load(".env")

	cfg := Config{
		Environment: "dev",
		AllowedCORS: os.Getenv("ALLOWED_CORS"),
		Host:        os.Getenv("HOST_NAME"),
		Port:        "8080",
		LogLevel:    "info",
		LogFormat:   "text",
		Wait:        time.Second * 15,
	}

	flag.StringVar(&cfg.AllowedCORS, "cors", cfg.AllowedCORS, "Cross-Origin Resource Sharing")
	flag.StringVar(&cfg.Host, "host", cfg.Host, "HTTP service host name")

	if env, ok := os.LookupEnv("ENVIRONMENT"); ok && env != "" {
		cfg.Environment = env
	}
	flag.StringVar(&cfg.Environment, "environment", cfg.Environment, "Runtime environment [dev, hlg, stg, prd]. Default: dev")

	if env, ok := os.LookupEnv("HOST_PORT"); ok && env != "" {
		cfg.Port = env
	}
	flag.StringVar(&cfg.Port, "port", cfg.Port, "HTTP service port. Default: 8080")

	if env, ok := os.LookupEnv("LOG_LEVEL"); ok && env != "" {
		cfg.LogLevel = env
	}
	flag.StringVar(&cfg.LogLevel, "log-level", cfg.LogLevel, "Log output level for the server [debug, info, trace]. Default: info")

	if env, ok := os.LookupEnv("LOG_FORMAT"); ok && env != "" {
		cfg.LogFormat = env
	}
	flag.StringVar(&cfg.LogFormat, "log-format", cfg.LogFormat, "Log output format [text, json]. Default: text")

	if env, ok := os.LookupEnv("GRACEFUL_TIMEOUT"); ok && env != "" {
		if w, err := time.ParseDuration(env); err == nil {
			cfg.Wait = w
		}
	}
	flag.DurationVar(&cfg.Wait, "graceful-timeout", cfg.Wait, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m. Default: 15s")

	flag.Parse()
	return cfg
}
