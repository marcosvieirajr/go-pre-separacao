package config

import (
	"flag"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// default values
var (
	ENVIRONMENT  = "dev"
	HOST_NAME    = ""
	HOST_PORT    = "8080"
	ALLOWED_CORS = ""
	LOG_LEVEL    = "info"
	// LOG_FORMAT   = "text"
	GRACEFUL_TIMEOUT = time.Second * 15
)

func Load() {
	// load env variables
	loadEnvs()

	// flags take precedence over env vars
	loadFlags()
}

func loadEnvs() {
	// Load dev env vars from .env file. It's important to note that
	// it WILL NOT OVERRIDE an env variable that already exists
	godotenv.Load(".env")

	HOST_NAME = os.Getenv("HOST_NAME")
	ALLOWED_CORS = os.Getenv("ALLOWED_CORS")

	if env, ok := os.LookupEnv("ENVIRONMENT"); ok && env != "" {
		ENVIRONMENT = env
	}

	if env, ok := os.LookupEnv("HOST_PORT"); ok && env != "" {
		HOST_PORT = env
	}

	if env, ok := os.LookupEnv("LOG_LEVEL"); ok && env != "" {
		LOG_LEVEL = env
	}

	// if env, ok := os.LookupEnv("LOG_FORMAT"); ok && env != "" {
	// 	LOG_FORMAT = env
	// }

	if env, ok := os.LookupEnv("GRACEFUL_TIMEOUT"); ok && env != "" {
		if wait, err := time.ParseDuration(env); err == nil {
			GRACEFUL_TIMEOUT = wait
		}
	}
}

func loadFlags() {
	flag.StringVar(&ALLOWED_CORS, "cors", ALLOWED_CORS, "Cross-Origin Resource Sharing")
	flag.StringVar(&HOST_NAME, "host", HOST_NAME, "HTTP service host name")
	flag.StringVar(&ENVIRONMENT, "environment", ENVIRONMENT, "Runtime environment [dev, hlg, stg, prod]. Default: dev")
	flag.StringVar(&HOST_PORT, "port", HOST_PORT, "HTTP service port. Default: 8080")
	flag.StringVar(&LOG_LEVEL, "log-level", LOG_LEVEL, "Log output level for the server [debug, info, trace]. Default: info")
	// flag.StringVar(&LOG_FORMAT, "log-format", LOG_FORMAT, "Log output format [text, json]. Default: text")
	flag.DurationVar(&GRACEFUL_TIMEOUT, "wait", GRACEFUL_TIMEOUT, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m. Default: 15s")
	flag.Parse()
}
